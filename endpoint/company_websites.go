package endpoint

import (
	"context"
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
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
	a.queryFunc = a.Query
	return a
}

func (a *CompanyWebsites) Query(ctx context.Context, query string) ([]*pb.CompanyWebsite, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CompanyWebsiteResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Companywebsites, nil
}
