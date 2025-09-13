package handlers

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/mtlynch/picoshare/v2/garbagecollect"
	"github.com/mtlynch/picoshare/v2/space"
)

type (
	SpaceChecker interface {
		Check() (space.Usage, error)
	}

	Clock interface {
		Now() time.Time
	}

	Authenticator interface {
		StartSession(w http.ResponseWriter, r *http.Request)
		ClearSession(w http.ResponseWriter)
		Authenticate(r *http.Request) bool
	}

	Server struct {
		router        *mux.Router
		rootRouter    *mux.Router
		authenticator Authenticator
		store         Store
		spaceChecker  SpaceChecker
		collector     *garbagecollect.Collector
		clock         Clock
		basePath      string
	}
)

// Router returns the underlying router interface for the server.
func (s Server) Router() *mux.Router {
	if s.rootRouter != nil {
		return s.rootRouter
	}
	return s.router
}

// New creates a new server with all the state it needs to satisfy HTTP
// requests.
func New(authenticator Authenticator, store Store, spaceChecker SpaceChecker, collector *garbagecollect.Collector, clock Clock) Server {
	s := Server{
		router:        mux.NewRouter(),
		authenticator: authenticator,
		store:         store,
		spaceChecker:  spaceChecker,
		collector:     collector,
		clock:         clock,
		basePath:      normalizeBasePathFromEnv(),
	}

	// If a base path is configured, mount a subrouter at that path and use it
	// for route registration, while keeping the root router as the public entry.
	if s.basePath != "/" {
		root := mux.NewRouter()
		root.Use(withBasePath(s.basePath))
		s.rootRouter = root
		s.router = root.PathPrefix(s.basePath).Subrouter()
		// Redirect the base path without trailing slash to include it, so
		// http://host/prefix -> http://host/prefix/
		s.rootRouter.HandleFunc(s.basePath, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, s.basePath+"/", http.StatusMovedPermanently)
		}).Methods(http.MethodGet)
	} else {
		// Expose base path to all requests via context.
		s.router.Use(withBasePath(s.basePath))
		s.rootRouter = s.router
	}

	s.routes()
	return s
}
