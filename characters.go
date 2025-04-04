package igdb

import (
	"fmt"
	"strings"

	pb "github/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCharacters(query string) ([]*pb.Character, error) {
	resp, err := g.Request("https://api.igdb.com/v4/characters.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CharacterResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Characters) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Characters, nil
}

func (g *igdb) GetCharacterByID(id uint64) (*pb.Character, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	characters, err := g.GetCharacters(query)
	if err != nil {
		return nil, err
	}
	return characters[0], nil
}

func (g *igdb) GetCharactersByIDs(ids []uint64) ([]*pb.Character, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCharacters(idStr)
}
