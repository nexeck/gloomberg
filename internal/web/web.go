package web

//
//  gloomberg websockets server & web ui
//
//  created with help of this awesome articly by @percybolmer 👌 thanks!
//  https://programmingpercy.tech/blog/mastering-websockets-with-go/
//

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/charmbracelet/log"
	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/webtransport-go"
)

func StartWebUI(queueWsOutTokenTransactions chan *totra.TokenTransaction) {
	// Create a Manager instance used to handle WebSocket Connections
	hub := NewHub(queueWsOutTokenTransactions)

	listenOn := ":8080"
	certPath := "./home.benleb.de.crt"
	keyPath := "./home.benleb.de.key"

	// load index template
	// tmpl := template.Must(template.ParseFiles("www/index.html"))

	tmplFiles := []string{"www/index.html", "www/style.html", "www/javascript.html"}

	tmpl, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		gbl.Log.Error(err)
	}

	go func() {
		// create a new webtransport.Server, listening on (UDP) port 443 (8080)
		s := webtransport.Server{H3: http3.Server{Addr: listenOn}}

		// Create a new HTTP endpoint /webtransport.
		http.HandleFunc("/webtransport", func(w http.ResponseWriter, r *http.Request) {
			conn, err := s.Upgrade(w, r)
			if err != nil {
				log.Printf("upgrading failed: %s", err)
				w.WriteHeader(http.StatusInternalServerError)

				return
			}

			// Handle the connection. Here goes the application logic.
			log.Print(fmt.Sprintf("new connection from %s | %+v", conn.RemoteAddr(), conn))
		})

		log.Fatal(s.ListenAndServeTLS(certPath, keyPath))
	}()

	// index page
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		data := map[string]string{
			"Title": "gloomberg | " + internal.GloombergVersion,
		}

		if err := tmpl.Execute(w, data); err != nil {
			gbl.Log.Error("Error executing template: ", err)
		}
	})

	// static js files (the stripping feels a bit weird...)
	http.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./www/js"))))

	// websocket endpoint
	http.HandleFunc("/ws", hub.serveWS)

	// load tls certificate
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatal(err)
	}

	// create http server
	server := &http.Server{
		Addr:              listenOn,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           nil,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
			MinVersion:   tls.VersionTLS12,
			MaxVersion:   0,
		},
	}

	// start http server
	log.Fatal(server.ListenAndServe())
}
