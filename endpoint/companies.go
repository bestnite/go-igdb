package endpoint

import (
	"errors"
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"
	"github.com/go-resty/resty/v2"

	"google.golang.org/protobuf/proto"
)

type Companies struct {
	BaseEndpoint[pb.Company]
}

func NewCompanies(request func(URL string, dataBody any) (*resty.Response, error)) *Companies {
	a := &Companies{
		BaseEndpoint[pb.Company]{
			endpointName: EPCompanies,
			request:      request,
		},
	}
	a.queryFunc = a.Query
	return a
}

func (a *Companies) Query(query string) ([]*pb.Company, error) {
	resp, err := a.request(fmt.Sprintf("https://api.igdb.com/v4/%s.pb", a.endpointName), query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CompanyResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Companies) == 0 {
		return nil, errors.New("no results")
	}

	return data.Companies, nil
}
