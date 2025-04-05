package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type CharacterGenders struct{ BaseEndpoint }

func (a *CharacterGenders) Query(query string) ([]*pb.CharacterGender, error) {
	resp, err := a.request("https://api.igdb.com/v4/character_genders.pb", query)
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
