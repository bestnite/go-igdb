package endpoint

import (
	pb "git.nite07.com/nite/go-igdb/proto"
)

type WebsiteTypes struct {
	BaseEndpoint[pb.WebsiteType]
}

func NewWebsiteTypes(request RequestFunc) *WebsiteTypes {
	a := &WebsiteTypes{
		BaseEndpoint[pb.WebsiteType]{
			endpointName: EPWebsiteTypes,
			request:      request,
		},
	}
	a.queryFunc = a.queryPB(func(r *pb.WebsiteTypeResult) []*pb.WebsiteType { return r.Websitetypes })
	return a
}
