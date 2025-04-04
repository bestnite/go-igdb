package igdb

import (
	"fmt"
	pb "github/bestnite/go-igdb/proto"
	"strings"

	"google.golang.org/protobuf/proto"
)

func (g *igdb) GetGenres(query string) ([]*pb.Genre, error) {
	resp, err := g.Request("https://api.igdb.com/v4/genres.pb", query)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}

	data := pb.GenreResult{}
	if err = proto.Unmarshal(resp.Body(), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	if len(data.Genres) == 0 {
		return nil, fmt.Errorf("no results: %s", query)
	}

	return data.Genres, nil
}

func (g *igdb) GetGenreByID(id uint64) (*pb.Genre, error) {
	query := fmt.Sprintf(`where id=%d; fields *;`, id)
	genres, err := g.GetGenres(query)
	if err != nil {
		return nil, err
	}
	return genres[0], nil
}

func (g *igdb) GetGenresByIDs(ids []uint64) ([]*pb.Genre, error) {
	idStrSlice := make([]string, len(ids))
	for i, id := range ids {
		idStrSlice[i] = fmt.Sprintf("%d", id)
	}

	idStr := fmt.Sprintf(`where id = (%s); fields *;`, strings.Join(idStrSlice, ","))

	return g.GetGenres(idStr)
}
