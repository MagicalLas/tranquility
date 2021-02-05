package adapter

import "magical.dev/tranquility/internal/app/musicmanagement/domain/music"

type MySQLMusicRepository struct{}

func (i MySQLMusicRepository) GetByID(id string) (*music.Music, error) {
	panic("implement me")
}

func (i MySQLMusicRepository) Save(music music.Music) error {
	panic("implement me")
}
