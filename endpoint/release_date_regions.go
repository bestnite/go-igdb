package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type ReleaseDateRegions struct {
	BaseEndpoint[pb.ReleaseDateRegion]
}

func NewReleaseDateRegions(request RequestFunc) *ReleaseDateRegions {
	a := &ReleaseDateRegions{
		BaseEndpoint[pb.ReleaseDateRegion]{
			endpointName: EPReleaseDateRegions,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.ReleaseDateRegionResult) []*pb.ReleaseDateRegion { return r.Releasedateregions })
	return a
}
