package endpoint

import (
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type CompanyLogos struct{ BaseEndpoint }

func (a *CompanyLogos) Query(query string) ([]*pb.CompanyLogo, error) {
	resp, err := a.request("https://api.igdb.com/v4/company_logos.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CompanyLogoResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Companylogos) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Companylogos, nil
}
