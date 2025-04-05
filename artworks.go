package igdb

import (
	"fmt"
	"strings"

	pb "github.com/bestnite/go-igdb/proto"

	"google.golang.org/protobuf/proto"
)

func (g *Client) GetArtworks(query string) ([]*pb.Artwork, error) {
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

func (g *Client) GetArtworkByID(id uint64) (*pb.Artwork, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	artworks, err := g.GetArtworks(query)
	if err != nil {
		return nil, err
	}
	return artworks[0], nil
}

func (g *Client) GetArtworksByIDs(ids []uint64) ([]*pb.Artwork, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetArtworks(idStr)
}

func (g *Client) GetArtworksByGameID(id uint64) ([]*pb.Artwork, error) {
	query := fmt.Sprintf(`where game = %d; fields *;`, id)
	return g.GetArtworks(query)
}

func (g *Client) GetArtworksByGameIDs(ids []uint64) ([]*pb.Artwork, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where game = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetArtworks(idStr)
}

func (g *Client) GetArtworksLength() (int, error) {
	query := `fields *; sort id desc; limit 1;`
	artworks, err := g.GetArtworks(query)
	if err != nil {
		return 0, err
	}
	return int(artworks[0].Id), nil
}
