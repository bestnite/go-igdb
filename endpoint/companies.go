package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Companies struct {
	BaseEndpoint[pb.Company]
}

func NewCompanies(request RequestFunc) *Companies {
	a := &Companies{
		BaseEndpoint[pb.Company]{
			endpointName: EPCompanies,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CompanyResult) []*pb.Company { return r.Companies })
	return a
}
