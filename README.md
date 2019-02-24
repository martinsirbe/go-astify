# Go ASTify
  
An Abstract Syntax Tree parser written in Go.
  
![code_coverage_badge](code_coverage.svg) ![json_ast_badge](json_ast.svg) 
  
The main goal of this project is to eventually provide a single library for all of your AST ([Abstract Syntax Tree][AST]) 
parsing needs. Currently only JSON parsing is supported, however feel free to contribute to this! 😉
  
## JSON
Parsing to the AST supports the following nodes based on the JSON syntax:  
- [x] Null
- [x] String
- [x] Integer
- [x] Boolean
- [x] Object
- [x] Array
  
### Structure
- **json/ast.go** - contains an AST node interface and JSON node definitions
- **json/tokenizer.go** - allows to scan the JSON input to produce a lexical representation
- **json/parser.go** - parses the JSON input to the AST by using the tokenizer
- **json/example/main.go** - contains an example of parsing the test JSON (`test_data/test.json`) to the AST
  
### How to use this?
First you will need to initialise a new tokenizer and to provide a JSON input as `[]rune`. Once the tokenizer has been 
successfully initialised, you then can initialise parser and call its `Parse` function which will return the AST Node.  
  
## Tests
To run tests, just run `make test`
```bash
go test -v -cover ./...
=== RUN   TestCorrectlyParsedJSONInputToAST
--- PASS: TestCorrectlyParsedJSONInputToAST (0.00s)
=== RUN   TestCorrectlyTokenizedTestJSON
--- PASS: TestCorrectlyTokenizedTestJSON (0.00s)
=== RUN   TestSuccessfullyPreviewedNextToken
--- PASS: TestSuccessfullyPreviewedNextToken (0.00s)
=== RUN   TestWhenPreviewingNextTokenDoesNotAdvanceToNextToken
--- PASS: TestWhenPreviewingNextTokenDoesNotAdvanceToNextToken (0.00s)
PASS
coverage: 88.6% of statements
ok      github.com/martinsirbe/go-astify/json   (cached)        coverage: 88.6% of statements
?       github.com/martinsirbe/go-astify/json/example   [no test files]
```
  
## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENCE.md) file for details
  
## Contributing
  
1. Go get it `go get github.com/martinsirbe/go-astify`
2. Create your feature branch (`git checkout -b my-feature-branch`)
3. Commit your changes (`git commit -m 'Add ...'`)
4. Push to the branch (`git push origin my-feature-branch`)
5. Create a new pull request
  
[AST]: https://en.wikipedia.org/wiki/Abstract_syntax_tree