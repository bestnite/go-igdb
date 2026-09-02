package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type PlayerPerspectives struct {
	BaseEndpoint[pb.PlayerPerspective]
}

func NewPlayerPerspectives(request RequestFunc) *PlayerPerspectives {
	a := &PlayerPerspectives{
		BaseEndpoint[pb.PlayerPerspective]{
			endpointName: EPPlayerPerspectives,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.PlayerPerspectiveResult) []*pb.PlayerPerspective { return r.Playerperspectives })
	return a
}
