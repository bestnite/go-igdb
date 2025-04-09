package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Keywords struct {
	BaseEndpoint[pb.Keyword]
}

func NewKeywords(request RequestFunc) *Keywords {
	a := &Keywords{
		BaseEndpoint[pb.Keyword]{
			endpointName: EPKeywords,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Keywords) Query(query string) ([]*pb.Keyword, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.KeywordResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Keywords) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Keywords, nil
}
