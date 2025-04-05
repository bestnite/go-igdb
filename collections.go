package igdb

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetCollections(query string) ([]*pb.Collection, error) {
	resp, err := g.Request("https://api.igdb.com/v4/collections.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Collections) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Collections, nil
}
