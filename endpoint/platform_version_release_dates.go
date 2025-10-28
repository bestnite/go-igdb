package endpoint

import (
	"context"
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type PlatformVersionReleaseDates struct {
	BaseEndpoint[pb.PlatformVersionReleaseDate]
}

func NewPlatformVersionReleaseDates(request RequestFunc) *PlatformVersionReleaseDates {
	a := &PlatformVersionReleaseDates{
		BaseEndpoint: BaseEndpoint[pb.PlatformVersionReleaseDate]{
			endpointName: EPPlatformVersionReleaseDates,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlatformVersionReleaseDates) Query(ctx context.Context, query string) ([]*pb.PlatformVersionReleaseDate, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformVersionReleaseDateResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Platformversionreleasedates, nil
}
