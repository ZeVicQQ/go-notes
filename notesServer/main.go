package main

import (
	"Nodeprj/controllers/stdhttp"
	"Nodeprj/gate/psg"
	"fmt"
	"net/http"
)

func main() {
	s, err := psg.NewPsg("postgres://Alex:1234@localhost:5432/Node", "Alex", "1234")
	if err != nil {
		fmt.Println("Error")
	}
	stdhttp.NewController("localhost:8080", s)
	http.ListenAndServe(":8080", nil)
}
