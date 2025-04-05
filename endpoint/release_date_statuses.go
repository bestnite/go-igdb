package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type ReleaseDateStatuses struct {
	BaseEndpoint[pb.ReleaseDateStatus]
}

func NewReleaseDateStatuses(request func(URL string, dataBody any) (*resty.Response, error)) *ReleaseDateStatuses {
	a := &ReleaseDateStatuses{
		BaseEndpoint[pb.ReleaseDateStatus]{
			endpointName: EPReleaseDateStatuses,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *ReleaseDateStatuses) Query(query string) ([]*pb.ReleaseDateStatus, error) {
	resp, err := a.request("https://api.igdb.com/v4/release_date_statuses.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ReleaseDateStatusResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Releasedatestatuses) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Releasedatestatuses, nil
}
