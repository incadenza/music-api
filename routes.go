package main

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/incadenza/music-api/interal/data"
	"github.com/rs/cors"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: app.config.trustedOrigins,
	})

	r.Use(c.Handler)
	r.Get("/tracks", func(w http.ResponseWriter, r *http.Request) {
		tracks, err := app.models.Track.GetAll()

		if err != nil {
			internalErrorResponse(w, r, err)
			return
		}

		successResponse(w, r, response{"tracks": tracks})
	})

	r.Get("/artists/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIDParam(r)

		if err != nil {
			badRequestResponse(w, r, err)
			return
		}

		artist, err := app.models.Artist.GetByID(id)

		if err != nil {
			if errors.Is(err, data.ErrNoRecordFound) {
				notFoundResponse(w, r)
				return
			}

			internalErrorResponse(w, r, err)
			return
		}

		successResponse(w, r, response{"artist": artist})
	})

	r.Get("/artists/{id}/albums", func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIDParam(r)
		if err != nil {
			badRequestResponse(w, r, err)
			return
		}

		albums, err := app.models.Artist.GetAlbums(id)

		if err != nil {
			if errors.Is(err, data.ErrNoRecordFound) {
				notFoundResponse(w, r)
				return
			}

			internalErrorResponse(w, r, err)
			return
		}

		successResponse(w, r, response{"albums": albums})
	})

	r.Get("/albums/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIDParam(r)

		if err != nil {
			badRequestResponse(w, r, err)
			return
		}

		album, err := app.models.Album.GetByID(id)

		if err != nil {
			if errors.Is(err, data.ErrNoRecordFound) {
				notFoundResponse(w, r)
				return
			}

			internalErrorResponse(w, r, err)
			return
		}

		successResponse(w, r, response{"album": album})
	})

	r.Get("/albums/{id}/tracks", func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIDParam(r)
		if err != nil {
			badRequestResponse(w, r, err)
			return
		}

		tracks, err := app.models.Album.GetTracks(id)

		if err != nil {
			if errors.Is(err, data.ErrNoRecordFound) {
				notFoundResponse(w, r)
				return
			}

			internalErrorResponse(w, r, err)
			return
		}

		successResponse(w, r, response{"tracks": tracks})
	})

	return r
}
