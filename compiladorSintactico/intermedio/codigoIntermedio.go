package codigoIntermedio

import (
	d "../definiciones"
	lm "../logger"
	tt "../tokenTypes"
	"strconv"
)

var variablesTemporales = 0
var etiquetasTemporales = 0

func GenCode(Nodo *d.Nodo) {
	if Nodo != nil {
		switch Nodo.TokenType.TypeLexema {
		case tt.Asig:
			result := genCodeExpression(Nodo.Der)
			lm.AgregarCode("(asn," + result + "," + Nodo.Izq.TokenType.Lexema + ",_)")
		case tt.If:
			expression := genCodeExpression(Nodo.Izq)
			etiquetasTemporales++
			label1 := etiquetasTemporales
			lm.AgregarCode("(if_f," + expression + "," + "L" + strconv.Itoa(label1) + "," + "_)")
			if Nodo.Med != nil {
				GenCode(Nodo.Med)
			}
			etiquetasTemporales++
			label2 := etiquetasTemporales
			lm.AgregarCode("(goto," + "L" + strconv.Itoa(label2) + ",_,_)")
			lm.AgregarCode("(label," + "L" + strconv.Itoa(label1) + ",_,_)")
			if Nodo.Der != nil {
				GenCode(Nodo.Der)
			}
			lm.AgregarCode("(label," + "L" + strconv.Itoa(label2) + ",_,_)")
		case tt.While:
			etiquetasTemporales++
			label1 := etiquetasTemporales
			lm.AgregarCode("(label," + "L" + strconv.Itoa(label1) + ",_,_)")
			expression := genCodeExpression(Nodo.Izq)
			etiquetasTemporales++
			label2 := etiquetasTemporales
			lm.AgregarCode("(if_f," + expression + "," + "L" + strconv.Itoa(label2) + "," + "_)")
			if Nodo.Med != nil {
				GenCode(Nodo.Med)
			}
			lm.AgregarCode("(goto," + "L" + strconv.Itoa(label1) + ",_,_)")
			lm.AgregarCode("(label," + "L" + strconv.Itoa(label2) + ",_,_)")
		case tt.Do:
			etiquetasTemporales++
			label1 := etiquetasTemporales
			lm.AgregarCode("(label," + "L" + strconv.Itoa(label1) + ",_,_)")
			if Nodo.Izq != nil {
				GenCode(Nodo.Izq)
			}
			expression := genCodeExpression(Nodo.Med)
			etiquetasTemporales++
			label2 := etiquetasTemporales
			lm.AgregarCode("(if_f," + expression + "," + "L" + strconv.Itoa(label2) + "," + "_)")
			lm.AgregarCode("(goto," + "L" + strconv.Itoa(label1) + ",_,_)")
			lm.AgregarCode("(label," + "L" + strconv.Itoa(label2) + ",_,_)")
		case tt.Read:
			lm.AgregarCode("(rd," + Nodo.Med.TokenType.Lexema + ",_,_)")
		case tt.Write:
			center := genCodeExpression(Nodo.Med)
			lm.AgregarCode("(wri," + center + ",_,_)")
		}
		if Nodo.Bro != nil {
			GenCode(Nodo.Bro)
		}
	}
}

func genCodeExpression(Nodo *d.Nodo) string {
	switch Nodo.TokenType.TypeLexema {
	case tt.Plus:
		left := genCodeExpression(Nodo.Izq)
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(add," + left + "," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ")")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.Minus:
		left := genCodeExpression(Nodo.Izq)
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(sub," + left + "," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ")")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.Div:
		left := genCodeExpression(Nodo.Izq)
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(div," + left + "," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ")")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.Multi:
		left := genCodeExpression(Nodo.Izq)
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(mul," + left + "," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ")")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.And:
		left := genCodeExpression(Nodo.Izq)
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(and," + left + "," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ")")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.Or:
		left := genCodeExpression(Nodo.Izq)
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(or," + left + "," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ")")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.Not:
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(not," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ",_)")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.Menor:
		left := genCodeExpression(Nodo.Izq)
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(lt," + left + "," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ")")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.MenorIgual:
		left := genCodeExpression(Nodo.Izq)
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(leq," + left + "," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ")")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.Mayor:
		left := genCodeExpression(Nodo.Izq)
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(gt," + left + "," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ")")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.MayorIgual:
		left := genCodeExpression(Nodo.Izq)
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(geq," + left + "," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ")")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.Comparacion:
		left := genCodeExpression(Nodo.Izq)
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(eq," + left + "," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ")")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.DiferenteDe:
		left := genCodeExpression(Nodo.Izq)
		right := genCodeExpression(Nodo.Der)
		variablesTemporales++
		lm.AgregarCode("(ineq," + left + "," + right + "," + "t" + strconv.Itoa(variablesTemporales) + ")")
		return "t" + strconv.Itoa(variablesTemporales)
	case tt.Numero, tt.Decimal, tt.Bool, tt.Ident, tt.True, tt.False:
		return Nodo.TokenType.Lexema
	}
	return "ERROR CODE GEN"
}
