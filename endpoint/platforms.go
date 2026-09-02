package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Platforms struct {
	BaseEndpoint[pb.Platform]
}

func NewPlatforms(request RequestFunc) *Platforms {
	a := &Platforms{
		BaseEndpoint[pb.Platform]{
			endpointName: EPPlatforms,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.PlatformResult) []*pb.Platform { return r.Platforms })
	return a
}
