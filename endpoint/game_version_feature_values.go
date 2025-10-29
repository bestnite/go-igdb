package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type GameVersionFeatureValues struct {
	BaseEndpoint[pb.GameVersionFeatureValue]
}

func NewGameVersionFeatureValues(request RequestFunc) *GameVersionFeatureValues {
	a := &GameVersionFeatureValues{
		BaseEndpoint: BaseEndpoint[pb.GameVersionFeatureValue]{
			endpointName: EPGameVersionFeatureValues,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameVersionFeatureValues) Query(ctx context.Context, query string) ([]*pb.GameVersionFeatureValue, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameVersionFeatureValueResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Gameversionfeaturevalues, nil
}
