package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type InvolvedCompanies struct {
	BaseEndpoint[pb.InvolvedCompany]
}

func NewInvolvedCompanies(request RequestFunc) *InvolvedCompanies {
	a := &InvolvedCompanies{
		BaseEndpoint[pb.InvolvedCompany]{
			endpointName: EPInvolvedCompanies,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *InvolvedCompanies) Query(query string) ([]*pb.InvolvedCompany, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.InvolvedCompanyResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Involvedcompanies) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Involvedcompanies, nil
}
