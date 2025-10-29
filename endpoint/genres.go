package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Genres struct {
	BaseEndpoint[pb.Genre]
}

func NewGenres(request RequestFunc) *Genres {
	a := &Genres{
		BaseEndpoint: BaseEndpoint[pb.Genre]{
			endpointName: EPGenres,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Genres) Query(ctx context.Context, query string) ([]*pb.Genre, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GenreResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Genres, nil
}
