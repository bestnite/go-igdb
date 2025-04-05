package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type Artworks struct {
	BaseEndpoint[pb.Artwork]
}

func NewArtworks(request func(URL string, dataBody any) (*resty.Response, error)) *Artworks {
	a := &Artworks{
		BaseEndpoint[pb.Artwork]{
			endpointName: EPArtworks,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Artworks) Query(query string) ([]*pb.Artwork, error) {
	resp, err := a.request("https://api.igdb.com/v4/artworks.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ArtworkResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Artworks) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Artworks, nil
}
