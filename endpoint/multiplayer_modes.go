package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type MultiplayerModes struct {
	BaseEndpoint[pb.MultiplayerMode]
}

func NewMultiplayerModes(request RequestFunc) *MultiplayerModes {
	a := &MultiplayerModes{
		BaseEndpoint[pb.MultiplayerMode]{
			endpointName: EPMultiplayerModes,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *MultiplayerModes) Query(query string) ([]*pb.MultiplayerMode, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.MultiplayerModeResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Multiplayermodes) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Multiplayermodes, nil
}
