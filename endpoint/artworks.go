package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Artworks struct {
	BaseEndpoint[pb.Artwork]
}

func NewArtworks(request RequestFunc) *Artworks {
	a := &Artworks{
		BaseEndpoint[pb.Artwork]{
			endpointName: EPArtworks,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.ArtworkResult) []*pb.Artwork { return r.Artworks })
	return a
}
