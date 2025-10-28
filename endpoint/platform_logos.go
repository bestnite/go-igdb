package endpoint

import (
	"context"
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type PlatformLogos struct {
	BaseEndpoint[pb.PlatformLogo]
}

func NewPlatformLogos(request RequestFunc) *PlatformLogos {
	a := &PlatformLogos{
		BaseEndpoint: BaseEndpoint[pb.PlatformLogo]{
			endpointName: EPPlatformLogos,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlatformLogos) Query(ctx context.Context, query string) ([]*pb.PlatformLogo, error) {
	resp, err := a.request(ctx, "POST", fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformLogoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return data.Platformlogos, nil
}
