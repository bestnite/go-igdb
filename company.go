package igdb

import (
	"errors"
	"fmt"
	"github/bestnite/go-igdb/constant"

	pb "github/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetIGDBCompany(id uint64) (*pb.Company, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	resp, err := g.Request(constant.IGDBCompaniesURL, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch IGDB company for ID %d: %w", id, err)
	}

	var data pb.CompanyResult
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal IGDB companies response: %w", err)
	}

	if len(data.Companies) == 0 {
		return nil, errors.New("company not found")
	}

	return data.Companies[0], nil
}
