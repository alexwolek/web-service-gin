package models

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// returns all albums in the collection
func GetAllAlbums() []Album {
	return albums
}

// GetAlbumByID locates the album whose ID value matches the id
func GetAlbumByID(id string) *Album {
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			return &a
		}
	}

	return nil
}

// adds a new album to the collection
func AddNewAlbum(newAlbum Album) {
	albums = append(albums, newAlbum)
}
