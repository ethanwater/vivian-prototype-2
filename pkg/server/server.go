package server

import (
	"context"
	"database/sql"
	"sync"
	"vivian/pkg/frontend"

	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ServiceWeaver/weaver"
	"github.com/pelletier/go-toml"
)

type Server struct {
	weaver.Implements[weaver.Main]
	serverControls weaver.Ref[handleInterface]
	listener       weaver.Listener `weaver:"vivian"`

	addr          string
	read_timeout  time.Duration
	write_timeout time.Duration
	mu            sync.Mutex

	db_name  string
	handler  http.Handler
	database *sql.DB
}

func Deploy(ctx context.Context, app *Server) error {
	toml, err := toml.LoadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("[vivian:%s]\n", toml.Get("vivian.version"))
	}

	appHandler := http.NewServeMux()
	app.addr = ":9000"
	app.handler = appHandler
	app.read_timeout = 10 * time.Second
	app.write_timeout = 10 * time.Second
	app.database = EstablishLinkDatabase(ctx, app)
	app.db_name = "users"

	app.Logger(ctx).Info("vivian: app deployed", "address", app.listener)

	appHandler.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(frontend.WebUI))))
	appHandler.Handle("/kill", serverControls{}.kill(ctx, app))
	appHandler.Handle("/add", serverControls{}.add(ctx, app))
	appHandler.Handle("/echo", serverControls{}.echo(ctx, app))
	//appHandler.Handle("/ping", serverControls{}.ping(ctx, app))
	appHandler.HandleFunc(weaver.HealthzURL, weaver.HealthzHandler)

	return http.Serve(app.listener, app.handler)
}