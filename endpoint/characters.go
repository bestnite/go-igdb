package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Characters struct{ BaseEndpoint }

func (a *Characters) Query(query string) ([]*pb.Character, error) {
	resp, err := a.request("https://api.igdb.com/v4/characters.pb", query)
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
