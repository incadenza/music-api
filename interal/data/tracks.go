package data

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Track struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre,omitempty"`
	GenreID  int64  `json:"genreId,omitempty"`
	AlbumID  int64  `json:"albumId,omitempty"`
	Album    string `json:"album,omitempty"`
	ArtistID int64  `json:"artistId,omitempty"`
	Artist   string `json:"artist,omitempty"`
}

type TrackRepo struct {
	DB *pgxpool.Pool
}

func (repo TrackRepo) GetAll() ([]*Track, error) {
	stmt := `
	SELECT
		tracks.id,
		tracks.name,
		tracks.genre_id,
		genres.name,
		tracks.album_id,
		albums.title,
		albums.artist_id,
		artists.name
	FROM
		tracks
		JOIN genres ON genres.id = tracks.genre_id
		JOIN albums ON albums.id = tracks.album_id
		JOIN artists ON artists.id = albums.artist_id
	`
	rows, err := repo.DB.Query(context.Background(), stmt)

	tracks := []*Track{}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		track := Track{}
		if err := rows.Scan(&track.ID, &track.Name, &track.GenreID, &track.Genre, &track.AlbumID, &track.Album, &track.ArtistID, &track.Artist); err != nil {
			return nil, err
		}
		tracks = append(tracks, &track)
	}

	return tracks, nil
}
