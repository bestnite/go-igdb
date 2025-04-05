package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetMultiplayerModes(query string) ([]*pb.MultiplayerMode, error) {
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

func (g *Client) GetMultiplayerModeByID(id uint64) (*pb.MultiplayerMode, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	multiplayerModes, err := g.GetMultiplayerModes(query)
	if err != nil {
		return nil, err
	}
	return multiplayerModes[0], nil
}

func (g *Client) GetMultiplayerModesByIDs(ids []uint64) ([]*pb.MultiplayerMode, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetMultiplayerModes(idStr)
}

func (g *Client) GetMultiplayerModesByGameID(id uint64) ([]*pb.MultiplayerMode, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetMultiplayerModes(query)
}

func (g *Client) GetMultiplayerModesByGameIDs(ids []uint64) ([]*pb.MultiplayerMode, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetMultiplayerModes(idStr)
}

func (g *Client) GetMultiplayerModesByPlatformID(id uint64) ([]*pb.MultiplayerMode, error) {
	query := fmt.Sprintf(`where platform = %d; fields *;`, id)
	return g.GetMultiplayerModes(query)
}

func (g *Client) GetMultiplayerModesByPlatformIDs(ids []uint64) ([]*pb.MultiplayerMode, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where platform = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetMultiplayerModes(idStr)
}

func (g *Client) GetMultiplayerModesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	multiplayerModes, err := g.GetMultiplayerModes(query)
	if err != nil {
		return 0, err
	}
	return int(multiplayerModes[0].Id), nil
}
