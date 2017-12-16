package service

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/v1/user/login", userLoginHandler(formatter)).Methods("POST")

	mx.HandleFunc("/v1/users", postUserHandler(formatter)).Methods("POST")
	//mx.PathPrefix("/v1").Handler(http.HandlerFunc(testHandler))
	mx.HandleFunc("/v1/users", getAllUserHandler(formatter)).Methods("GET")

	mx.HandleFunc("/v1/users", deleteUserHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/v1/meetings", postMeetingHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/meetings", getAllMeetingsHandler(formatter)).Methods("GET")

}

func testHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello world")
}
