package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type GameModes struct {
	BaseEndpoint[pb.GameMode]
}

func NewGameModes(request RequestFunc) *GameModes {
	a := &GameModes{
		BaseEndpoint[pb.GameMode]{
			endpointName: EPGameModes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameModeResult) []*pb.GameMode { return r.Gamemodes })
	return a
}
