package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type AgeRatings struct {
	BaseEndpoint[pb.AgeRating]
}

func NewAgeRatings(request func(URL string, dataBody any) (*resty.Response, error)) *AgeRatings {
	a := &AgeRatings{
		BaseEndpoint[pb.AgeRating]{
			endpointName: EPAgeRatings,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *AgeRatings) Query(query string) ([]*pb.AgeRating, error) {
	resp, err := a.request("https://api.igdb.com/v4/age_ratings.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.AgeRatingResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Ageratings) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Ageratings, nil
}
