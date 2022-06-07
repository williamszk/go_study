// https://www.youtube.com/watch?v=8S30eoBSojU

package main

import (
	"net/http"
	"regexp"
)

var (
	listUsersRe  = regexp.MustCompile(`^\/users[\/]*$`)
	getUserRe    = regexp.MustCompile(`^\/users\/(\d+)*$`) // /users/123
	createUserRe = regexp.MustCompile(`^\/users[\/]*$`)
)

// this is a struct with no key and values
type userHandler struct{}

// x = userHandler()
// h = &x
// typeof(h) -> *userHandler

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && listUsersRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getUserRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return
	}
	case r.Me
}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {

}

// Endpoints:
// GET /users
// GET /users/{id}
// POST /users

// JSON based

func main() {
	mux := http.ServeMux()
	mux.Handle("/users/", &userHandler{})
	// typeof(&userHandler()) -> *userHandler
	// mux.Handle("/posts/", &postsHandler{})
	http.ListenAndServe("localhost:8080", mux)

}
