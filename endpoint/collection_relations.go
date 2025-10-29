package endpoint

import (
	"context"
	"fmt"

	pb "git.nite07.com/nite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type CollectionRelations struct {
	BaseEndpoint[pb.CollectionRelation]
}

func NewCollectionRelations(request RequestFunc) *CollectionRelations {
	a := &CollectionRelations{
		BaseEndpoint[pb.CollectionRelation]{
			endpointName: EPCollectionRelations,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *CollectionRelations) Query(ctx context.Context, query string) ([]*pb.CollectionRelation, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionRelationResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Collectionrelations, nil
}
