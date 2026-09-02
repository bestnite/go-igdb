package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type LanguageSupports struct {
	BaseEndpoint[pb.LanguageSupport]
}

func NewLanguageSupports(request RequestFunc) *LanguageSupports {
	a := &LanguageSupports{
		BaseEndpoint[pb.LanguageSupport]{
			endpointName: EPLanguageSupports,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.LanguageSupportResult) []*pb.LanguageSupport { return r.Languagesupports })
	return a
}
