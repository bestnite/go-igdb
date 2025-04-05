package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetMultiplayerModes(query string) ([]*pb.MultiplayerMode, error) {
	resp, err := g.Request("https://api.igdb.com/v4/multiplayer_modes.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.MultiplayerModeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Multiplayermodes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Multiplayermodes, nil
}

func (g *igdb) GetMultiplayerModeByID(id uint64) (*pb.MultiplayerMode, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	multiplayerModes, err := g.GetMultiplayerModes(query)
	if err != nil {
		return nil, err
	}
	return multiplayerModes[0], nil
}

func (g *igdb) GetMultiplayerModesByIDs(ids []uint64) ([]*pb.MultiplayerMode, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetMultiplayerModes(idStr)
}

func (g *igdb) GetMultiplayerModesByGameID(id uint64) ([]*pb.MultiplayerMode, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetMultiplayerModes(query)
}

func (g *igdb) GetMultiplayerModesByGameIDs(ids []uint64) ([]*pb.MultiplayerMode, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetMultiplayerModes(idStr)
}

func (g *igdb) GetMultiplayerModesByPlatformID(id uint64) ([]*pb.MultiplayerMode, error) {
	query := fmt.Sprintf(`where platform = %d; fields *;`, id)
	return g.GetMultiplayerModes(query)
}

func (g *igdb) GetMultiplayerModesByPlatformIDs(ids []uint64) ([]*pb.MultiplayerMode, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where platform = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetMultiplayerModes(idStr)
}

func (g *igdb) GetMultiplayerModesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	multiplayerModes, err := g.GetMultiplayerModes(query)
	if err != nil {
		return 0, err
	}
	return int(multiplayerModes[0].Id), nil
}
