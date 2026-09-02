package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Games struct {
	BaseEndpoint[pb.Game]
}

func NewGames(request RequestFunc) *Games {
	a := &Games{
		BaseEndpoint[pb.Game]{
			endpointName: EPGames,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameResult) []*pb.Game { return r.Games })
	return a
}
