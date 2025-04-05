package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type CollectionMembershipTypes struct{ BaseEndpoint }

func (a *CollectionMembershipTypes) Query(query string) ([]*pb.CollectionMembershipType, error) {
	resp, err := a.request("https://api.igdb.com/v4/collection_membership_types.pb", query)
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
