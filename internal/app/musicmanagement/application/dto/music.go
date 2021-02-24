package dto

import "magical.dev/tranquility/internal/app/musicmanagement/domain/music"

type Music struct{}

func TransferToDTO(music *music.Music) *Music {
	return nil
}

func TransferListToDTOList(musicList []*music.Music) []*Music {
	return nil
}
