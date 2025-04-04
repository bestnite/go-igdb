package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetPlatformVersionCompanies(query string) ([]*pb.PlatformVersionCompany, error) {
	resp, err := g.Request("https://api.igdb.com/v4/platform_version_companies.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.PlatformVersionCompanyResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Platformversioncompanies) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Platformversioncompanies, nil
}

func (g *igdb) GetPlatformVersionCompanyByID(id uint64) (*pb.PlatformVersionCompany, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	platformVersionCompanies, err := g.GetPlatformVersionCompanies(query)
	if err != nil {
		return nil, err
	}
	return platformVersionCompanies[0], nil
}

func (g *igdb) GetPlatformVersionCompaniesByIDs(ids []uint64) ([]*pb.PlatformVersionCompany, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetPlatformVersionCompanies(idStr)
}
