package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type CompanyStatuses struct {
	BaseEndpoint[pb.CompanyStatus]
}

func NewCompanyStatuses(request func(URL string, dataBody any) (*resty.Response, error)) *CompanyStatuses {
	a := &CompanyStatuses{
		BaseEndpoint[pb.CompanyStatus]{
			endpointName: EPCompanyStatuses,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *CompanyStatuses) Query(query string) ([]*pb.CompanyStatus, error) {
	resp, err := a.request("https://api.igdb.com/v4/company_statuses.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CompanyStatusResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Companystatuses) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Companystatuses, nil
}
