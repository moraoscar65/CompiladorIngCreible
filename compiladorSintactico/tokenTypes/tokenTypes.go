package tokenTypes

import (
	d "../definiciones"
)

const (
	Error       = iota
	Program     //1
	If          //2
	Else        //3
	Fi          //4
	Do          //5
	Until       //6
	While       //7
	Read        //8
	Write       //9
	Float       //10
	Int         //11
	Bool        //12
	Not         //13
	And         //14
	Or          //15
	Plus        //16
	Minus       //17
	Multi       //18
	Div         //19
	Poten       //20
	Menor       //21
	MenorIgual  //22
	Mayor       //23
	MayorIgual  //24
	Comparacion //25
	DiferenteDe //26
	Asig        //27
	PuntoComa   //28
	Coma        //29
	Punto       //30
	ParenIzq    //31
	ParenDer    //32
	LlaveIzq    //33
	LlaveDer    //34
	Ident       //35
	Numero      //36
	Decimal     //37
	True        //38
	False       //39
	Then        //40
	Eof         //44
)

func GetPalabrasReservadas() []d.Token {
	return []d.Token{
		{TypeLexema: Program, Lexema: "program"},
		{TypeLexema: If, Lexema: "if"},
		{TypeLexema: Then, Lexema: "then"},
		{TypeLexema: True, Lexema: "true"},
		{TypeLexema: False, Lexema: "false"},
		{TypeLexema: Else, Lexema: "else"},
		{TypeLexema: Fi, Lexema: "fi"},
		{TypeLexema: Do, Lexema: "do"},
		{TypeLexema: Until, Lexema: "until"},
		{TypeLexema: While, Lexema: "while"},
		{TypeLexema: Read, Lexema: "read"},
		{TypeLexema: Write, Lexema: "write"},
		{TypeLexema: Float, Lexema: "float"},
		{TypeLexema: Int, Lexema: "int"},
		{TypeLexema: Bool, Lexema: "bool"},
		{TypeLexema: Not, Lexema: "not"},
		{TypeLexema: And, Lexema: "and"},
		{TypeLexema: Or, Lexema: "or"},
	}

}
