package semantico

import (
	d "../definiciones"
	lm "../logger"
	tt "../tokenTypes"
	"errors"
	"fmt"
	"strconv"
)

var TSimbolos = make(map[string]d.Simbolo)
var currentType int
var currentVar string

func Semantico(arbol *d.Nodo) {
	partida(arbol)
}

func partida(nodo *d.Nodo) {
	if nodo.Izq != nil {
		declaraciones(nodo.Izq)
	}
	if nodo.Der != nil {
		sentencias(nodo.Der)
	}
}

func declaraciones(nodo *d.Nodo) {
	if nodo != nil {
		if checarTipo(nodo.TokenType.TypeLexema) {
			currentType = nodo.TokenType.TypeLexema
			nodo.Dtype = currentType
		}
		declaraciones(nodo.Med)
		if !checarTipo(nodo.TokenType.TypeLexema) {
			TSimbolos[nodo.TokenType.Lexema] = d.Simbolo{Dtype: currentType, Valor: "0"}
			nodo.Dtype = currentType
			//fmt.Println("Agregado a tabla de simbolos: " + nodo.TokenType.Lexema)
			lm.AgregarSemantico("Variable añadida a la tabla de simbolos { " + nodo.TokenType.Lexema + " - No inicializado }")
		}
		if nodo.Bro != nil {
			declaraciones(nodo.Bro)
		}
	}
}

func checarTipo(tipo int) bool {
	if tipo == tt.Int || tipo == tt.Float || tipo == tt.Bool {

		return true
	}
	return false
}

func checarTabla(token d.Token) bool {
	if token.TypeLexema == TSimbolos[currentVar].Dtype {
		aux := TSimbolos[currentVar].Dtype
		TSimbolos[currentVar] = d.Simbolo{Dtype: aux, Valor: token.Lexema}
		return true
	}
	return false
}

func sentencias(nodo *d.Nodo) {
	if nodo != nil {
		switch nodo.TokenType.TypeLexema {
		case tt.Asig:
			switch TSimbolos[nodo.Izq.TokenType.Lexema].Dtype {
			case tt.Int:
				err := checarIntExp(nodo.Der)
				if err != nil {
					fmt.Println(err)
					lm.AgregarError(err.Error())
				} else {
					nodo.Dtype = tt.Int
					nodo.Izq.Dtype = tt.Int
					TSimbolos[nodo.Izq.TokenType.Lexema] = d.Simbolo{Dtype: tt.Int, Valor: nodo.Der.AtrValor}
					lm.AgregarSemantico("Tabla de simbolos { 'variable' -> " + nodo.Izq.TokenType.Lexema + " ;; 'Valor' -> " + nodo.Der.AtrValor + ";; 'dType' -> Int }")
					//fmt.Println("Tabla modificada: " + nodo.Izq.TokenType.Lexema + " -> " + nodo.Der.AtrValor)
				}
				break
			case tt.Float:
				err := checarFloatExp(nodo.Der)
				if err != nil {
					fmt.Println(err.Error())
					lm.AgregarError(err.Error())
				} else {
					nodo.Dtype = tt.Float
					nodo.Izq.Dtype = tt.Float
					TSimbolos[nodo.Izq.TokenType.Lexema] = d.Simbolo{Dtype: tt.Float, Valor: nodo.Der.AtrValor}
					lm.AgregarSemantico("Tabla de simbolos { 'variable' -> " + nodo.Izq.TokenType.Lexema + " ;; 'Valor' -> " + nodo.Der.AtrValor + ";; 'dType' -> Float }")
					//fmt.Println("Tabla modificada: " + nodo.Izq.TokenType.Lexema + " -> " + nodo.Der.AtrValor)
				}
				break

			case tt.Bool:
				err := checarBoolExp(nodo.Der)
				if err != nil {
					fmt.Println(err.Error())
					lm.AgregarError(err.Error())
				} else {
					nodo.Dtype = tt.Bool
					nodo.Izq.Dtype = tt.Bool
					TSimbolos[nodo.Izq.TokenType.Lexema] = d.Simbolo{Dtype: tt.Bool, Valor: nodo.Der.TokenType.Lexema}
					lm.AgregarSemantico("Tabla de simbolos { 'variable' -> " + nodo.Izq.TokenType.Lexema + " ;; 'Valor' -> " + nodo.Der.TokenType.Lexema + ";; 'dType' -> Bool }")
					//fmt.Println("Tabla modificada: " + nodo.Izq.TokenType.Lexema + " -> " + nodo.Der.TokenType.Lexema)
				}
				break
			default:
				fmt.Println("Valor Dtype simbolo indefinido")
				fmt.Println(nodo.TokenType)
			}
			break
		case tt.If:
			var err error
			err = checarBoolExp(nodo.Izq)
			if err != nil {
				fmt.Println(err.Error())
				lm.AgregarError(err.Error())
			}
			sentencias(nodo.Med)
			sentencias(nodo.Der)
			break
		case tt.While:
			var err error
			err = checarBoolExp(nodo.Izq)
			if err != nil {
				lm.AgregarError(err.Error())
				fmt.Println(err.Error())
			}
			sentencias(nodo.Med)
			break
		case tt.Do:
			var err error
			err = checarBoolExp(nodo.Med)
			if err != nil {
				fmt.Println(err.Error())
				lm.AgregarError(err.Error())
			}
			sentencias(nodo.Izq)
			break
		default:
			fmt.Println("Operación indefinida")
			fmt.Println(nodo.TokenType)
		}
		if nodo.Bro != nil {
			sentencias(nodo.Bro)
		}
	}
}

