package adapter

import "magical.dev/tranquility/internal/app/musicmanagement/domain/music"

type InMemoryMusicRepository struct{}

func (i InMemoryMusicRepository) GetByID(id string) (*music.Music, error) {
	panic("implement me")
}

func (i InMemoryMusicRepository) Save(music music.Music) error {
	panic("implement me")
}
