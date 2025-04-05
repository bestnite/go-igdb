package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetPlayerPerspectives(query string) ([]*pb.PlayerPerspective, error) {
	resp, err := g.Request("https://api.igdb.com/v4/player_perspectives.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlayerPerspectiveResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Playerperspectives) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Playerperspectives, nil
}

func (g *igdb) GetPlayerPerspectiveByID(id uint64) (*pb.PlayerPerspective, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	playerPerspectives, err := g.GetPlayerPerspectives(query)
	if err != nil {
		return nil, err
	}
	return playerPerspectives[0], nil
}

func (g *igdb) GetPlayerPerspectivesByIDs(ids []uint64) ([]*pb.PlayerPerspective, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlayerPerspectives(idStr)
}

func (g *igdb) GetPlayerPerspectivesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	playerPerspectives, err := g.GetPlayerPerspectives(query)
	if err != nil {
		return 0, err
	}
	return int(playerPerspectives[0].Id), nil
}
