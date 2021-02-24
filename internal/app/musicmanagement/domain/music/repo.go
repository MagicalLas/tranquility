package music

type Repository interface {
	GetByID(id string) (*Music, error)
	GetMusicsByArtistID(id string) ([]*Music, error)
	Save(music *Music) error
}
