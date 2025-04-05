package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type InvolvedCompanies struct{ BaseEndpoint }

func (a *InvolvedCompanies) Query(query string) ([]*pb.InvolvedCompany, error) {
	resp, err := a.request("https://api.igdb.com/v4/involved_companies.pb", query)
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
