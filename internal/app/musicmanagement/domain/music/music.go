package music

type Music struct{
	Title    string
	ArtistID string
	Version  int
}

func (m *Music) UpdateTitle(newTitle string) {
	m.Title = newTitle
	m.Version++
}