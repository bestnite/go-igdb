package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type Themes struct {
	BaseEndpoint[pb.Theme]
}

func NewThemes(request RequestFunc) *Themes {
	a := &Themes{
		BaseEndpoint[pb.Theme]{
			endpointName: EPThemes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.ThemeResult) []*pb.Theme { return r.Themes })
	return a
}
