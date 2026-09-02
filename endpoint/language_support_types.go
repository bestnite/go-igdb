package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type LanguageSupportTypes struct {
	BaseEndpoint[pb.LanguageSupportType]
}

func NewLanguageSupportTypes(request RequestFunc) *LanguageSupportTypes {
	a := &LanguageSupportTypes{
		BaseEndpoint[pb.LanguageSupportType]{
			endpointName: EPLanguageSupportTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.LanguageSupportTypeResult) []*pb.LanguageSupportType { return r.Languagesupporttypes })
	return a
}
