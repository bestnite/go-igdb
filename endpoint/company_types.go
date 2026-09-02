package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CompanyTypes struct {
	BaseEndpoint[pb.CompanyType]
}

func NewCompanyTypes(request RequestFunc) *CompanyTypes {
	a := &CompanyTypes{
		BaseEndpoint[pb.CompanyType]{
			endpointName: EPCompanyTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CompanyTypeResult) []*pb.CompanyType { return r.Companytypes })
	return a
}
