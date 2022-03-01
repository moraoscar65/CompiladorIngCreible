package lexico

import (
	"bufio"
	
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"unicode"

	d "../definiciones"
	"../logger"
	s "../states"
	tt "../tokenTypes"
)

var palabrasReservadas = tt.GetPalabrasReservadas()
var nCol int = 0
var nFila int = 1

func loggers(token d.Token) {
	if token.TypeLexema == tt.Error {
		logger.AgregarError("Token no identificado (" + strconv.Itoa(token.NumFila) + "," + strconv.Itoa(token.NumCol) + "): " + token.Lexema)
		//fmt.Println("Token no identificado (" + strconv.Itoa(token.Row) + "," + strconv.Itoa(token.Col) + "): " + token.Lexema)
	} else {
		logger.AgregarLexico("Token (" + strconv.Itoa(token.NumFila) + "," + strconv.Itoa(token.NumCol) + "): " + token.Lexema)
		//fmt.Println("Token (" + strconv.Itoa(row) + "," + strconv.Itoa(index) + "): " + token.Lexema)
	}
}

func ReadFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {

		return "", err
	}
	text := string(data)
	return text, nil
}

func unGetRune(r *bufio.Reader) {
	r.UnreadRune()
}

func getRune(r *bufio.Reader) string {

	if c, _, err := r.ReadRune(); err != nil {
		if err == io.EOF {
			//fmt.Printf("%q [%d]\n", string(c), sz)
			//fmt.Println(c)
			//fmt.Println("----------final del archivo 2-----------")
			return "EOF"
		} else {
			log.Fatal(err)
		}
	} else {
		//fmt.Printf("%q [%d]\n", string(c), sz)
		nCol++
		if (string(c)) == "\n" {
			nCol = 0
			nFila++
		}
		return string(c)
	}
	return "ERROR"
}

