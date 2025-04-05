package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type AgeRatingContentDescriptions struct {
	BaseEndpoint[pb.AgeRatingContentDescription]
}

func NewAgeRatingContentDescriptions(request func(URL string, dataBody any) (*resty.Response, error)) *AgeRatingContentDescriptions {
	a := &AgeRatingContentDescriptions{
		BaseEndpoint[pb.AgeRatingContentDescription]{
			endpointName: EPAgeRatingContentDescriptions,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *AgeRatingContentDescriptions) Query(query string) ([]*pb.AgeRatingContentDescription, error) {
	resp, err := a.request("https://api.igdb.com/v4/age_rating_content_descriptions.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.AgeRatingContentDescriptionResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Ageratingcontentdescriptions) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Ageratingcontentdescriptions, nil
}
