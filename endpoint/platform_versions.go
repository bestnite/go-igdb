package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type PlatformVersions struct {
	BaseEndpoint[pb.PlatformVersion]
}

func NewPlatformVersions(request RequestFunc) *PlatformVersions {
	a := &PlatformVersions{
		BaseEndpoint[pb.PlatformVersion]{
			endpointName: EPPlatformVersions,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.PlatformVersionResult) []*pb.PlatformVersion { return r.Platformversions })
	return a
}
