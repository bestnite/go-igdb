package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type GameTypes struct {
	BaseEndpoint[pb.GameType]
}

func NewGameTypes(request RequestFunc) *GameTypes {
	a := &GameTypes{
		BaseEndpoint[pb.GameType]{
			endpointName: EPGameTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameTypeResult) []*pb.GameType { return r.Gametypes })
	return a
}
