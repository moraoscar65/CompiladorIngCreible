package sintactico

import (
	"errors"

	d "../definiciones"
	"../tokenTypes"
)

func bExpresion() (*d.Nodo, error) {
	var root, err = bTerm()
	if err != nil {
		root.TokenType.TypeLexema = tokenTypes.Error
		return root, err
	}
	if token.TypeLexema == tokenTypes.Or {
		var current, next *d.Nodo
		current = Inicializar()
		current.TokenType = token
		if match(tokenTypes.Or) {
			root.TokenType.TypeLexema = tokenTypes.Error
			err = errors.New("Se esperaba or")
			return root, err
		}
		current.Izq = root
		root = current
		root.Der, err = bTerm()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
			return root, err
		}
		for token.TypeLexema == tokenTypes.Or {
			next = Inicializar()
			next.TokenType = token
			if match(tokenTypes.Or) {
				root.TokenType.TypeLexema = tokenTypes.Error
				err = errors.New("Se esperaba or")
				return root, err
			}
			next.Izq = current.Der
			current.Der = next
			current = next
			current.Der, err = bTerm()
			if err != nil {
				root.TokenType.TypeLexema = tokenTypes.Error
				return root, err
			}
		}
	}
	return root, err
}

func bTerm() (*d.Nodo, error) {
	var root, err = notFactor()
	if err != nil {
		root.TokenType.TypeLexema = tokenTypes.Error
		return root, err
	}
	if token.TypeLexema == tokenTypes.And {
		var current, next *d.Nodo
		current = Inicializar()
		current.TokenType = token
		if match(tokenTypes.And) {
			root.TokenType.TypeLexema = tokenTypes.Error
			err = errors.New("Se esperaba and")
			return root, err
		}
		current.Izq = root
		root = current
		root.Der, err = notFactor()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
			return root, err
		}
		for token.TypeLexema == tokenTypes.And {
			next = Inicializar()
			next.TokenType = token
			if match(tokenTypes.And) {
				root.TokenType.TypeLexema = tokenTypes.Error
				err = errors.New("Se esperaba and")
				return root, err
			}
			next.Izq = current.Der
			current.Der = next
			current = next
			current.Der, err = notFactor()
			if err != nil {
				root.TokenType.TypeLexema = tokenTypes.Error
				return root, err
			}
		}
	}
	return root, err
}

func notFactor() (*d.Nodo, error) {
	var root, current, next *d.Nodo
	var err error = nil
	if token.TypeLexema == tokenTypes.Not {
		root = Inicializar()
		root.TokenType = token
		if match(tokenTypes.Not) {
			root.TokenType.TypeLexema = tokenTypes.Error
			err = errors.New("Se espera not")
			return root, err
		}
		current = root
		for token.TypeLexema == tokenTypes.Not {
			next = Inicializar()
			next.TokenType = token
			current.Der = next
			current = next
			if match(tokenTypes.Not) {
				root.TokenType.TypeLexema = tokenTypes.Error
				err = errors.New("Se espera not")
				return root, err
			}
		}
		current.Der, err = bFactor()
		if err != nil {
			root.TokenType.TypeLexema = tokenTypes.Error
			return root, err
		}
	} else {
		root, err = bFactor()
	}
	return root, err
}

func bFactor() (*d.Nodo, error) {
	var root *d.Nodo
	var err error = nil
	if token.TypeLexema == tokenTypes.True || token.TypeLexema == tokenTypes.False {
		root = Inicializar()
		root.TokenType = token
		root.Dtype = tokenTypes.Bool
		if match(token.TypeLexema) {
			root.TokenType.TypeLexema = tokenTypes.Error
			err = errors.New("Se esperaba un true | false")
		}
	} else {
		root, err = relacion()
	}
	return root, err
}
