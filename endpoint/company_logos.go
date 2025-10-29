package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
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
	a.queryFunc = a.Query
	return a
}

func (a *CompanyLogos) Query(ctx context.Context, query string) ([]*pb.CompanyLogo, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CompanyLogoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Companylogos, nil
}
