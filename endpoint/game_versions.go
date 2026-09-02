package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type GameVersions struct {
	BaseEndpoint[pb.GameVersion]
}

func NewGameVersions(request RequestFunc) *GameVersions {
	a := &GameVersions{
		BaseEndpoint[pb.GameVersion]{
			endpointName: EPGameVersions,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GameVersionResult) []*pb.GameVersion { return r.Gameversions })
	return a
}
