package models

import (
	"github/bedel225/db"
)

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func GetAllAlbum() ([]Album, error) {
	query := "SELECT * FROM albums"
	rows, err := db.DB.Query(query)

	if err != nil {
		panic(err)
		//return nil, err
	}

	var albums []Album
	for rows.Next() {
		var album Album
		rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		albums = append(albums, album)
	}
	return albums, nil

}

func (a *Album) Save() {
	query := `
		INSERT INTO albums(title, artist, price)
		VALUES (?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(a.Title, a.Artist, a.Price)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	a.ID = id
}
