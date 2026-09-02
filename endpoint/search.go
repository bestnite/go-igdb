package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Search struct {
	BaseEndpoint[pb.Search]
}

func NewSearch(request RequestFunc) *Search {
	a := &Search{
		BaseEndpoint[pb.Search]{
			endpointName: EPSearch,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.SearchResult) []*pb.Search { return r.Searches })
	return a
}
