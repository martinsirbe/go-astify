package json_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/martinsirbe/go-astify/json"
)

func TestCorrectlyTokenizedTestJSON(t *testing.T) {
	jsonFile, err := os.Open("../test_data/test.json")
	defer func() {
		assert.Nil(t, jsonFile.Close())
	}()
	assert.Nil(t, err)

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	assert.Nil(t, err)

	tokenizer := json.NewTokenizer([]rune(string(jsonBytes)))
	for _, expectedToken := range expectedTokens() {
		token := tokenizer.GetToken()
		assert.Equal(t, expectedToken.Type, token.Type)
		assert.Equal(t, expectedToken.Value, token.Value)
	}
}

func TestSuccessfullyPreviewedNextToken(t *testing.T) {
	jsonFile, err := os.Open("../test_data/test.json")
	defer func() {
		assert.Nil(t, jsonFile.Close())
	}()
	assert.Nil(t, err)

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	assert.Nil(t, err)

	tokenizer := json.NewTokenizer([]rune(string(jsonBytes)))
	nextToken := tokenizer.PeekToken()
	assert.Equal(t, "{", nextToken.Value)
	assert.Equal(t, json.LCBRACKET, nextToken.Type)
}

func TestWhenPreviewingNextTokenDoesNotAdvanceToNextToken(t *testing.T) {
	jsonFile, err := os.Open("../test_data/test.json")
	defer func() {
		assert.Nil(t, jsonFile.Close())
	}()
	assert.Nil(t, err)

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	assert.Nil(t, err)

	tokenizer := json.NewTokenizer([]rune(string(jsonBytes)))
	nextToken := tokenizer.PeekToken()
	assert.Equal(t, "{", nextToken.Value)
	assert.Equal(t, json.LCBRACKET, nextToken.Type)

	nextToken = tokenizer.PeekToken()
	assert.Equal(t, "{", nextToken.Value)
	assert.Equal(t, json.LCBRACKET, nextToken.Type)
}

func expectedTokens() []json.Token {
	return []json.Token{
		{
			Value: "{",
			Type: json.LCBRACKET,
		},
		{
			Value: "test_object",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "{",
			Type: json.LCBRACKET,
		},
		{
			Value: "test_string",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "hello world",
			Type: json.STRING,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test_integer",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "1234567890",
			Type: json.INTEGER,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test_boolean",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "true",
			Type: json.BOOLEAN,
		},
		{
			Value: "}",
			Type: json.RCBRACKET,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test_string",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "hello world",
			Type: json.STRING,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test_integer",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "1234567890",
			Type: json.INTEGER,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test_boolean",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "true",
			Type: json.BOOLEAN,
		},
		{
			Value: "test_string_array",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "[",
			Type: json.LSBRACKET,
		},
		{
			Value: "test 1",
			Type: json.STRING,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test 2",
			Type: json.STRING,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test 3",
			Type: json.STRING,
		},
		{
			Value: "]",
			Type: json.RSBRACKET,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test_object_array",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "[",
			Type: json.LSBRACKET,
		},
		{
			Value: "{",
			Type: json.LCBRACKET,
		},
		{
			Value: "test_string",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "hello world 1",
			Type: json.STRING,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test_integer",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "123",
			Type: json.INTEGER,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test_boolean",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "false",
			Type: json.BOOLEAN,
		},
		{
			Value: "}",
			Type: json.RCBRACKET,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "{",
			Type: json.LCBRACKET,
		},
		{
			Value: "test_string",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "hello world 2",
			Type: json.STRING,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test_integer",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "456",
			Type: json.INTEGER,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test_boolean",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "true",
			Type: json.BOOLEAN,
		},
		{
			Value: "}",
			Type: json.RCBRACKET,
		},
		{
			Value: "]",
			Type: json.RSBRACKET,
		},
		{
			Value: ",",
			Type: json.COMMA,
		},
		{
			Value: "test_empty_array",
			Type: json.STRING,
		},
		{
			Value: ":",
			Type: json.COLON,
		},
		{
			Value: "[",
			Type: json.LSBRACKET,
		},
		{
			Value: "]",
			Type: json.RSBRACKET,
		},
		{
			Value: "}",
			Type: json.RCBRACKET,
		},
	}
}
