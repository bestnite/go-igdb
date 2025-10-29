package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Keywords struct {
	BaseEndpoint[pb.Keyword]
}

func NewKeywords(request RequestFunc) *Keywords {
	a := &Keywords{
		BaseEndpoint: BaseEndpoint[pb.Keyword]{
			endpointName: EPKeywords,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Keywords) Query(ctx context.Context, query string) ([]*pb.Keyword, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.KeywordResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Keywords, nil
}
