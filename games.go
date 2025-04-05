package igdb

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetGames(query string) ([]*pb.Game, error) {
	resp, err := g.Request("https://api.igdb.com/v4/games.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Games) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Games, nil
}

func (g *Client) GetParentGameID(id uint64) (uint64, error) {
	detail, err := GetItemByID(id, g.GetGames)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch IGDB app detail for parent: %d: %w", id, err)
	}
	hasParent := false
	if detail.ParentGame != nil && detail.ParentGame.Id != 0 {
		hasParent = true
		detail, err = GetItemByID(detail.ParentGame.Id, g.GetGames)
		if err != nil {
			return 0, fmt.Errorf("failed to fetch IGDB version parent: %d: %w", detail.VersionParent.Id, err)
		}
	}
	for detail.VersionParent != nil && detail.VersionParent.Id != 0 {
		hasParent = true
		detail, err = GetItemByID(detail.VersionParent.Id, g.GetGames)
		if err != nil {
			return 0, fmt.Errorf("failed to fetch IGDB version parent: %d: %w", detail.VersionParent.Id, err)
		}
	}

	if hasParent {
		return detail.Id, nil
	}

	return id, nil
}
