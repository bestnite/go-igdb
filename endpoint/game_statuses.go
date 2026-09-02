package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type GameStatuses struct {
	BaseEndpoint[pb.GameStatus]
}

func NewGameStatuses(request RequestFunc) *GameStatuses {
	a := &GameStatuses{
		BaseEndpoint[pb.GameStatus]{
			endpointName: EPGameStatuses,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameStatusResult) []*pb.GameStatus { return r.Gamestatuses })
	return a
}
