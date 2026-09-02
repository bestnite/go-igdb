package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type AlternativeNames struct {
	BaseEndpoint[pb.AlternativeName]
}

func NewAlternativeNames(request RequestFunc) *AlternativeNames {
	a := &AlternativeNames{
		BaseEndpoint[pb.AlternativeName]{
			endpointName: EPAlternativeNames,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.AlternativeNameResult) []*pb.AlternativeName { return r.Alternativenames })
	return a
}
