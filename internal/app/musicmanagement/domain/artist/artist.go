package artist

import "magical.dev/tranquility/internal/app/musicmanagement/domain/music"

type Artist struct {
	id string
}

func (a *Artist) CreateMusic(title string) *music.Music {
	return &music.Music{
		Title:    title,
		ArtistID: a.id,
		Version:  0,
	}
}