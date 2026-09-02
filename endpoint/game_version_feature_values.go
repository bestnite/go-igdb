package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type GameVersionFeatureValues struct {
	BaseEndpoint[pb.GameVersionFeatureValue]
}

func NewGameVersionFeatureValues(request RequestFunc) *GameVersionFeatureValues {
	a := &GameVersionFeatureValues{
		BaseEndpoint[pb.GameVersionFeatureValue]{
			endpointName: EPGameVersionFeatureValues,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameVersionFeatureValueResult) []*pb.GameVersionFeatureValue {
		return r.Gameversionfeaturevalues
	})
	return a
}
