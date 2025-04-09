package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type CharacterMugShots struct {
	BaseEndpoint[pb.CharacterMugShot]
}

func NewCharacterMugShots(request RequestFunc) *CharacterMugShots {
	a := &CharacterMugShots{
		BaseEndpoint[pb.CharacterMugShot]{
			endpointName: EPCharacterMugShots,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *CharacterMugShots) Query(query string) ([]*pb.CharacterMugShot, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
