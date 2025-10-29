package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type PlatformVersionCompanies struct {
	BaseEndpoint[pb.PlatformVersionCompany]
}

func NewPlatformVersionCompanies(request RequestFunc) *PlatformVersionCompanies {
	a := &PlatformVersionCompanies{
		BaseEndpoint: BaseEndpoint[pb.PlatformVersionCompany]{
			endpointName: EPPlatformVersionCompanies,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlatformVersionCompanies) Query(ctx context.Context, query string) ([]*pb.PlatformVersionCompany, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformVersionCompanyResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Platformversioncompanies, nil
}
