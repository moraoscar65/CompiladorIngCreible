package sintactico

import (
	"errors"
	"fmt"

	d "../definiciones"
	"../tokenTypes"
)

func listaDeclaracion() *d.Nodo {
	var root *d.Nodo
	if token.TypeLexema == tokenTypes.Int || token.TypeLexema == tokenTypes.Float || token.TypeLexema == tokenTypes.Bool {
		root = declaracion()
		var current, next *d.Nodo
		current = root
		for token.TypeLexema == tokenTypes.Int || token.TypeLexema == tokenTypes.Float || token.TypeLexema == tokenTypes.Bool {
			next = declaracion()
			current.Bro = next
			current = next
		}
	}
	return root
}

func declaracion() *d.Nodo {
	var root *d.Nodo
	var err error
	root = tipo()
	root.Med, err = listaId()
	if err != nil {
		fmt.Println("Se llamara a PanicMode para buscar ;")
		PanicMode(tokenTypes.PuntoComa)
		return root
	}
	if match(tokenTypes.PuntoComa) {
		fmt.Println("Se llamara a PanicMode para buscar ;")
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	return root
}

func tipo() *d.Nodo {
	var root = Inicializar()
	root.TokenType.TypeLexema = tokenTypes.Error
	root.TokenType = token
	match(token.TypeLexema)
	return root
}

func listaId() (*d.Nodo, error) {
	root, err := identificador()
	if err != nil {
		return root, err
	}
	var current, next *d.Nodo
	current = root
	for token.TypeLexema == tokenTypes.Coma {
		if match(tokenTypes.Coma) {
			fmt.Println("Error al hacer match de la coma, regresando el error")
			root.TokenType.TypeLexema = tokenTypes.Error
			err = errors.New("Error: Se esperaba una coma")
			return root, err
		}
		next, err = identificador()
		if err != nil {
			return root, err
		}
		current.Med = next
		current = next
	}
	return root, nil
}
