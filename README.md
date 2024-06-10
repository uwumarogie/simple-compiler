# Simple Compiler in Go

This repository contains the source code for a simple compiler implemented in Go. The purpose of this project is to demonstrate the basic components of a compiler, including lexical analysis, parsing, and evaluation.

## Project Structure

- **ast.go**: Contains the Abstract Syntax Tree (AST) nodes and related structures.
- **evaluator.go**: Handles the evaluation of the parsed AST.
- **go.mod**: Go module file that specifies the module path and dependencies.
- **lexer.go**: Implements the lexical analysis, converting source code into tokens.
- **main.go**: The entry point of the compiler, coordinating the compilation process.
- **parser.go**: Parses tokens into an AST.
- **token.go**: Defines the tokens used by the lexer and parser.

## Getting Started

### Prerequisites

- Go 1.16 or higher installed on your machine.

### Installation

Clone the repository:

```bash
git clone https://github.com/uwumarogie/simple_compiler.git
cd simple_compiler
```

### Usage

To run the compiler, use the following command:

```bash
 go run main.go
```
This will execute the main.go file, which initiates the compilation process.


## Components

### Lexer
The lexer (lexer.go) reads the source code and breaks it into tokens. Tokens are the smallest units of meaningful text, such as keywords, identifiers, operators, and literals.

### Parser
The parser (parser.go) takes the tokens produced by the lexer and organizes them into a hierarchical structure known as an Abstract Syntax Tree (AST). The AST represents the grammatical structure of the source code.

### Evaluator
The evaluator (evaluator.go) processes the AST to perform the necessary computations or actions specified by the source code.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Initial Motivation

I created this project as a learning exercise to understand the basic components of a compiler and how they work together after listening to a theoretical computer science lecture on compilers.
