package json_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/martinsirbe/go-astify/json"
)

func TestCorrectlyParsedJSONInputIntoAST(t *testing.T) {
	jsonFile, err := os.Open("../test_data/test.json")
	defer func() {
		assert.Nil(t, jsonFile.Close())
	}()
	assert.Nil(t, err)

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	assert.Nil(t, err)

	tokenizer := json.NewTokenizer([]rune(string(jsonBytes)))
	parser := json.NewParser(tokenizer)

	ast, err := parser.Parse()
	assert.Nil(t, err)

	assert.IsType(t, json.Object{}, ast.Get())
	rootObjProps := ast.Get().(json.Object)
	assert.Equal(t, 6, len(rootObjProps.Properties))

	testObject(t, rootObjProps)
	testString(t, rootObjProps)
	testInteger(t, rootObjProps)
	testBoolean(t, rootObjProps)
	testSimpleArray(t, rootObjProps)
	testObjectArray(t, rootObjProps)
}

func testObject(t *testing.T, rootObjProps json.Object) {
	obj := rootObjProps.Properties["test_object"]
	assert.IsType(t, json.Object{}, obj.Get())
	simpleObjProps := obj.Get().(json.Object).Properties
	assert.Equal(t, 3, len(simpleObjProps))

	objStr := simpleObjProps["test_string"]
	assert.IsType(t, json.String{}, objStr.Get())
	assert.Equal(t, "hello world", objStr.Get().(json.String).Value)

	objInt := simpleObjProps["test_integer"]
	assert.IsType(t, json.Integer{}, objInt.Get())
	assert.Equal(t, 1234567890, objInt.Get().(json.Integer).Value)

	objBoolTrue := simpleObjProps["test_boolean"]
	assert.IsType(t, json.Boolean{}, objBoolTrue.Get())
	assert.Equal(t, true, objBoolTrue.Get().(json.Boolean).Value)
}

func testString(t *testing.T, rootObjProps json.Object) {
	s := rootObjProps.Properties["test_string"]
	assert.IsType(t, json.String{}, s.Get())
	assert.Equal(t, "hello world", s.Get().(json.String).Value)
}

func testInteger(t *testing.T, rootObjProps json.Object) {
	i := rootObjProps.Properties["test_integer"]
	assert.IsType(t, json.Integer{}, i.Get())
	assert.Equal(t, 1234567890, i.Get().(json.Integer).Value)
}

func testBoolean(t *testing.T, rootObjProps json.Object) {
	b := rootObjProps.Properties["test_boolean"]
	assert.IsType(t, json.Boolean{}, b.Get())
	assert.Equal(t, true, b.Get().(json.Boolean).Value)
}

func testSimpleArray(t *testing.T, rootObjProps json.Object) {
	a := rootObjProps.Properties["test_string_array"]
	assert.IsType(t, json.Array{}, a.Get())

	av := a.Get().(json.Array).Elements
	assert.Equal(t, 3, len(av))

	for i := 0; i < 3; i++ {
		s := av[i]
		assert.IsType(t, json.String{}, s.Get())
		assert.Equal(t, fmt.Sprintf("test %d", i+1), s.Get().(json.String).Value)
	}
}

func testObjectArray(t *testing.T, rootObjProps json.Object) {
	a := rootObjProps.Properties["test_object_array"]
	assert.IsType(t, json.Array{}, a.Get())

	av := a.Get().(json.Array).Elements
	assert.Equal(t, 2, len(av))

	for _, expected := range []struct {
		index        int
		stringValue  string
		integerValue int
		booleanValue bool
	}{
		{
			index:        0,
			stringValue:  "hello world 1",
			integerValue: 123,
			booleanValue: false,
		},
		{
			index:        1,
			stringValue:  "hello world 2",
			integerValue: 456,
			booleanValue: true,
		},
	} {
		s := av[expected.index]
		assert.IsType(t, json.Object{}, s.Get())

		o := s.(json.Object).Properties

		objStr := o["test_string"]
		assert.IsType(t, json.String{}, objStr.Get())
		assert.Equal(t, expected.stringValue, objStr.Get().(json.String).Value)

		objInt := o["test_integer"]
		assert.IsType(t, json.Integer{}, objInt.Get())
		assert.Equal(t, expected.integerValue, objInt.Get().(json.Integer).Value)

		objBool := o["test_boolean"]
		assert.IsType(t, json.Boolean{}, objBool.Get())
		assert.Equal(t, expected.booleanValue, objBool.Get().(json.Boolean).Value)
	}
}
