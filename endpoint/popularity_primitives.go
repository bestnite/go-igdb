package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type PopularityPrimitives struct {
	BaseEndpoint[pb.PopularityPrimitive]
}

func NewPopularityPrimitives(request RequestFunc) *PopularityPrimitives {
	a := &PopularityPrimitives{
		BaseEndpoint[pb.PopularityPrimitive]{
			endpointName: EPPopularityPrimitives,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.PopularityPrimitiveResult) []*pb.PopularityPrimitive { return r.Popularityprimitives })
	return a
}
