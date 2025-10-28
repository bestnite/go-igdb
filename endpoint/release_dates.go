package endpoint

import (
	"context"
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type ReleaseDates struct {
	BaseEndpoint[pb.ReleaseDate]
}

func NewReleaseDates(request RequestFunc) *ReleaseDates {
	a := &ReleaseDates{
		BaseEndpoint: BaseEndpoint[pb.ReleaseDate]{
			endpointName: EPReleaseDates,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *ReleaseDates) Query(ctx context.Context, query string) ([]*pb.ReleaseDate, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ReleaseDateResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Releasedates, nil
}
