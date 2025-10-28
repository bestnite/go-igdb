package endpoint

import (
	"context"
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type CompanyStatuses struct {
	BaseEndpoint[pb.CompanyStatus]
}

func NewCompanyStatuses(request RequestFunc) *CompanyStatuses {
	a := &CompanyStatuses{
		BaseEndpoint[pb.CompanyStatus]{
			endpointName: EPCompanyStatuses,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *CompanyStatuses) Query(ctx context.Context, query string) ([]*pb.CompanyStatus, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CompanyStatusResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Companystatuses, nil
}
