package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type CollectionMembershipTypes struct {
	BaseEndpoint[pb.CollectionMembershipType]
}

func NewCollectionMembershipTypes(request func(URL string, dataBody any) (*resty.Response, error)) *CollectionMembershipTypes {
	a := &CollectionMembershipTypes{
		BaseEndpoint[pb.CollectionMembershipType]{
			endpointName: EPCollectionMembershipTypes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *CollectionMembershipTypes) Query(query string) ([]*pb.CollectionMembershipType, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionMembershipTypeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Collectionmembershiptypes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Collectionmembershiptypes, nil
}
