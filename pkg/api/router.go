package api

import (
	"github.com/breadchris/sifty/gen/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/log"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"time"
)

func NewAPIHandler(assets fs.FS, apiServer api.TwirpServer) http.Handler {
	muxRoot := chi.NewRouter()

	muxRoot.Use(middleware.RequestID)
	muxRoot.Use(middleware.RealIP)
	muxRoot.Use(middleware.Logger)
	//muxRoot.Use(session.Sessioner(session.Options{
	//	Provider:           "file",
	//	CookieName:         "session",
	//	FlashEncryptionKey: "SomethingSuperSecretThatShouldChange",
	//}))

	//muxRoot.Use(middleware.Recoverer)
	muxRoot.Use(middleware.Timeout(time.Second * 5))

	fs := http.FS(assets)
	httpFileServer := http.FileServer(fs)

	muxRoot.Handle("/*", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, apiServer.PathPrefix()) {
			apiServer.ServeHTTP(rw, r)
			return
		}

		filePath := r.URL.Path
		if strings.Index(r.URL.Path, "/") == 0 {
			filePath = r.URL.Path[1:]
		}

		f, err := assets.Open(filePath)
		if os.IsNotExist(err) {
			r.URL.Path = "/"
		}

		if err == nil {
			f.Close()
		}

		log.Warn().Err(err).Str("url", r.URL.Path).Msg("could not open file")

		httpFileServer.ServeHTTP(rw, r)
	}))
	return muxRoot
}
