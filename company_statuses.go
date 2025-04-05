package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCompanyStatuses(query string) ([]*pb.CompanyStatus, error) {
	resp, err := g.Request("https://api.igdb.com/v4/company_statuses.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CompanyStatusResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Companystatuses) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Companystatuses, nil
}

func (g *igdb) GetCompanyStatusByID(id uint64) (*pb.CompanyStatus, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	companyStatuses, err := g.GetCompanyStatuses(query)
	if err != nil {
		return nil, err
	}
	return companyStatuses[0], nil
}

func (g *igdb) GetCompanyStatusesByIDs(ids []uint64) ([]*pb.CompanyStatus, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanyStatuses(idStr)
}

func (g *igdb) GetCompanyStatusesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	companyStatuses, err := g.GetCompanyStatuses(query)
	if err != nil {
		return 0, err
	}
	return int(companyStatuses[0].Id), nil
}
