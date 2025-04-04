package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type Genres struct {
	BaseEndpoint[pb.Genre]
}

func NewGenres(request func(URL string, dataBody any) (*resty.Response, error)) *Genres {
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
	resp, err := a.request("https://api.igdb.com/v4/genres.pb", query)
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
