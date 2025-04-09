package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Themes struct {
	BaseEndpoint[pb.Theme]
}

func NewThemes(request RequestFunc) *Themes {
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
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
