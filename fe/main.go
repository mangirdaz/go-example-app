/*
Copyright 2016 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Command template is a trivial web server that uses the text/template (and
// html/template) package's "block" feature to implement a kind of template
// inheritance.
//
// It should be executed from the directory in which the source resides,
// as it will look for its template files in the current directory.
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/bobbydeveaux/go-example-app/config"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

type Handlers struct {
	images map[string]*Image
}

func main() {
	log.Info("Start Frontend")

	var images = map[string]*Image{
		"go":         {"The Go Gopher", "https://golang.org/doc/gopher/frontpage.png"},
		"redhat":     {"Red Hat Logo", "https://upload.wikimedia.org/wikipedia/en/thumb/6/6c/RedHat.svg/1280px-RedHat.svg.png"},
		"openshift":  {"OCP Logo", "https://blog.openshift.com/wp-content/uploads/openshift_container_platform.png"},
		"kubernetes": {"K8S Logo", "https://blog.osones.com/en/images/docker/kubernetes.png"},
		"random":     {"Random", config.Get("EnvAPIServiceURL") + "/api/v1/img"},
	}

	hn := Handlers{
		images: images,
	}

	http.HandleFunc("/", hn.indexHandler)
	http.HandleFunc("/image/", hn.imageHandler)
	url := fmt.Sprintf("%s:%s", config.Get("EnvFEIP"), config.Get("EnvFEPort"))
	log.Fatal(http.ListenAndServe(url, nil))
}

// indexTemplate is the main site template.
// The default template includes two template blocks ("sidebar" and "content")
// that may be replaced in templates derived from this one.
var indexTemplate = template.Must(template.ParseFiles("/deployment/index.tmpl"))

// Index is a data structure used to populate an indexTemplate.
type Index struct {
	Title    string
	Body     string
	Links    []Link
	Revealed string
}

type Link struct {
	URL, Title string
}

// indexHandler is an HTTP handler that serves the index page.
func (hn Handlers) indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("/ hit")
	data := &Index{
		Title:    "UKCloud Image gallery",
		Body:     "Welcome to the UKCloud gallery.",
		Revealed: config.Get("EnvAPIPassword"),
	}

	// images specifies the site content: a collection of images.

	for name, img := range hn.images {
		data.Links = append(data.Links, Link{
			URL:   "/image/" + name,
			Title: img.Title,
		})
	}
	if err := indexTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

// imageTemplate is a clone of indexTemplate that provides
// alternate "sidebar" and "content" templates.
var imageTemplate = template.Must(template.Must(indexTemplate.Clone()).ParseFiles("/deployment/image.tmpl"))

// Image is a data structure used to populate an imageTemplate.
type Image struct {
	Title string
	URL   string
}

// imageHandler is an HTTP handler that serves the image pages.
func (hn Handlers) imageHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("/image/ hit")
	// images specifies the site content: a collection of images.

	data, ok := hn.images[strings.TrimPrefix(r.URL.Path, "/image/")]
	if !ok {
		http.NotFound(w, r)
		return
	}
	if err := imageTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}
