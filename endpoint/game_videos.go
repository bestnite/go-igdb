package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type GameVideos struct {
	BaseEndpoint[pb.GameVideo]
}

func NewGameVideos(request RequestFunc) *GameVideos {
	a := &GameVideos{
		BaseEndpoint[pb.GameVideo]{
			endpointName: EPGameVideos,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameVideoResult) []*pb.GameVideo { return r.Gamevideos })
	return a
}
