package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetCharacterGenders(query string) ([]*pb.CharacterGender, error) {
	resp, err := g.Request("https://api.igdb.com/v4/character_genders.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CharacterGenderResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Charactergenders) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Charactergenders, nil
}

func (g *Client) GetCharacterGenderByID(id uint64) (*pb.CharacterGender, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	characterGenders, err := g.GetCharacterGenders(query)
	if err != nil {
		return nil, err
	}
	return characterGenders[0], nil
}

func (g *Client) GetCharacterGendersByIDs(ids []uint64) ([]*pb.CharacterGender, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCharacterGenders(idStr)
}

func (g *Client) GetCharacterGendersLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	characterGenders, err := g.GetCharacterGenders(query)
	if err != nil {
		return 0, err
	}
	return int(characterGenders[0].Id), nil
}
