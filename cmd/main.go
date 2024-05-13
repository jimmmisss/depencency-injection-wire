package cmd

import (
	"net/http"
	"wire/user"
)

func main() {
	userHandler := user.Wire(nil)
	http.Handle("/user", userHandler.FetchByUserName())
	http.ListenAndServe(":8080", nil)
}
