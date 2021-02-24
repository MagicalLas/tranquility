package presentation

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"magical.dev/tranquility/internal/app/recommendation/application"
)

type EndpointSet struct {
	RecommendMusicEndpoint endpoint.Endpoint
}

func CreateEndpointSet(service application.MusicRecommendationService) EndpointSet {
	var recommendMusicEndpoint endpoint.Endpoint
	{
		recommendMusicEndpoint = MakePresentMusicEndpoint(service)
		// Sum is limited to 1 request per second with burst of 1 request.
		// Note, rate is defined as a time interval between requests.
	}
	return EndpointSet{
		RecommendMusicEndpoint: recommendMusicEndpoint,
	}
}

func MakePresentMusicEndpoint(s application.MusicRecommendationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RecommendMusicRequest)
		s.RecommendMusic(ctx, req.UserID)
		return RecommendMusicResponse{}, nil
	}
}
