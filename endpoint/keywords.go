package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Keywords struct {
	BaseEndpoint[pb.Keyword]
}

func NewKeywords(request RequestFunc) *Keywords {
	a := &Keywords{
		BaseEndpoint[pb.Keyword]{
			endpointName: EPKeywords,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.KeywordResult) []*pb.Keyword { return r.Keywords })
	return a
}
