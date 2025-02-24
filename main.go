package main

import (
	"fmt"
	"log"

	"github.com/MudassirDev/all-in-one-business-tool/api/server"
)

func main() {
	srv := server.CreateServer()

	fmt.Printf("listening at port %v\n", srv.Addr)
	err := srv.ListenAndServe()
	log.Printf("failed to start server! err: %v\n", err.Error())
}
