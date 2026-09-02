package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type GameReleaseFormats struct {
	BaseEndpoint[pb.GameReleaseFormat]
}

func NewGameReleaseFormats(request RequestFunc) *GameReleaseFormats {
	a := &GameReleaseFormats{
		BaseEndpoint[pb.GameReleaseFormat]{
			endpointName: EPGameReleaseFormats,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameReleaseFormatResult) []*pb.GameReleaseFormat { return r.Gamereleaseformats })
	return a
}
