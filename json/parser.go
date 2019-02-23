package json

import (
	"strconv"

	"github.com/pkg/errors"
)

// Parser parses the given JSON input to an AST
type Parser struct {
	tokenizer *Tokenizer
}

// NewParser used to initialise a new instance of JSON AST Parser
func NewParser(tokenizer *Tokenizer) *Parser {
	return &Parser{
		tokenizer: tokenizer,
	}
}

// Parse used to parse the given JSON input (provided via tokenizer) to AST
func (p *Parser) Parse() (Node, error) {
	token := p.tokenizer.GetToken()
	switch token.Type {
	case STRING:
		return String{Value: token.Value}, nil
	case INTEGER:
		i, err := strconv.Atoi(token.Value)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to convert int string value - %s", token.Value)
		}
		return Integer{Value: i}, nil
	case BOOLEAN:
		b, err := strconv.ParseBool(token.Value)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to convert bool string value - %s", token.Value)
		}
		return Boolean{Value: b}, nil
	case LCBRACKET:
		return p.getObject()
	case LSBRACKET:
		return p.getArray()
	}
	return nil, nil
}

func (p *Parser) getObject() (Node, error) {
	object := make(map[string]Node)
LOOP:
	for {
		token := p.tokenizer.GetToken()
		key := token.Value

		switch {
		case token.Type == RCBRACKET:
			break LOOP
		case token.Type == COMMA:
			continue
		case p.tokenizer.GetToken().Type != COLON:
			return nil, errors.New("expected token to be a colon separator")
		}

		bn, err := p.Parse() // calls internally tokenizer GetToken func to recursively resolve JSON property value
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse JSON object value")
		}
		object[key] = bn
	}
	return Object{Properties: object}, nil
}

func (p *Parser) getArray() (Node, error) {
	token := p.tokenizer.PeekToken()

	if token.Type == RSBRACKET {
		return Array{Elements: nil}, nil
	}

	bns := make([]Node, 0)
	for {
		bn, err := p.Parse()
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse an element from the array")
		}
		bns = append(bns, bn)

		token = p.tokenizer.GetToken()
		switch {
		case token.Type == RSBRACKET:
			return Array{Elements: bns}, nil
		case token.Type != COMMA:
			return nil, errors.New("invalid token provided in the array")
		}
	}
}
