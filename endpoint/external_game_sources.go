package endpoint

import (
	"context"
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type ExternalGameSources struct {
	BaseEndpoint[pb.ExternalGameSource]
}

func NewExternalGameSources(request RequestFunc) *ExternalGameSources {
	a := &ExternalGameSources{
		BaseEndpoint[pb.ExternalGameSource]{
			endpointName: EPExternalGameSources,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *ExternalGameSources) Query(ctx context.Context, query string) ([]*pb.ExternalGameSource, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ExternalGameSourceResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Externalgamesources, nil
}
