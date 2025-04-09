package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type PlatformLogos struct {
	BaseEndpoint[pb.PlatformLogo]
}

func NewPlatformLogos(request RequestFunc) *PlatformLogos {
	a := &PlatformLogos{
		BaseEndpoint[pb.PlatformLogo]{
			endpointName: EPPlatformLogos,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlatformLogos) Query(query string) ([]*pb.PlatformLogo, error) {
	resp, err := a.request("POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformLogoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformlogos) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformlogos, nil
}
