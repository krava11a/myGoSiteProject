package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func StartPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	text := "Hello user!"
	fmt.Fprint(rw, text)
}
