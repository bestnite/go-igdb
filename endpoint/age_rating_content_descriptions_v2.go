package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type AgeRatingContentDescriptionsV2 struct {
	BaseEndpoint
}

func (a *AgeRatingContentDescriptionsV2) Query(query string) ([]*pb.AgeRatingContentDescriptionV2, error) {
	resp, err := a.request("https://api.igdb.com/v4/age_rating_content_descriptions_v2.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.AgeRatingContentDescriptionV2Result{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Ageratingcontentdescriptionsv2) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Ageratingcontentdescriptionsv2, nil
}
