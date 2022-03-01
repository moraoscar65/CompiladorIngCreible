package sintactico

import (
	d "../definiciones"
	"../tokenTypes"
)

func listaSentencias() *d.Nodo {
	var root *d.Nodo
	var current, next *d.Nodo
	if token.TypeLexema != tokenTypes.Eof {
		switch token.TypeLexema {
		case tokenTypes.If:
			root = seleccion()
		case tokenTypes.While:
			root = iteracion()
		case tokenTypes.Do:
			root = repeticion()
		case tokenTypes.Read:
			root = sentRead()
		case tokenTypes.Write:
			root = sentWrite()
		case tokenTypes.LlaveIzq:
			root = bloque()
		case tokenTypes.Ident:
			root = asignacion()
		case tokenTypes.LlaveDer:
			return root
		default:
			root = Inicializar()
			root.TokenType.TypeLexema = tokenTypes.Error
			root.TokenType.Lexema = token.Lexema
			PanicMode(tokenTypes.PuntoComa)
		}
		current = root
		for token.TypeLexema != tokenTypes.Eof {
			switch token.TypeLexema {
			case tokenTypes.If:
				next = seleccion()
			case tokenTypes.While:
				next = iteracion()
			case tokenTypes.Do:
				next = repeticion()
			case tokenTypes.Read:
				next = sentRead()
			case tokenTypes.Write:
				next = sentWrite()
			case tokenTypes.LlaveIzq:
				next = bloque()
			case tokenTypes.Ident:
				next = asignacion()
			case tokenTypes.LlaveDer:
				return root
			default:
				next = Inicializar()
				next.TokenType.TypeLexema = tokenTypes.Error
				next.TokenType.Lexema = token.Lexema
				PanicMode(tokenTypes.PuntoComa)
			}
			current.Bro = next
			current = next
		}
	}
	return root
}

func checkListaSenteciasType() bool {
	return token.TypeLexema == tokenTypes.Ident || token.TypeLexema == tokenTypes.Read || token.TypeLexema == tokenTypes.LlaveIzq || token.TypeLexema == tokenTypes.Write || token.TypeLexema == tokenTypes.While || token.TypeLexema == tokenTypes.If || token.TypeLexema == tokenTypes.Do
}

func seleccion() *d.Nodo {
	var err error
	var root = Inicializar()
	root.TokenType = token
	if match(tokenTypes.If) {
		PanicMode(tokenTypes.Fi)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	if match(tokenTypes.ParenIzq) {
		PanicMode(tokenTypes.Fi)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	root.Izq, err = bExpresion()
	if err != nil {
		PanicMode(tokenTypes.Fi)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	if match(tokenTypes.ParenDer) {
		PanicMode(tokenTypes.Fi)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	if match(tokenTypes.Then) {
		PanicMode(tokenTypes.Fi)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	root.Med = bloque()
	if token.TypeLexema == tokenTypes.Else {
		if match(tokenTypes.Else) {
			PanicMode(tokenTypes.Fi)
			root.TokenType.TypeLexema = tokenTypes.Error
			return root
		}
		root.Der = bloque()
	}
	if match(tokenTypes.Fi) {
		PanicMode(tokenTypes.Fi)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	return root
}

func iteracion() *d.Nodo {
	var err error
	var root = Inicializar()
	root.TokenType = token
	if match(tokenTypes.While) {
		PanicMode(tokenTypes.LlaveDer)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	if match(tokenTypes.ParenIzq) {
		PanicMode(tokenTypes.LlaveDer)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	root.Izq, err = bExpresion()
	if err != nil {
		PanicMode(tokenTypes.LlaveDer)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	if match(tokenTypes.ParenDer) {
		PanicMode(tokenTypes.LlaveDer)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	root.Med = bloque()
	return root
}

func repeticion() *d.Nodo {
	var err error
	var root = Inicializar()
	root.TokenType = token
	if match(tokenTypes.Do) {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	root.Izq = bloque()
	if match(tokenTypes.Until) {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	if match(tokenTypes.ParenIzq) {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	root.Med, err = bExpresion()
	if err != nil {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	if match(tokenTypes.ParenDer) {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	if match(tokenTypes.PuntoComa) {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	return root
}

func sentRead() *d.Nodo {
	var root *d.Nodo
	var err error
	root = Inicializar()
	root.TokenType = token
	if match(tokenTypes.Read) {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	root.Med, err = identificador()
	if err != nil {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	if match(tokenTypes.PuntoComa) {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	return root
}

func sentWrite() *d.Nodo {
	var root *d.Nodo
	var err error = nil
	root = Inicializar()
	root.TokenType = token
	if match(tokenTypes.Write) {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	root.Med, err = bExpresion() //Pendiente
	if err != nil {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	if match(tokenTypes.PuntoComa) {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	return root
}

func bloque() *d.Nodo {
	var root *d.Nodo
	if match(tokenTypes.LlaveIzq) {
		PanicMode(tokenTypes.LlaveDer)
		root = Inicializar()
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	root = listaSentencias()
	if match(tokenTypes.LlaveDer) {
		PanicMode(tokenTypes.LlaveDer)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	return root
}

func asignacion() *d.Nodo {
	var root = &d.Nodo{}
	var err error = nil
	root.Izq, err = identificador()
	if err != nil {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	root.TokenType = token
	if match(tokenTypes.Asig) {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	root.Der, err = bExpresion()
	if err != nil {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	if match(tokenTypes.PuntoComa) {
		PanicMode(tokenTypes.PuntoComa)
		root.TokenType.TypeLexema = tokenTypes.Error
		return root
	}
	return root
}
