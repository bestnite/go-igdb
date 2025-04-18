package endpoint

import (
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

func (a *CompanyWebsites) Query(query string) ([]*pb.CompanyWebsite, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CompanyWebsiteResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Companywebsites) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Companywebsites, nil
}
