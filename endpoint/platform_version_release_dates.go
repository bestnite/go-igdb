package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type PlatformVersionReleaseDates struct {
	BaseEndpoint[pb.PlatformVersionReleaseDate]
}

func NewPlatformVersionReleaseDates(request func(URL string, dataBody any) (*resty.Response, error)) *PlatformVersionReleaseDates {
	a := &PlatformVersionReleaseDates{
		BaseEndpoint[pb.PlatformVersionReleaseDate]{
			endpointName: EPPlatformVersionReleaseDates,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlatformVersionReleaseDates) Query(query string) ([]*pb.PlatformVersionReleaseDate, error) {
	resp, err := a.request("https://api.igdb.com/v4/platform_version_release_dates.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformVersionReleaseDateResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformversionreleasedates) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformversionreleasedates, nil
}
