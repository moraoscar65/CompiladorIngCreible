package sintactico

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	d "../definiciones"
	lexico "../lexico"
	lg "../logger"
	"../tokenTypes"
)

var token d.Token
var r *bufio.Reader

func Parse(file string) *d.Nodo {
	if text, err := (lexico.ReadFile(file)); text != "" { //cargamos el archivo dentro de un string y validamos su apertura
		r = bufio.NewReader(strings.NewReader(text)) //creamos el buffer para ir obteniendo caracter por caracter
		token = lexico.GetToken(r)
		var tree = Inicializar()
		tree = program()
		lg.AgregarTree(*tree)

		return tree
	} else {
		fmt.Println("File reading error", err)
	}
	return nil
}

func program() *d.Nodo {
	var tmp = Inicializar()
	if token.TypeLexema == tokenTypes.Program {
		tmp.TokenType = token
		match(tokenTypes.Program)
		match(tokenTypes.LlaveIzq)
		tmp.Izq = listaDeclaracion()
		tmp.Der = listaSentencias()
		match(tokenTypes.LlaveDer)
	}
	return tmp
}

func match(tokenType int) bool {
	var err bool = false
	if token.TypeLexema == tokenType {
		tmp := lexico.GetToken(r)

		lg.AgregarSintactico("Match: " + token.Lexema + " (" + strconv.Itoa(token.NumFila) + "," + strconv.Itoa(token.NumCol) + ")")
		token = tmp

	} else {
		str := "Error en" + " (" + strconv.Itoa(token.NumFila) + "," + strconv.Itoa(token.NumCol) + ")"
		lg.AgregarError(str)
		err = true
	}
	return err
}

func PanicMode(tokenType int) {
	for token.TypeLexema != tokenType {
		tmp := lexico.GetToken(r)
		lg.AgregarSintactico("Token Ignorado: " + token.Lexema)
		token = tmp

	}
	match(tokenType)
}

func identificador() (*d.Nodo, error) {
	root := Inicializar()
	root.TokenType = token

	var err error = nil
	if match(tokenTypes.Ident) {
		root.TokenType.TypeLexema = tokenTypes.Error
		err = errors.New("Error en el identificador")
	}
	return root, err
}

func Inicializar() *d.Nodo {

	nodo := d.Nodo{TokenType: d.Token{TypeLexema: 0, Lexema: "", NumFila: 0, NumCol: 0}, AtrValor: "", Izq: nil, Med: nil, Der: nil, Bro: nil}
	return &nodo

}

func numero() (*d.Nodo, error) {
	root := Inicializar()
	root.TokenType = token
	root.Dtype = tokenTypes.Int
	var err error = nil
	if match(tokenTypes.Numero) {
		root.TokenType.TypeLexema = tokenTypes.Error
		err = errors.New("Error en la constante numerica entera")
	}
	return root, err
}

func decimal() (*d.Nodo, error) {
	root := Inicializar()
	root.TokenType = token
	root.Dtype = tokenTypes.Float
	var err error = nil
	if match(tokenTypes.Decimal) {
		root.TokenType.TypeLexema = tokenTypes.Error
		err = errors.New("Error en la constante numerica decimal")
	}
	return root, err
}

func TreeToString(node *d.Nodo) string {
	var str string
	jsonData, err := json.MarshalIndent(node, "", "    ")
	if err != nil {
		fmt.Println(err)
	} else {
		str = string(jsonData)
		str = strings.ReplaceAll(str, "\\u003c", "<")
		str = strings.ReplaceAll(str, "\\u003e", ">")
		str = strings.ReplaceAll(str, "null", "[]")
	}
	return str
}
