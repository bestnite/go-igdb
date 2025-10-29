package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
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
	a.queryFunc = a.Query
	return a
}

func (a *Companies) Query(ctx context.Context, query string) ([]*pb.Company, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CompanyResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Companies, nil
}
