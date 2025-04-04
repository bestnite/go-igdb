package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetInvolvedCompanies(query string) ([]*pb.InvolvedCompany, error) {
	resp, err := g.Request("https://api.igdb.com/v4/involved_companies.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.InvolvedCompanyResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Involvedcompanies) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Involvedcompanies, nil
}

func (g *igdb) GetInvolvedCompanyByID(id uint64) (*pb.InvolvedCompany, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	involvedCompanies, err := g.GetInvolvedCompanies(query)
	if err != nil {
		return nil, err
	}
	return involvedCompanies[0], nil
}

func (g *igdb) GetInvolvedCompaniesByIDs(ids []uint64) ([]*pb.InvolvedCompany, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetInvolvedCompanies(idStr)
}

func (g *igdb) GetInvolvedCompaniesByGameID(id uint64) ([]*pb.InvolvedCompany, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetInvolvedCompanies(query)
}

func (g *igdb) GetInvolvedCompaniesByGameIDs(ids []uint64) ([]*pb.InvolvedCompany, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetInvolvedCompanies(idStr)
}

func (g *igdb) GetInvolvedCompaniesByCompanyID(id uint64) ([]*pb.InvolvedCompany, error) {
	query := fmt.Sprintf(`where company = %d; fields *;`, id)
	return g.GetInvolvedCompanies(query)
}

func (g *igdb) GetInvolvedCompaniesByCompanyIDs(ids []uint64) ([]*pb.InvolvedCompany, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where company = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetInvolvedCompanies(idStr)
}

func (g *igdb) GetInvolvedCompaniesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	involvedCompanies, err := g.GetInvolvedCompanies(query)
	if err != nil {
		return 0, err
	}
	return int(involvedCompanies[0].Id), nil
}
