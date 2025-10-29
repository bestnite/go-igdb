package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type CollectionTypes struct {
	BaseEndpoint[pb.CollectionType]
}

func NewCollectionTypes(request RequestFunc) *CollectionTypes {
	a := &CollectionTypes{
		BaseEndpoint[pb.CollectionType]{
			endpointName: EPCollectionTypes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *CollectionTypes) Query(ctx context.Context, query string) ([]*pb.CollectionType, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Collectiontypes, nil
}
