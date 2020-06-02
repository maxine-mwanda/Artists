package models

import (
	"api/utils"
	"fmt"
)

func Fetchsingers() (artists []Artist, err error) {
	var artist Artist
	db, err := utils.Connecttodb()
	if err != nil {
		fmt.Println("unable to connect todb")
		return
	}
	query := "select * from artist"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("unable to fetch singers")
		return
	}

	for rows.Next() {
		if err = rows.Scan(&artist.Id, &artist.StageName, &artist.FullName, &artist.DateOfBirth,
			&artist.Nationality, &artist.CreatedAt); err != nil {
			fmt.Println("Unable to scan because ", err)
			continue
		}
		artists = append(artists, artist)
	}
	return
}

func Create_artist(artist NewArtist) (id int64, err error) {
	query := "insert into artist (stage_name, full_name, date_of_birth, nationality)" +
		"values (?,?,?,?)"
	db, err := utils.Connecttodb()
	if err != nil {
		fmt.Println("unable to connect todb")
		return
	}
	row, err := db.Exec(query, artist.StageName, artist.FullName, artist.DateOfBirth, artist.Nationality)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, _ = row.LastInsertId()
	return

}

func Fetchsinger(singer string) (artist Artist) {
	db, err := utils.Connecttodb()
	if err != nil {
		fmt.Println("unable to connect todb")
		return
	}
	query := "select * from artist where stage_name = ?;"
	row := db.QueryRow(query, singer)
	if row == nil {
		fmt.Println("unable to fetch singers")
		return
	}

	err = row.Scan(&artist.Id, &artist.StageName, &artist.FullName, &artist.DateOfBirth,
		&artist.Nationality, &artist.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
