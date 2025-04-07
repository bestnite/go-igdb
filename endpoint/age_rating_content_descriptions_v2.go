package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type AgeRatingContentDescriptionsV2 struct {
	BaseEndpoint[pb.AgeRatingContentDescriptionV2]
}

func NewAgeRatingContentDescriptionsV2(request func(URL string, dataBody any) (*resty.Response, error)) *AgeRatingContentDescriptionsV2 {
	a := &AgeRatingContentDescriptionsV2{
		BaseEndpoint[pb.AgeRatingContentDescriptionV2]{
			endpointName: EPAgeRatingContentDescriptionsV2,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *AgeRatingContentDescriptionsV2) Query(query string) ([]*pb.AgeRatingContentDescriptionV2, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
