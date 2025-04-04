package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCharacterMugShots(query string) ([]*pb.CharacterMugShot, error) {
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

func (g *igdb) GetCharacterMugShotByID(id uint64) (*pb.CharacterMugShot, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	characterMugShots, err := g.GetCharacterMugShots(query)
	if err != nil {
		return nil, err
	}
	return characterMugShots[0], nil
}

func (g *igdb) GetCharacterMugShotsByIDs(ids []uint64) ([]*pb.CharacterMugShot, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCharacterMugShots(idStr)
}
