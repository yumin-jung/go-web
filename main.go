package main

import (
	"net/http"

	"github.com/yumin-jung/go-web/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
