package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type GameVersionFeatureValues struct {
	BaseEndpoint[pb.GameVersionFeatureValue]
}

func NewGameVersionFeatureValues(request func(URL string, dataBody any) (*resty.Response, error)) *GameVersionFeatureValues {
	a := &GameVersionFeatureValues{
		BaseEndpoint[pb.GameVersionFeatureValue]{
			endpointName: EPGameVersionFeatureValues,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameVersionFeatureValues) Query(query string) ([]*pb.GameVersionFeatureValue, error) {
	resp, err := a.request("https://api.igdb.com/v4/game_version_feature_values.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameVersionFeatureValueResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gameversionfeaturevalues) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gameversionfeaturevalues, nil
}
