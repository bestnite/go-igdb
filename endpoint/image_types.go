package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type ImageTypes struct {
	BaseEndpoint[pb.ImageType]
}

func NewImageTypes(request RequestFunc) *ImageTypes {
	a := &ImageTypes{
		BaseEndpoint[pb.ImageType]{
			endpointName: EPImageTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.ImageTypeResult) []*pb.ImageType { return r.Imagetypes })
	return a
}
