package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Genres struct {
	BaseEndpoint[pb.Genre]
}

func NewGenres(request RequestFunc) *Genres {
	a := &Genres{
		BaseEndpoint[pb.Genre]{
			endpointName: EPGenres,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Genres) Query(query string) ([]*pb.Genre, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GenreResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Genres) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Genres, nil
}