func obtenerValorNumerico(nodo *d.Nodo) string {
	var val1, val2 int
	if nodo.Izq.TokenType.TypeLexema == tt.Ident {
		val1, _ = strconv.Atoi(TSimbolos[nodo.Izq.TokenType.Lexema].Valor)
	} else {
		val1, _ = strconv.Atoi(nodo.Izq.TokenType.Lexema)
	}
	if nodo.Der.TokenType.TypeLexema == tt.Ident {
		val2, _ = strconv.Atoi(TSimbolos[nodo.Der.TokenType.Lexema].Valor)
	} else {
		val2, _ = strconv.Atoi(nodo.Der.TokenType.Lexema)
	}
	var val int
	switch nodo.TokenType.TypeLexema {
	case tt.Plus:
		val = val1 + val2
	case tt.Minus:
		val = val1 - val2
	case tt.Div:
		val = val1 / val2
	case tt.Multi:
		val = val1 * val2
	}

	return strconv.Itoa(val)
}

func obtenerValorDecimal(nodo *d.Nodo) string {
	var val1, val2 float64
	tamFloat := 32
	if nodo.Izq.TokenType.TypeLexema == tt.Ident {
		val1, _ = strconv.ParseFloat(TSimbolos[nodo.Izq.TokenType.Lexema].Valor, tamFloat)
	} else {
		val1, _ = strconv.ParseFloat(nodo.Izq.TokenType.Lexema, tamFloat)
	}
	if nodo.Der.TokenType.TypeLexema == tt.Ident {
		val2, _ = strconv.ParseFloat(TSimbolos[nodo.Der.TokenType.Lexema].Valor, tamFloat)
	} else {
		val2, _ = strconv.ParseFloat(nodo.Der.TokenType.Lexema, tamFloat)
	}
	var val float64
	switch nodo.TokenType.TypeLexema {
	case tt.Plus:
		val = val1 + val2
	case tt.Minus:
		val = val1 - val2
	case tt.Div:
		val = val1 / val2
	case tt.Multi:
		val = val1 * val2
	}

	return fmt.Sprintf("%f", val)
}

func checarIntExp(nodo *d.Nodo) error {
	var err error
	switch nodo.TokenType.TypeLexema {
	case tt.Plus, tt.Minus, tt.Div, tt.Multi:
		err1 := checarIntExp(nodo.Izq)
		err2 := checarIntExp(nodo.Der)
		if err1 != nil {
			err = err1
		} else if err2 != nil {
			err = err2
		} else {
			err = nil
			nodo.AtrValor = obtenerValorNumerico(nodo)
		}
	case tt.Ident:
		if TSimbolos[nodo.TokenType.Lexema].Dtype == tt.Int {
			err = nil
			nodo.AtrValor = TSimbolos[nodo.TokenType.Lexema].Valor
		} else {
			err = errors.New("No es Int" + "(" + strconv.Itoa(nodo.TokenType.NumFila) + "," + strconv.Itoa(nodo.TokenType.NumCol) + ")")

		}
	case tt.Numero:
		if nodo.Dtype == tt.Int {
			err = nil
			nodo.AtrValor = nodo.TokenType.Lexema
		} else {
			err = errors.New("No es Int" + "(" + strconv.Itoa(nodo.TokenType.NumFila) + "," + strconv.Itoa(nodo.TokenType.NumCol) + ")")

		}
	case tt.Decimal:
		err = errors.New("No es Int" + "(" + strconv.Itoa(nodo.TokenType.NumFila) + "," + strconv.Itoa(nodo.TokenType.NumCol) + ")")

	default:
		err = errors.New("expresion int indefinido")
	}

	return err
}

