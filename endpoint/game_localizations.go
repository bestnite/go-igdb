package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type GameLocalizations struct {
	BaseEndpoint[pb.GameLocalization]
}

func NewGameLocalizations(request RequestFunc) *GameLocalizations {
	a := &GameLocalizations{
		BaseEndpoint[pb.GameLocalization]{
			endpointName: EPGameLocalizations,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameLocalizationResult) []*pb.GameLocalization { return r.Gamelocalizations })
	return a
}
