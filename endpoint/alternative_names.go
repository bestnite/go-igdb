package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type AlternativeNames struct {
	BaseEndpoint[pb.AlternativeName]
}

func NewAlternativeNames(request RequestFunc) *AlternativeNames {
	a := &AlternativeNames{
		BaseEndpoint[pb.AlternativeName]{
			endpointName: EPAlternativeNames,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *AlternativeNames) Query(query string) ([]*pb.AlternativeName, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.AlternativeNameResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Alternativenames) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Alternativenames, nil
}
