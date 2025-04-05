package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Artworks struct{ BaseEndpoint }

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
