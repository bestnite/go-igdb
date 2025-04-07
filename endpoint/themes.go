package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type Themes struct {
	BaseEndpoint[pb.Theme]
}

func NewThemes(request func(URL string, dataBody any) (*resty.Response, error)) *Themes {
	a := &Themes{
		BaseEndpoint[pb.Theme]{
			endpointName: EPThemes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Themes) Query(query string) ([]*pb.Theme, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ThemeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Themes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Themes, nil
}
