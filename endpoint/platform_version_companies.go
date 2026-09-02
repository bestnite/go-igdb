package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type PlatformVersionCompanies struct {
	BaseEndpoint[pb.PlatformVersionCompany]
}

func NewPlatformVersionCompanies(request RequestFunc) *PlatformVersionCompanies {
	a := &PlatformVersionCompanies{
		BaseEndpoint[pb.PlatformVersionCompany]{
			endpointName: EPPlatformVersionCompanies,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.PlatformVersionCompanyResult) []*pb.PlatformVersionCompany {
		return r.Platformversioncompanies
	})
	return a
}
