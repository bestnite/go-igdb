package igdb

import (
	"fmt"
	"strings"

	pb "github/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetArtworks(query string) ([]*pb.Artwork, error) {
	resp, err := g.Request("https://api.igdb.com/v4/artworks.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.ArtworkResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Artworks) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Artworks, nil
}

func (g *igdb) GetArtworkByID(id uint64) (*pb.Artwork, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	artworks, err := g.GetArtworks(query)
	if err != nil {
		return nil, err
	}
	return artworks[0], nil
}

func (g *igdb) GetArtworksByIDs(ids []uint64) ([]*pb.Artwork, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetArtworks(idStr)
}

func (g *igdb) GetArtworksByGameID(id uint64) ([]*pb.Artwork, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetArtworks(query)
}

func (g *igdb) GetArtworksByGameIDs(ids []uint64) ([]*pb.Artwork, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetArtworks(idStr)
}
