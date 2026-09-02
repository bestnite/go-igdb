package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type PopularityTypes struct {
	BaseEndpoint[pb.PopularityType]
}

func NewPopularityTypes(request RequestFunc) *PopularityTypes {
	a := &PopularityTypes{
		BaseEndpoint[pb.PopularityType]{
			endpointName: EPPopularityTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.PopularityTypeResult) []*pb.PopularityType { return r.Popularitytypes })
	return a
}
