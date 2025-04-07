package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type GameVideos struct {
	BaseEndpoint[pb.GameVideo]
}

func NewGameVideos(request func(URL string, dataBody any) (*resty.Response, error)) *GameVideos {
	a := &GameVideos{
		BaseEndpoint[pb.GameVideo]{
			endpointName: EPGameVideos,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *GameVideos) Query(query string) ([]*pb.GameVideo, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GameVideoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Gamevideos) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Gamevideos, nil
}
