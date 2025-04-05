package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type PlatformVersionCompanies struct {
	BaseEndpoint[pb.PlatformVersionCompany]
}

func NewPlatformVersionCompanies(request func(URL string, dataBody any) (*resty.Response, error)) *PlatformVersionCompanies {
	a := &PlatformVersionCompanies{
		BaseEndpoint[pb.PlatformVersionCompany]{
			endpointName: EPPlatformVersionCompanies,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *PlatformVersionCompanies) Query(query string) ([]*pb.PlatformVersionCompany, error) {
	resp, err := a.request("https://api.igdb.com/v4/platform_version_companies.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformVersionCompanyResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformversioncompanies) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformversioncompanies, nil
}
