package data

import (
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrNoRecordFound = errors.New("no record found")
)

type Models struct {
	Track  TrackRepo
	Artist ArtistRepo
	Album  AlbumRepo
}

func NewModels(db *pgxpool.Pool) *Models {
	return &Models{
		Track:  TrackRepo{DB: db},
		Artist: ArtistRepo{DB: db},
		Album:  AlbumRepo{DB: db},
	}
}
