package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Screenshots struct {
	BaseEndpoint[pb.Screenshot]
}

func NewScreenshots(request RequestFunc) *Screenshots {
	a := &Screenshots{
		BaseEndpoint[pb.Screenshot]{
			endpointName: EPScreenshots,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Screenshots) Query(query string) ([]*pb.Screenshot, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ScreenshotResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Screenshots) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Screenshots, nil
}
