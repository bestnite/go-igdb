package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type PlatformLogos struct {
	BaseEndpoint[pb.PlatformLogo]
}

func NewPlatformLogos(request func(URL string, dataBody any) (*resty.Response, error)) *PlatformLogos {
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
	resp, err := a.request("https://api.igdb.com/v4/platform_logos.pb", query)
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
