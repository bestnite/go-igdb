package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type AgeRatingOrganizations struct{ BaseEndpoint }

func (a *AgeRatingOrganizations) Query(query string) ([]*pb.AgeRatingOrganization, error) {
	resp, err := a.request("https://api.igdb.com/v4/age_rating_organizations.pb", query)
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
