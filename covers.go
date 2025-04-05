package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetCovers(query string) ([]*pb.Cover, error) {
	resp, err := g.Request("https://api.igdb.com/v4/covers.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CoverResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Covers) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Covers, nil
}

func (g *Client) GetCoverByID(id uint64) (*pb.Cover, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	covers, err := g.GetCovers(query)
	if err != nil {
		return nil, err
	}
	return covers[0], nil
}

func (g *Client) GetCoversByIDs(ids []uint64) ([]*pb.Cover, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCovers(idStr)
}

func (g *Client) GetCoversByGameID(id uint64) ([]*pb.Cover, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetCovers(query)
}

func (g *Client) GetCoversByGameIDs(ids []uint64) ([]*pb.Cover, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCovers(idStr)
}

func (g *Client) GetCoversByGameLocalizationID(id uint64) ([]*pb.Cover, error) {
	query := fmt.Sprintf(`where game_localization = %d; fields *;`, id)
	return g.GetCovers(query)
}

func (g *Client) GetCoversByGameLocalizationIDs(ids []uint64) ([]*pb.Cover, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game_localization = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCovers(idStr)
}

func (g *Client) GetCoversLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	covers, err := g.GetCovers(query)
	if err != nil {
		return 0, err
	}
	return int(covers[0].Id), nil
}
