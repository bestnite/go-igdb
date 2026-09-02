package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CompanyWebsites struct {
	BaseEndpoint[pb.CompanyWebsite]
}

func NewCompanyWebsites(request RequestFunc) *CompanyWebsites {
	a := &CompanyWebsites{
		BaseEndpoint[pb.CompanyWebsite]{
			endpointName: EPCompanyWebsites,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CompanyWebsiteResult) []*pb.CompanyWebsite { return r.Companywebsites })
	return a
}
