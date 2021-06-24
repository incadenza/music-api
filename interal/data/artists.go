package data

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Artist struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ArtistRepo struct {
	DB *pgxpool.Pool
}

func (repo ArtistRepo) GetByID(id int64) (*Artist, error) {
	stmt := `SELECT id, name FROM artists WHERE id = $1`

	artist := Artist{}
	err := repo.DB.QueryRow(context.Background(), stmt, id).Scan(&artist.ID, &artist.Name)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecordFound

		}
		return nil, err
	}

	return &artist, nil
}

func (repo ArtistRepo) GetAlbums(id int64) ([]*Album, error) {
	stmt := `SELECT id, title FROM albums WHERE artist_id = $1`

	if _, err := repo.GetByID(id); err != nil {
		return nil, err
	}

	albums := []*Album{}
	rows, err := repo.DB.Query(context.Background(), stmt, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		album := Album{}
		if err := rows.Scan(&album.ID, &album.Title); err != nil {
			return nil, err
		}
		albums = append(albums, &album)
	}

	return albums, nil
}
