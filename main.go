package main

import (
	"net/http"

	"github.com/richeney/pluralsight-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
