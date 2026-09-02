package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type CompanyLogos struct {
	BaseEndpoint[pb.CompanyLogo]
}

func NewCompanyLogos(request RequestFunc) *CompanyLogos {
	a := &CompanyLogos{
		BaseEndpoint[pb.CompanyLogo]{
			endpointName: EPCompanyLogos,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.CompanyLogoResult) []*pb.CompanyLogo { return r.Companylogos })
	return a
}
