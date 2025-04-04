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

func (g *igdb) GetCompanyByChangeDateFormatID(id uint64) ([]*pb.Company, error) {
	query := fmt.Sprintf(`where change_date_format = %d; fields *;`, id)
	return g.GetCompanies(query)
}

func (g *igdb) GetCompanyByChangeDateFormatsIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where change_date_format = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *igdb) GetCompanyByChangedCompanyID(id uint64) ([]*pb.Company, error) {
	query := fmt.Sprintf(`where changed_company_id = %d; fields *;`, id)
	return g.GetCompanies(query)
}

func (g *igdb) GetCompanyByChangedCompanyIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where changed_company_id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *igdb) GetCompanyByLogoID(id uint64) ([]*pb.Company, error) {
	query := fmt.Sprintf(`where logo = %d; fields *;`, id)
	return g.GetCompanies(query)
}

func (g *igdb) GetCompanyByLogoIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where logo = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *igdb) GetCompanyByParentID(id uint64) ([]*pb.Company, error) {
	query := fmt.Sprintf(`where parent = %d; fields *;`, id)
	return g.GetCompanies(query)
}

func (g *igdb) GetCompanyByParentIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where parent = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *igdb) GetCompanyByStartDateFormatID(id uint64) ([]*pb.Company, error) {
	query := fmt.Sprintf(`where start_date_format = %d; fields *;`, id)
	return g.GetCompanies(query)
}

func (g *igdb) GetCompanyByStartDateFormatsIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where start_date_format = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *igdb) GetCompanyByStatusID(id uint64) ([]*pb.Company, error) {
	query := fmt.Sprintf(`where status = %d; fields *;`, id)
	return g.GetCompanies(query)
}

func (g *igdb) GetCompanyByStatusIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where status = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *igdb) GetCompaniesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	companies, err := g.GetCompanies(query)
	if err != nil {
		return 0, err
	}
	return int(companies[0].Id), nil
}
