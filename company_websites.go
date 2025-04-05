package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetCompanyWebsites(query string) ([]*pb.CompanyWebsite, error) {
	resp, err := g.Request("https://api.igdb.com/v4/company_websites.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.CompanyWebsiteResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Companywebsites) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Companywebsites, nil
}

func (g *Client) GetCompanyWebsiteByID(id uint64) (*pb.CompanyWebsite, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	companyWebsites, err := g.GetCompanyWebsites(query)
	if err != nil {
		return nil, err
	}
	return companyWebsites[0], nil
}

func (g *Client) GetCompanyWebsitesByIDs(ids []uint64) ([]*pb.CompanyWebsite, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanyWebsites(idStr)
}

func (g *Client) GetCompanyWebsitesByTypeID(id uint64) ([]*pb.CompanyWebsite, error) {
	query := fmt.Sprintf(`where type = %d; fields *;`, id)
	return g.GetCompanyWebsites(query)
}

func (g *Client) GetCompanyWebsitesByTypeIDs(ids []uint64) ([]*pb.CompanyWebsite, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanyWebsites(idStr)
}

func (g *Client) GetCompanyWebsitesLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	companyWebsites, err := g.GetCompanyWebsites(query)
	if err != nil {
		return 0, err
	}
	return int(companyWebsites[0].Id), nil
}
