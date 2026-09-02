package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type GameEngineLogos struct {
	BaseEndpoint[pb.GameEngineLogo]
}

func NewGameEngineLogos(request RequestFunc) *GameEngineLogos {
	a := &GameEngineLogos{
		BaseEndpoint[pb.GameEngineLogo]{
			endpointName: EPGameEngineLogos,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameEngineLogoResult) []*pb.GameEngineLogo { return r.Gameenginelogos })
	return a
}
