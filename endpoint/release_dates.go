package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type ReleaseDates struct {
	BaseEndpoint[pb.ReleaseDate]
}

func NewReleaseDates(request func(URL string, dataBody any) (*resty.Response, error)) *ReleaseDates {
	a := &ReleaseDates{
		BaseEndpoint[pb.ReleaseDate]{
			endpointName: EPReleaseDates,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *ReleaseDates) Query(query string) ([]*pb.ReleaseDate, error) {
	resp, err := a.request("https://api.igdb.com/v4/release_dates.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ReleaseDateResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Releasedates) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Releasedates, nil
}
