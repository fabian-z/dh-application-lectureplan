package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/crewjam/saml/samlsp"
	"github.com/fabian-z/dh-application-lectureplan/backend/template"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/acme/autocert"
)

var (
	// Globals must be either:
	// already concurrency safe, protected with locks or not be changed after init
	templates           *template.Templates
	db                  *sqlx.DB
	executableDirectory string

	useTLS        = false
	useSSO        = false
	listeningAddr = ":8080"
)

func main() {

	// Setup globals

	var err error
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	executablePath, err := os.Executable()
	if err != nil {
		log.Fatal("error getting executable path: ", err)
	}
	executableDirectory = filepath.Dir(executablePath)

	templates = new(template.Templates)
	err = templates.Init(filepath.Join(executableDirectory, "template", "html"))
	if err != nil {
		log.Fatal(err)
	}

	// Setup database
	db = openDB()

	ssoHandler := emptyHandler
	var ssoMiddleware *samlsp.Middleware
	// Setup SSO Handler if enable
	if useSSO {
		ssoMiddleware, err = loadSSO()

		if err != nil {
			log.Fatal("error initializing SAML handling:", err)
		}

		ssoHandler = ssoMiddleware.RequireAccount
	}

	// Setup router and HTTP server
	router := chi.NewRouter()

	// Set security headers
	router.Use(middleware.SetHeader("Content-Security-Policy", "default-src 'none'; script-src 'self'; font-src 'self'; connect-src 'self' wss: ws:; img-src 'self'; style-src 'self' 'unsafe-inline';"))
	router.Use(middleware.SetHeader("X-Frame-Options", "deny"))
	router.Use(middleware.SetHeader("X-XSS-Protection", "1; mode=block"))
	router.Use(middleware.SetHeader("X-Content-Type-Options", "nosniff"))
	// Feature policy would be nice, but very long https://github.com/w3c/webappsec-permissions-policy/issues/189

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	var staticPath = filepath.Join(executableDirectory, "static")
	router.Get("/css/*", http.FileServer(http.Dir(staticPath)).ServeHTTP)
	router.Get("/js/*", http.FileServer(http.Dir(staticPath)).ServeHTTP)
	router.Get("/gfx/*", http.FileServer(http.Dir(staticPath)).ServeHTTP)
	router.Get("/fonts/*", http.FileServer(http.Dir(staticPath)).ServeHTTP)

	if useSSO {
		router.Handle("/saml/*", ssoMiddleware)
	}

	router.Group(func(r chi.Router) {
		// Enforce SSO policy for these routes
		r.Use(ssoHandler)
		r.Get("/api/events", handleListEvents)
		//router.Post("/api/changeEvent", handleChangeEvents)

		r.Get("/events", handleEvents)
	})

	router.Get("/dataprotection", handleDataProtection)
	router.Get("/", handleRoot)

	// Manually specify timeout values
	// https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	srv := &http.Server{
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           router,
		Addr:              listeningAddr,
	}

	// Initialize additional routines

	// signal handling for graceful shutdown

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	signal.Ignore(syscall.SIGHUP) // do not exit if controlling terminal closes

	go func() {
		<-signalChan
		log.Println("Graceful shutdown requested")

		// TODO shutdown http server after requests are finished

		os.Exit(0)
	}()

	// Start listening for client requests

	if useTLS {
		log.Fatal(srv.Serve(autocert.NewListener("dh-application.nerdwiese.de")))
	} else {
		log.Fatal(srv.ListenAndServe())
	}
}
