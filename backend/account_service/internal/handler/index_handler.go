package handler

import (
	"log"
	"net/http"
	"path"
	"strings"
)

func IndexHandler(publicDir string) http.Handler {
	handler := http.FileServer(http.Dir(publicDir))

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_path := req.URL.Path

		log.Println("req.URL.Path: ", _path)

		// static files
		if strings.Contains(_path, ".") || _path == "/" {
			log.Println("Server static")
			handler.ServeHTTP(w, req)
			return
		}

		// the all 404 gonna be served as root
		log.Println("Server index")
		http.ServeFile(w, req, path.Join(publicDir, "/index.html"))
	})
}