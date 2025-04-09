package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type AgeRatingOrganizations struct {
	BaseEndpoint[pb.AgeRatingOrganization]
}

func NewAgeRatingOrganizations(request RequestFunc) *AgeRatingOrganizations {
	a := &AgeRatingOrganizations{
		BaseEndpoint[pb.AgeRatingOrganization]{
			endpointName: EPAgeRatingOrganizations,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *AgeRatingOrganizations) Query(query string) ([]*pb.AgeRatingOrganization, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.AgeRatingOrganizationResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Ageratingorganizations) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Ageratingorganizations, nil
}
