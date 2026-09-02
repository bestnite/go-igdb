package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type PlatformLogos struct {
	BaseEndpoint[pb.PlatformLogo]
}

func NewPlatformLogos(request RequestFunc) *PlatformLogos {
	a := &PlatformLogos{
		BaseEndpoint[pb.PlatformLogo]{
			endpointName: EPPlatformLogos,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.PlatformLogoResult) []*pb.PlatformLogo { return r.Platformlogos })
	return a
}
