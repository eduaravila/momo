package main

import (
	"log"

	"github.com/eduaravila/momo/packages/postgres/queries"
)

func main() {
	var user queries.User
	log.Println(user)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// router := v1.Handler(queries.New(db), adapter.NewTwitchAPI())
	// log.Fatal(http.ListenAndServe(":8080", router))
}
