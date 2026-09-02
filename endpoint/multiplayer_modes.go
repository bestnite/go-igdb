package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type MultiplayerModes struct {
	BaseEndpoint[pb.MultiplayerMode]
}

func NewMultiplayerModes(request RequestFunc) *MultiplayerModes {
	a := &MultiplayerModes{
		BaseEndpoint[pb.MultiplayerMode]{
			endpointName: EPMultiplayerModes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.MultiplayerModeResult) []*pb.MultiplayerMode { return r.Multiplayermodes })
	return a
}
