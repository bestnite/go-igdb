package igdb

import (
	"errors"
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetCompanies(query string) ([]*pb.Company, error) {
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

func (g *Client) GetCompanyByID(id uint64) (*pb.Company, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	companys, err := g.GetCompanies(query)
	if err != nil {
		return nil, err
	}
	return companys[0], nil
}

func (g *Client) GetCompanyByIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *Client) GetCompanyByChangeDateFormatID(id uint64) ([]*pb.Company, error) {
	query := fmt.Sprintf(`where change_date_format = %d; fields *;`, id)
	return g.GetCompanies(query)
}

func (g *Client) GetCompanyByChangeDateFormatsIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where change_date_format = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *Client) GetCompanyByChangedCompanyID(id uint64) ([]*pb.Company, error) {
	query := fmt.Sprintf(`where changed_company_id = %d; fields *;`, id)
	return g.GetCompanies(query)
}

func (g *Client) GetCompanyByChangedCompanyIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where changed_company_id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *Client) GetCompanyByLogoID(id uint64) ([]*pb.Company, error) {
	query := fmt.Sprintf(`where logo = %d; fields *;`, id)
	return g.GetCompanies(query)
}

func (g *Client) GetCompanyByLogoIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where logo = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *Client) GetCompanyByParentID(id uint64) ([]*pb.Company, error) {
	query := fmt.Sprintf(`where parent = %d; fields *;`, id)
	return g.GetCompanies(query)
}

func (g *Client) GetCompanyByParentIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where parent = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *Client) GetCompanyByStartDateFormatID(id uint64) ([]*pb.Company, error) {
	query := fmt.Sprintf(`where start_date_format = %d; fields *;`, id)
	return g.GetCompanies(query)
}

func (g *Client) GetCompanyByStartDateFormatsIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where start_date_format = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *Client) GetCompanyByStatusID(id uint64) ([]*pb.Company, error) {
	query := fmt.Sprintf(`where status = %d; fields *;`, id)
	return g.GetCompanies(query)
}

func (g *Client) GetCompanyByStatusIDs(ids []uint64) ([]*pb.Company, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where status = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanies(idStr)
}

func (g *Client) GetCompaniesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	companies, err := g.GetCompanies(query)
	if err != nil {
		return 0, err
	}
	return int(companies[0].Id), nil
}
