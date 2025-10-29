package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Themes struct {
	BaseEndpoint[pb.Theme]
}

func NewThemes(request RequestFunc) *Themes {
	a := &Themes{
		BaseEndpoint: BaseEndpoint[pb.Theme]{
			endpointName: EPThemes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Themes) Query(ctx context.Context, query string) ([]*pb.Theme, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ThemeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Themes, nil
}
