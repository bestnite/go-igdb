package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCharacterSpecies(query string) ([]*pb.CharacterSpecie, error) {
	resp, err := g.Request("https://api.igdb.com/v4/character_species.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CharacterSpecieResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Characterspecies) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Characterspecies, nil
}

func (g *igdb) GetCharacterSpecieByID(id uint64) (*pb.CharacterSpecie, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	characterSpecies, err := g.GetCharacterSpecies(query)
	if err != nil {
		return nil, err
	}
	return characterSpecies[0], nil
}

func (g *igdb) GetCharacterSpeciesByIDs(ids []uint64) ([]*pb.CharacterSpecie, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCharacterSpecies(idStr)
}
