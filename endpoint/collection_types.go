package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type CollectionTypes struct {
	BaseEndpoint[pb.CollectionType]
}

func NewCollectionTypes(request func(URL string, dataBody any) (*resty.Response, error)) *CollectionTypes {
	a := &CollectionTypes{
		BaseEndpoint[pb.CollectionType]{
			endpointName: EPCollectionTypes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *CollectionTypes) Query(query string) ([]*pb.CollectionType, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Collectiontypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Collectiontypes, nil
}
