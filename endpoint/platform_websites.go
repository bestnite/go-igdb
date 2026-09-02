package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type PlatformWebsites struct {
	BaseEndpoint[pb.PlatformWebsite]
}

func NewPlatformWebsites(request RequestFunc) *PlatformWebsites {
	a := &PlatformWebsites{
		BaseEndpoint[pb.PlatformWebsite]{
			endpointName: EPPlatformWebsites,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.PlatformWebsiteResult) []*pb.PlatformWebsite { return r.Platformwebsites })
	return a
}
