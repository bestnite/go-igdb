package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type CollectionMemberships struct {
	BaseEndpoint[pb.CollectionMembership]
}

func NewCollectionMemberships(request func(URL string, dataBody any) (*resty.Response, error)) *CollectionMemberships {
	a := &CollectionMemberships{
		BaseEndpoint[pb.CollectionMembership]{
			endpointName: EPCollectionMemberships,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *CollectionMemberships) Query(query string) ([]*pb.CollectionMembership, error) {
	resp, err := a.request("https://api.igdb.com/v4/collection_memberships.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CollectionMembershipResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Collectionmemberships) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Collectionmemberships, nil
}
