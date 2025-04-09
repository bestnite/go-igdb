package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type CollectionMemberships struct {
	BaseEndpoint[pb.CollectionMembership]
}

func NewCollectionMemberships(request RequestFunc) *CollectionMemberships {
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
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
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
