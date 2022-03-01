package sintactico

import (
	"errors"

	d "../definiciones"
	"../tokenTypes"
)

func relacion() (*d.Nodo, error) {
	root, err := expresion()
	if err != nil {
		root.TokenType.TypeLexema = tokenTypes.Error
		return root, err
	}
	if token.TypeLexema == tokenTypes.MenorIgual || token.TypeLexema == tokenTypes.Menor || token.TypeLexema == tokenTypes.MayorIgual || token.TypeLexema == tokenTypes.Mayor || token.TypeLexema == tokenTypes.Comparacion || token.TypeLexema == tokenTypes.DiferenteDe {
		var current *d.Nodo
		current, err = relOp()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
			return root, err
		}
		current.Izq = root
		root = current
		root.Der, err = expresion()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
			return root, err
		}
	}
	return root, err
}

func relOp() (*d.Nodo, error) {
	var root = Inicializar()
	root.TokenType = token
	

	var err error = nil
	if match(token.TypeLexema) {
		root.TokenType.TypeLexema = tokenTypes.Error
		err = errors.New("Se esperaba <= | < | > | >= | == | !=")
	}
	return root, err
}

func expresion() (*d.Nodo, error) {
	root, err := termino()
	if err != nil {
		root.TokenType.TypeLexema = tokenTypes.Error
		return root, err
	}
	if token.TypeLexema == tokenTypes.Plus || token.TypeLexema == tokenTypes.Minus {
		var current, next *d.Nodo
		current, err = sumaOp()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
			return root, err
		}
		current.Izq = root
		root = current
		root.Der, err = termino()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
			return root, err
		}
		for token.TypeLexema == tokenTypes.Plus || token.TypeLexema == tokenTypes.Minus {
			next, err = sumaOp()
			if err != nil {
				root.TokenType.TypeLexema = tokenTypes.Error
				return root, err
			}
			next.Izq = current.Der
			current.Der = next
			current = next
			current.Der, err = termino()
			if err != nil {
				root.TokenType.TypeLexema = tokenTypes.Error
				return root, err
			}
		}
	}
	return root, err
}

func sumaOp() (*d.Nodo, error) {
	var root = Inicializar()
	root.TokenType = token
	var err error = nil
	if match(token.TypeLexema) {
		root.TokenType.TypeLexema = tokenTypes.Error
		err = errors.New("Se esperaba + | -")
	}
	return root, err
}

func termino() (*d.Nodo, error) {
	root, err := signoFactor()
	if err != nil {
		root.TokenType.TypeLexema = tokenTypes.Error
		return root, err
	}
	if token.TypeLexema == tokenTypes.Multi || token.TypeLexema == tokenTypes.Div {
		var current, next *d.Nodo
		current, err = multOp()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
			return root, err
		}
		current.Izq = root
		root = current
		root.Der, err = signoFactor()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
			return root, err
		}
		for token.TypeLexema == tokenTypes.Multi || token.TypeLexema == tokenTypes.Div {
			next, err = multOp()
			root.TokenType.TypeLexema = tokenTypes.Error
			return root, err
			next.Izq = current.Der
			current.Der = next
			current = next
			current.Der, err = signoFactor()
			if err != nil {
				root.TokenType.TypeLexema = tokenTypes.Error
				return root, err
			}
		}
	}
	return root, err
}

func multOp() (*d.Nodo, error) {
	var root = Inicializar()
	root.TokenType = token

	var err error = nil
	if match(token.TypeLexema) {
		root.TokenType.TypeLexema = tokenTypes.Error
		err = errors.New("Se esperaba un  * | /")
	}
	return root, err
}

func signoFactor() (*d.Nodo, error) {
	var root = Inicializar()
	var err error = nil
	if token.TypeLexema == tokenTypes.Plus || token.TypeLexema == tokenTypes.Minus {
		root, err = sumaOp()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
			return root, err
		}
		root.Der, err = factor()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
		}
	} else {
		root, err = factor()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
		}
	}
	return root, err
}

func factor() (*d.Nodo, error) {
	var root = Inicializar()
	var err error = nil
	if token.TypeLexema == tokenTypes.ParenIzq {
		if match(tokenTypes.ParenIzq) {
			root.TokenType.TypeLexema = tokenTypes.Error
			err = errors.New("Se esperaba un  (")
			return root, err
		}
		root, err = bExpresion()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
			return root, err
		}
		if match(tokenTypes.ParenDer) {
			root.TokenType.TypeLexema = tokenTypes.Error
			err = errors.New("Se esperaba un  )")
			return root, err
		}
	} else if token.TypeLexema == tokenTypes.Ident {
		root, err = identificador()
	} else {
		if(token.TypeLexema == tokenTypes.Numero){
			root, err = numero()
		}else{
			root, err = decimal()
		}
		
	}
	return root, err
}
