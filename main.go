package main

import (
	"log"
	"github.com/lutrueba/microblog/handlers"
	"github.com/lutrueba/microblog/bd"
)
func main(){
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}