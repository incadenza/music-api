package data

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Album struct {
	ID    int64  `json:"id"`
	Title string `json:"name"`
}

type AlbumRepo struct {
	DB *pgxpool.Pool
}

func (repo AlbumRepo) GetByID(id int64) (*Album, error) {
	stmt := `SELECT id, title FROM albums WHERE id = $1`

	album := Album{}
	err := repo.DB.QueryRow(context.Background(), stmt, id).Scan(&album.ID, &album.Title)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecordFound

		}
		return nil, err
	}

	return &album, nil
}

func (repo AlbumRepo) GetTracks(id int64) ([]*Track, error) {
	stmt := `SELECT id, name FROM tracks WHERE album_id = $1`

	if _, err := repo.GetByID(id); err != nil {
		return nil, err
	}

	tracks := []*Track{}
	rows, err := repo.DB.Query(context.Background(), stmt, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		track := Track{}
		if err := rows.Scan(&track.ID, &track.Name); err != nil {
			return nil, err
		}
		tracks = append(tracks, &track)
	}

	return tracks, nil
}
