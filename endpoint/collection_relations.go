package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

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

func (a *CollectionRelations) Query(query string) ([]*pb.CollectionRelation, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionRelationResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Collectionrelations) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Collectionrelations, nil
}
