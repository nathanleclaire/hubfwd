package main

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func OfficialImageHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	libraryimg := vars["libraryimg"]
	log.WithField("libraryimg", libraryimg).Info("Received request to redirect to official image")
	http.Redirect(w, req, fmt.Sprintf("https://registry.hub.docker.com/_/%s", libraryimg), 301)
}

func UserImageHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	user := vars["user"]
	img := vars["img"]
	log.WithFields(log.Fields{
		"user": user,
		"img":  img,
	}).Info("Received request to redirect to user image")
	http.Redirect(w, req, fmt.Sprintf("https://hub.docker.com/u/%s/%s", user, img), 301)
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	resp := `
<!doctype html>
<html><head><title>Docker Image Redirect</title>
<link href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css" rel="stylesheet">
<style>
* { font-family: monospace; }
</style>
</head>
<body>
<center>
<h1>
Docker Image Redirector
</h1>
<p>
The Docker Hub URL schemes are hard to remember and type.
</p>
<p>
This website provides a shorthand for accessing Docker Hub images.
</p>
<p>To go to an official (library) image: <a href="http://dkrimg.com/golang">dkrimg.com/golang</a></p>
<p>To go to an unofficial (user) image: <a href="http://dkrimg.com/nathanleclaire/zshdev">dkrimg.com/nathanleclaire/zshdev</p>
</center>
</body>
</html>
`
	fmt.Fprintf(w, resp)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/{libraryimg}", OfficialImageHandler)
	router.HandleFunc("/{user}/{img}", UserImageHandler)
	n := negroni.New()
	n.UseHandler(router)
	n.Run(":3000")
}
