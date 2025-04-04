package igdb

import (
	"fmt"
)

func (g *igdb) GetParentGameID(id uint64) (uint64, error) {
	detail, err := g.GetGameByID(id)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch IGDB app detail for parent: %d: %w", id, err)
	}
	hasParent := false
	if detail.ParentGame != nil && detail.ParentGame.Id != 0 {
		hasParent = true
		detail, err = g.GetGameByID(detail.ParentGame.Id)
		if err != nil {
			return 0, fmt.Errorf("failed to fetch IGDB version parent: %d: %w", detail.VersionParent.Id, err)
		}
	}
	for detail.VersionParent != nil && detail.VersionParent.Id != 0 {
		hasParent = true
		detail, err = g.GetGameByID(detail.VersionParent.Id)
		if err != nil {
			return 0, fmt.Errorf("failed to fetch IGDB version parent: %d: %w", detail.VersionParent.Id, err)
		}
	}

	if hasParent {
		return detail.Id, nil
	}

	return id, nil
}
