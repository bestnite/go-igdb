package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Languages struct {
	BaseEndpoint[pb.Language]
}

func NewLanguages(request RequestFunc) *Languages {
	a := &Languages{
		BaseEndpoint[pb.Language]{
			endpointName: EPLanguages,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.LanguageResult) []*pb.Language { return r.Languages })
	return a
}
