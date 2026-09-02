package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type GameTimeToBeats struct {
	BaseEndpoint[pb.GameTimeToBeat]
}

func NewGameTimeToBeats(request RequestFunc) *GameTimeToBeats {
	a := &GameTimeToBeats{
		BaseEndpoint[pb.GameTimeToBeat]{
			endpointName: EPGameTimeToBeats,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameTimeToBeatResult) []*pb.GameTimeToBeat { return r.Gametimetobeats })
	return a
}
