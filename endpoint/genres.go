package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Genres struct {
	BaseEndpoint[pb.Genre]
}

func NewGenres(request RequestFunc) *Genres {
	a := &Genres{
		BaseEndpoint[pb.Genre]{
			endpointName: EPGenres,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.GenreResult) []*pb.Genre { return r.Genres })
	return a
}
