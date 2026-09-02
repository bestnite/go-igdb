package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Covers struct {
	BaseEndpoint[pb.Cover]
}

func NewCovers(request RequestFunc) *Covers {
	a := &Covers{
		BaseEndpoint[pb.Cover]{
			endpointName: EPCovers,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CoverResult) []*pb.Cover { return r.Covers })
	return a
}
