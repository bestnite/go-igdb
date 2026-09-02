package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type PlatformFamilies struct {
	BaseEndpoint[pb.PlatformFamily]
}

func NewPlatformFamilies(request RequestFunc) *PlatformFamilies {
	a := &PlatformFamilies{
		BaseEndpoint[pb.PlatformFamily]{
			endpointName: EPPlatformFamilies,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.PlatformFamilyResult) []*pb.PlatformFamily { return r.Platformfamilies })
	return a
}
