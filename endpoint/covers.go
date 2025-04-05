package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type Covers struct {
	BaseEndpoint[pb.Cover]
}

func NewCovers(request func(URL string, dataBody any) (*resty.Response, error)) *Covers {
	a := &Covers{
		BaseEndpoint[pb.Cover]{
			endpointName: EPCovers,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Covers) Query(query string) ([]*pb.Cover, error) {
	resp, err := a.request("https://api.igdb.com/v4/covers.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CoverResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Covers) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Covers, nil
}
