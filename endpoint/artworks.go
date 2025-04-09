package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Artworks struct {
	BaseEndpoint[pb.Artwork]
}

func NewArtworks(request RequestFunc) *Artworks {
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
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
