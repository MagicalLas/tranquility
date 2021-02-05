package application

import (
	"context"
	"magical.dev/tranquility/internal/app/recommendation/domain"
)

type MusicRecommendationService struct {}

func (s *MusicRecommendationService) RecommendMusic(ctx context.Context, userID string) []domain.Music {
	return nil
}
