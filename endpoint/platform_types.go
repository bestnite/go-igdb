package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type PlatformTypes struct {
	BaseEndpoint[pb.PlatformType]
}

func NewPlatformTypes(request RequestFunc) *PlatformTypes {
	a := &PlatformTypes{
		BaseEndpoint[pb.PlatformType]{
			endpointName: EPPlatformTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.PlatformTypeResult) []*pb.PlatformType { return r.Platformtypes })
	return a
}
