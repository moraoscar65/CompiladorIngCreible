package definiciones

type Token struct {
	TypeLexema int
	Lexema     string
	NumFila    int
	NumCol     int
}

type Nodo struct {
	TokenType Token `json:"TokenType,omitempty"`
	Dtype     int
	AtrValor  string `json:"Name,omitempty"`
	Izq       *Nodo  `json:"Izq,omitempty"`
	Med       *Nodo  `json:"Med,omitempty"`
	Der       *Nodo  `json:"Der,omitempty"`
	Bro       *Nodo  `json:"Bro,omitempty"`
}

type Simbolo struct {
	Dtype int
	Valor string
}
