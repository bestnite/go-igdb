package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type Screenshots struct {
	BaseEndpoint[pb.Screenshot]
}

func NewScreenshots(request func(URL string, dataBody any) (*resty.Response, error)) *Screenshots {
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
	resp, err := a.request("https://api.igdb.com/v4/screenshots.pb", query)
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