func isDelimi(c string) bool {
	if c == " " || c == "\t" || c == "\n" || c == "\r" {
		return true
	}
	return false
}
func GetToken(r *bufio.Reader) d.Token {

	var c string
	state := s.InicialState
	var token d.Token
	for state != s.DoneState {
		switch state {

		case s.InicialState:

			c = getRune(r)

			for isDelimi(c) { // mientras es delimitador
				c = getRune(r) // extrae el siguiente caracter
			}

			if c == "EOF" {
				token.TypeLexema = tt.Eof
				state = s.DoneState
				token.Lexema = c
			} else if IsLetter(c) {
				state = s.IdentState
				token.Lexema += c
			} else if _, err := strconv.Atoi(c); err == nil {
				state = s.NumeroState
				token.Lexema += c
			} else if c == "(" {
				token.TypeLexema = tt.ParenIzq
				state = s.DoneState
				token.Lexema += c
			} else if c == ")" {
				token.TypeLexema = tt.ParenDer
				state = s.DoneState
				token.Lexema += c
			} else if c == "{" {
				token.TypeLexema = tt.LlaveIzq
				state = s.DoneState
				token.Lexema += c
			} else if c == "}" {
				token.TypeLexema = tt.LlaveDer
				state = s.DoneState
				token.Lexema += c
			} else if c == "," {
				token.TypeLexema = tt.Coma
				state = s.DoneState
				token.Lexema += c
			} else if c == "=" {
				state = s.CompState
				token.Lexema += c
			} else if c == "!" {
				state = s.DifDeState
				token.Lexema += c
			} else if c == "+" {
				token.TypeLexema = tt.Plus
				state = s.DoneState
				token.Lexema += c
			} else if c == "-" {
				token.TypeLexema = tt.Minus
				state = s.DoneState
				token.Lexema += c
			} else if c == "*" {
				token.TypeLexema = tt.Multi
				state = s.DoneState
				token.Lexema += c
			} else if c == "/" {
				state = s.ComentarioState
			} else if c == "^" {
				token.TypeLexema = tt.Poten
				state = s.DoneState
				token.Lexema += c
			} else if c == "<" {
				token.TypeLexema = tt.Menor
				state = s.MenorState
				token.Lexema += c
			} else if c == ">" {
				token.TypeLexema = tt.Mayor
				state = s.MayorState
				token.Lexema += c
			} else if c == ";" {
				token.TypeLexema = tt.PuntoComa
				state = s.DoneState
				token.Lexema += c

			} else if c == "." {
				state = s.DecimalState
				token.Lexema += c

			} else {
				token.TypeLexema = tt.Error
				state = s.DoneState
			}
		case s.ComentarioState:
			c = getRune(r)
			if c == "*" {
				state = s.ComentarioBloqueState
			} else if c == "/" {
				state = s.ComentarioLineaState
			} else {
				unGetRune(r)
				token.TypeLexema = tt.Div
				state = s.DoneState
				token.Lexema += "/"
			}

		case s.ComentarioLineaState:
			c = getRune(r)
			for !isSalto(c) {
				c = getRune(r)
			}
			state = s.InicialState

		case s.ComentarioBloqueState:
			c = getRune(r)
			for !isAsterisco(c) {
				c = getRune(r)
			}
			c = getRune(r)
			if c == "/" {
				state = s.InicialState
			} else {
				state = s.ComentarioBloqueState
			}

		case s.NumeroState:
			c = getRune(r)
			token.Lexema += c
			if _, err := strconv.Atoi(c); err != nil {
				if c == "." {
					state = s.DecimalState
				} else {
					token.TypeLexema = tt.Numero
					state = s.DoneState
					unGetRune(r)
					last := len(token.Lexema) - 1
					token.Lexema = token.Lexema[:last]
				}
			}
		case s.DecimalState:
			c = getRune(r)
			token.Lexema += c
			if _, err := strconv.Atoi(c); err != nil {

				token.TypeLexema = tt.Decimal
				state = s.DoneState
				unGetRune(r)
				last := len(token.Lexema) - 1
				token.Lexema = token.Lexema[:last]
			}
		case s.IdentState:
			c = getRune(r)
			token.Lexema += c
			if _, err := strconv.Atoi(c); !(IsLetter(c) || c == "_" || err == nil) {
				token.TypeLexema = tt.Ident
				state = s.DoneState
				unGetRune(r)
				last := len(token.Lexema) - 1
				token.Lexema = token.Lexema[:last]
				token = lookUpReserverdWord(token.Lexema)
			}
		case s.MenorState:
			c = getRune(r)
			if c == "=" {
				token.Lexema += c
				token.TypeLexema = tt.MenorIgual
				state = s.DoneState
			} else {
				unGetRune(r)
				token.TypeLexema = tt.Menor
				state = s.DoneState
			}
		case s.MayorState:
			c = getRune(r)
			if c == "=" {
				token.Lexema += c
				token.TypeLexema = tt.MayorIgual
				state = s.DoneState
			} else {
				unGetRune(r)
				token.TypeLexema = tt.Mayor
				state = s.DoneState
			}
		case s.CompState:
			c = getRune(r)
			if c == "=" {
				token.Lexema += c
				token.TypeLexema = tt.Comparacion
				state = s.DoneState
			} else {
				unGetRune(r)
				token.TypeLexema = tt.Asig
				state = s.DoneState
			}
		case s.DifDeState:
			c = getRune(r)
			if c == "=" {
				token.Lexema += c
				token.TypeLexema = tt.DiferenteDe
				state = s.DoneState
			}
		default:
			token.TypeLexema = tt.Error
			state = s.DoneState
			token.Lexema += c

		}
	}
	token.NumFila = nFila
	token.NumCol = nCol
	loggers(token)

	return token
}

func isSalto(s string) bool {
	if s == "\n" || s == "\r" {
		return true
	}
	return false
}
func isAsterisco(s string) bool {
	if s == "*" {
		return true
	}
	return false
}

func lookUpReserverdWord(s string) d.Token {

	var tok d.Token

	for index, element := range palabrasReservadas {
		if element.Lexema == s {
			
			return palabrasReservadas[index]
		}
	}
	tok.Lexema = s
	tok.TypeLexema = tt.Ident
	return tok
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
