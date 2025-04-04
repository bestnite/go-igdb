package igdb

import (
	"errors"
	"fmt"
	"strings"

	pb "github/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCompanies(query string) ([]*pb.Company, error) {
	resp, err := g.Request("https://api.igdb.com/v4/companies.pb", query)
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

func (g *igdb) GetCompanyByID(id uint64) (*pb.Company, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	companys, err := g.GetCompanies(query)
	if err != nil {
		return nil, err
	}
	return companys[0], nil
}

func (g *igdb) GetCompanyByIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}
