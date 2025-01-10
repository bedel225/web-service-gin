package models

import "github/bedel225/db"

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func (e Album) Save() {
	Albums = append(Albums, e)
}

func GetAllAlbum() ([]Album, error) {
	query := "SELECT * FROM albums"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	var albums []Album
	for rows.Next() {
		var album Album
		rows.Scan(album.ID, album.Title, album.Artist, album.Price)
		albums = append(albums, album)
	}
	return albums, nil

}
