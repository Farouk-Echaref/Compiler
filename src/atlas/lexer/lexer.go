package lexer

import (
	"atlas/token"
)

//Lexer struct 
type Lexer struct {
	//store source code
	source string
	//pos of the current character 
	pos int
	//pos of reading char(next pos)
	readPos int
	//current char
	char byte
}

//create Lexer
func New(input string) *Lexer{
	lxr := &Lexer{
		source: input
	}

	return lxr
} 