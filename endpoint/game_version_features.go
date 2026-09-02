package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type GameVersionFeatures struct {
	BaseEndpoint[pb.GameVersionFeature]
}

func NewGameVersionFeatures(request RequestFunc) *GameVersionFeatures {
	a := &GameVersionFeatures{
		BaseEndpoint[pb.GameVersionFeature]{
			endpointName: EPGameVersionFeatures,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameVersionFeatureResult) []*pb.GameVersionFeature { return r.Gameversionfeatures })
	return a
}