func checarFloatExp(nodo *d.Nodo) error {
	var err error
	switch nodo.TokenType.TypeLexema {
	case tt.Plus, tt.Minus, tt.Div, tt.Multi:
		err1 := checarFloatExp(nodo.Izq)
		err2 := checarFloatExp(nodo.Der)
		if err1 != nil {
			err = err1
		} else if err2 != nil {
			err = err2
		} else {
			err = nil
			//nodo.AtrValor = obtenerValorDecimal(nodo)
		}
	case tt.Ident:
		if TSimbolos[nodo.TokenType.Lexema].Dtype == tt.Float {
			err = nil
			nodo.AtrValor = TSimbolos[nodo.TokenType.Lexema].Valor
		} else {
			err = errors.New("No es un Float" + "(" + strconv.Itoa(nodo.TokenType.NumFila) + "," + strconv.Itoa(nodo.TokenType.NumCol) + ")")

		}
	case tt.Decimal:
		if nodo.Dtype == tt.Float {
			err = nil
			nodo.AtrValor = nodo.TokenType.Lexema
		} else {
			err = errors.New("No es un Float" + "(" + strconv.Itoa(nodo.TokenType.NumFila) + "," + strconv.Itoa(nodo.TokenType.NumCol) + ")")

		}
	case tt.Numero:
		if nodo.Dtype == tt.Int {
			err = nil
			nodo.AtrValor = nodo.TokenType.Lexema + ".0"
		} else {
			err = errors.New("No es un Float" + "(" + strconv.Itoa(nodo.TokenType.NumFila) + "," + strconv.Itoa(nodo.TokenType.NumCol) + ")")

		}
	default:
		err = errors.New("No es un flotante")
	}
	return err
}

func checarBoolExp(nodo *d.Nodo) error {
	var err error
	switch nodo.TokenType.TypeLexema {
	case tt.And, tt.Or, tt.Not:
		err1 := checarBoolExp(nodo.Izq)
		err2 := checarBoolExp(nodo.Der)
		if err1 != nil {
			err = err1
		} else if err2 != nil {
			err = err2
		} else {
			err = nil
		}
	case tt.Menor, tt.MenorIgual, tt.Mayor, tt.MayorIgual, tt.Comparacion, tt.DiferenteDe:
		intErr1 := checarIntExp(nodo.Izq)
		intErr2 := checarIntExp(nodo.Der)
		floatErr1 := checarFloatExp(nodo.Izq)
		floatErr2 := checarFloatExp(nodo.Der)
		if (intErr1 == nil && intErr2 == nil) || (floatErr1 == nil && floatErr2 == nil) || (floatErr1 == nil && intErr2 == nil) || (intErr1 == nil && floatErr2 == nil) {
			err = nil
		} else {
			err = errors.New("Los tipos de datos en las expresiones no coninciden1")
		}
	case tt.Ident:
		if TSimbolos[nodo.TokenType.Lexema].Dtype == tt.Bool {
			err = nil
		} else {
			err = errors.New("Identificador No es un booleano1" + "(" + strconv.Itoa(nodo.TokenType.NumFila) + "," + strconv.Itoa(nodo.TokenType.NumCol) + ")")
		}

	case tt.True, tt.False:
		if nodo.Dtype == tt.Bool {
			err = nil
		} else {
			err = errors.New("No es un booleano" + "(" + strconv.Itoa(nodo.TokenType.NumFila) + "," + strconv.Itoa(nodo.TokenType.NumCol) + ")")
		}
	default:
		err = errors.New("No es un booleano")
	}
	return err
}
