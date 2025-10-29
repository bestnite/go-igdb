package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Websites struct {
	BaseEndpoint[pb.Website]
}

func NewWebsites(request RequestFunc) *Websites {
	a := &Websites{
		BaseEndpoint: BaseEndpoint[pb.Website]{
			endpointName: EPWebsites,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Websites) Query(ctx context.Context, query string) ([]*pb.Website, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.WebsiteResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Websites, nil
}
