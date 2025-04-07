package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type CollectionRelationTypes struct {
	BaseEndpoint[pb.CollectionRelationType]
}

func NewCollectionRelationTypes(request func(URL string, dataBody any) (*resty.Response, error)) *CollectionRelationTypes {
	a := &CollectionRelationTypes{
		BaseEndpoint[pb.CollectionRelationType]{
			endpointName: EPCollectionRelationTypes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *CollectionRelationTypes) Query(query string) ([]*pb.CollectionRelationType, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionRelationTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Collectionrelationtypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Collectionrelationtypes, nil
}
