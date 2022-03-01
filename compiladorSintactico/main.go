package main

import (
	"./intermedio"
	"./logger"
	as "./semantico"
	s "./sintactico"
	"fmt"
	"os"
)

func main() {

	if len(os.Args[1:]) != 0 {
		arguments := os.Args[1:]
		var file string = arguments[0]
		arbol := s.Parse(file)
		as.Semantico(arbol)
		codigoIntermedio.GenCode(arbol.Der)
		logger.AgregarCode("(halt,_,_,_)")
		logger.CrearArchivo()
		logger.CrearArchivoTree()
	} else {
		fmt.Println("Vacío")
	}
	return
}
