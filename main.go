package main

import (
	"github.com/pluralsight/webservice/controllers"
	"net/http"
	// "github.com/pluralsight/webservice/models"
	// "fmt"
)

func main() {
	controllers.RegisterControllers();
	http.ListenAndServe(":3000", nil)
}
