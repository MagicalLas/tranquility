package presentation

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"magical.dev/tranquility/internal/app/musicmanagement/application"
)

type EndpointSet struct {
	PresentMusicEndpoint endpoint.Endpoint
}

func CreateEndpointSet(service application.MusicManagementApplicationService) EndpointSet {
	var presentMusicEndpoint endpoint.Endpoint
	{
		presentMusicEndpoint = MakePresentMusicEndpoint(service)
		// Sum is limited to 1 request per second with burst of 1 request.
		// Note, rate is defined as a time interval between requests.
	}
	return EndpointSet{
		PresentMusicEndpoint: presentMusicEndpoint,
	}
}

func MakePresentMusicEndpoint(s application.MusicManagementApplicationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PresentMusicRequest)
		s.PresentNewMusic(ctx, req.ArtistID, req.Title)
		return PresentMusicResponse{}, nil
	}
}
