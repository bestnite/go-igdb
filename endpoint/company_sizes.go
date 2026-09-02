package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CompanySizes struct {
	BaseEndpoint[pb.CompanySize]
}

func NewCompanySizes(request RequestFunc) *CompanySizes {
	a := &CompanySizes{
		BaseEndpoint[pb.CompanySize]{
			endpointName: EPCompanySizes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CompanySizeResult) []*pb.CompanySize { return r.Companysizes })
	return a
}
