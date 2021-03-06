package handlerfn

import (
	"log"
	"net/http"

	"github.com/chilledoj/realworld-starter-kit/models"
)

// AppEnvironment holds the database connection pool and logger
type AppEnvironment struct {
	DB     *models.AppDB
	Logger *log.Logger
}

// AppHandler is a struct to manage error providing handlers
type AppHandler struct {
	env *AppEnvironment
	fn  func(e *AppEnvironment, w http.ResponseWriter, r *http.Request) error
}

func (ah AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := ctx.Err(); err != nil {
		return
	}
	if e := ah.fn(ah.env, w, r); e != nil {
		ah.env.Logger.Printf("ERROR:: %s: %v", r.URL, e)
		c := http.StatusInternalServerError
		if err, ok := e.(statusCoder); ok {
			c = err.StatusCode()
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(c)
		if err, ok := e.(dataError); ok {
			w.Write(err.Data())
		}
	}
}
