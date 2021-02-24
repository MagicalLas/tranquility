package application

import (
	"context"
	"magical.dev/tranquility/internal/app/musicmanagement/application/dto"
	"magical.dev/tranquility/internal/app/musicmanagement/domain/music"
	"magical.dev/tranquility/internal/app/musicmanagement/domain/vocaloid"
)

type MusicManagementApplicationService struct{
	musicRepository    music.Repository
	vocaloidRepository vocaloid.Repository
}

func NewMusicApplicationService() *MusicManagementApplicationService {
	return &MusicManagementApplicationService{}
}

func (s MusicManagementApplicationService) PresentNewMusic(ctx context.Context, artistID string, title string) {

}

func (s MusicManagementApplicationService) PresentNewMusicUsingVocaloid(ctx context.Context, musicID string, title string, vocaloidID string) {

}

func (s MusicManagementApplicationService) UpdateMusicTitle(ctx context.Context, musicID string, title string) error {
	updateTargetMusic, err := s.musicRepository.GetByID(musicID)
	if err != nil {
		return err
	}
	updateTargetMusic.UpdateTitle(title)
	return s.musicRepository.Save(updateTargetMusic)
}

func (s MusicManagementApplicationService) ShowMyAllMusic(ctx context.Context, artistID string) ([]*dto.Music, error) {
	musicList, err := s.musicRepository.GetMusicsByArtistID(artistID)
	if err != nil {
		return nil, err
	}
	return dto.TransferListToDTOList(musicList), nil
}
