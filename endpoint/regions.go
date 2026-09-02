package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Regions struct {
	BaseEndpoint[pb.Region]
}

func NewRegions(request RequestFunc) *Regions {
	a := &Regions{
		BaseEndpoint[pb.Region]{
			endpointName: EPRegions,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.RegionResult) []*pb.Region { return r.Regions })
	return a
}
