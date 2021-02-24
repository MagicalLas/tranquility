package vocaloid

type Repository interface {
	GetByID(id string) (*Vocaloid, error)
	Save(music Vocaloid) error
}
