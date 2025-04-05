package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetCharacterMugShots(query string) ([]*pb.CharacterMugShot, error) {
	resp, err := g.Request("https://api.igdb.com/v4/character_mug_shots.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CharacterMugShotResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Charactermugshots) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Charactermugshots, nil
}

func (g *Client) GetCharacterMugShotByID(id uint64) (*pb.CharacterMugShot, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	characterMugShots, err := g.GetCharacterMugShots(query)
	if err != nil {
		return nil, err
	}
	return characterMugShots[0], nil
}

func (g *Client) GetCharacterMugShotsByIDs(ids []uint64) ([]*pb.CharacterMugShot, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCharacterMugShots(idStr)
}

func (g *Client) GetCharacterMugShotsLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	characterMugShots, err := g.GetCharacterMugShots(query)
	if err != nil {
		return 0, err
	}
	return int(characterMugShots[0].Id), nil
}
