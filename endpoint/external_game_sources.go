package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type ExternalGameSources struct {
	BaseEndpoint[pb.ExternalGameSource]
}

func NewExternalGameSources(request RequestFunc) *ExternalGameSources {
	a := &ExternalGameSources{
		BaseEndpoint[pb.ExternalGameSource]{
			endpointName: EPExternalGameSources,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.ExternalGameSourceResult) []*pb.ExternalGameSource { return r.Externalgamesources })
	return a
}
