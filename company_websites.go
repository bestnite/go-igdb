package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetCompanyWebsites(query string) ([]*pb.CompanyWebsite, error) {
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

func (g *igdb) GetCompanyWebsiteByID(id uint64) (*pb.CompanyWebsite, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	companyWebsites, err := g.GetCompanyWebsites(query)
	if err != nil {
		return nil, err
	}
	return companyWebsites[0], nil
}

func (g *igdb) GetCompanyWebsitesByIDs(ids []uint64) ([]*pb.CompanyWebsite, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanyWebsites(idStr)
}

func (g *igdb) GetCompanyWebsitesByTypeID(id uint64) ([]*pb.CompanyWebsite, error) {
	query := fmt.Sprintf(`where type = %d; fields *;`, id)
	return g.GetCompanyWebsites(query)
}

func (g *igdb) GetCompanyWebsitesByTypeIDs(ids []uint64) ([]*pb.CompanyWebsite, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where type = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanyWebsites(idStr)
}
