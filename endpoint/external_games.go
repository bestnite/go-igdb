package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type ExternalGames struct {
	BaseEndpoint[pb.ExternalGame]
}

func NewExternalGames(request RequestFunc) *ExternalGames {
	a := &ExternalGames{
		BaseEndpoint[pb.ExternalGame]{
			endpointName: EPExternalGames,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.ExternalGameResult) []*pb.ExternalGame { return r.Externalgames })
	return a
}
