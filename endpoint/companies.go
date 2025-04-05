package endpoint

import (
	"errors"
	"fmt"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

type Companies struct{ BaseEndpoint }

func (a *Companies) Query(query string) ([]*pb.Company, error) {
	resp, err := a.request("https://api.igdb.com/v4/companies.pb", query)
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
