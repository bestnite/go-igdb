package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Search struct {
	endpointName Name
	request      RequestFunc
}

func NewSearch(request RequestFunc) *Search {
	return &Search{
		endpointName: EPSearch,
		request:      request,
	}
}

func (a *Search) Search(ctx context.Context, query string) ([]*pb.Search, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.SearchResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Searches, nil
}
