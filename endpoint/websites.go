package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Websites struct {
	BaseEndpoint[pb.Website]
}

func NewWebsites(request RequestFunc) *Websites {
	a := &Websites{
		BaseEndpoint[pb.Website]{
			endpointName: EPWebsites,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.WebsiteResult) []*pb.Website { return r.Websites })
	return a
}
