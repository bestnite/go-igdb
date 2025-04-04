package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type CharacterMugShots struct {
	BaseEndpoint[pb.CharacterMugShot]
}

func NewCharacterMugShots(request func(URL string, dataBody any) (*resty.Response, error)) *CharacterMugShots {
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
	resp, err := a.request("https://api.igdb.com/v4/character_mug_shots.pb", query)
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
