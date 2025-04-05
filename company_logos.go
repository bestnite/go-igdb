package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetCompanyLogos(query string) ([]*pb.CompanyLogo, error) {
	resp, err := g.Request("https://api.igdb.com/v4/company_logos.pb", query)
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

func (g *Client) GetCompanyLogoByID(id uint64) (*pb.CompanyLogo, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	companyLogos, err := g.GetCompanyLogos(query)
	if err != nil {
		return nil, err
	}
	return companyLogos[0], nil
}

func (g *Client) GetCompanyLogosByIDs(ids []uint64) ([]*pb.CompanyLogo, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetCompanyLogos(idStr)
}

func (g *Client) GetCompanyLogosLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	companyLogos, err := g.GetCompanyLogos(query)
	if err != nil {
		return 0, err
	}
	return int(companyLogos[0].Id), nil
}
