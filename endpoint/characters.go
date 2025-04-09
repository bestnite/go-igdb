package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Characters struct {
	BaseEndpoint[pb.Character]
}

func NewCharacters(request RequestFunc) *Characters {
	a := &Characters{
		BaseEndpoint[pb.Character]{
			endpointName: EPCharacters,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Characters) Query(query string) ([]*pb.Character, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
