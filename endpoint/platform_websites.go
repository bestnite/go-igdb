package endpoint

import (
	"context"
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type PlatformWebsites struct {
	BaseEndpoint[pb.PlatformWebsite]
}

func NewPlatformWebsites(request RequestFunc) *PlatformWebsites {
	a := &PlatformWebsites{
		BaseEndpoint: BaseEndpoint[pb.PlatformWebsite]{
			endpointName: EPPlatformWebsites,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlatformWebsites) Query(ctx context.Context, query string) ([]*pb.PlatformWebsite, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformWebsiteResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Platformwebsites, nil
}
