package tokens

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	//instructions
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	LABEL   = "LABEL"
	INT     = "INT"
	STRING  = "STRING"
	COMMA   = "COMMA"

	STORE = "STORE"

	ADD = "ADD"
	DES = "DES"
	DIV = "DIV"
	INC = "INC"
	MUL = "MUL"
	SUB = "SUB"
)

var keywords = map[string]Type{
	"store": STORE,
	"add":   ADD,
	"sub":   SUB,
	"mul":   MUL,
	"div":   DIV,
	"inc":   INC,
	"des":   DES,
}

func LookupIdentifier(identifier string) Type {
	if token, ok := keywords[identifier]; ok {
		return token
	}
	return IDENT
}
