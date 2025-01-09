# Luma Parser

A Go-based parser for processing UML-like class diagrams and relationships. The parser tokenizes and parses a custom language syntax for defining classes, interfaces, and their relationships.

## Features

- **Class Definitions**: Define classes with fields and methods
- **Interface Definitions**: Create interface declarations
- **Relationships**: Support for various UML relationships:
  - Associations (-->, <--, <-->)
  - Dependencies (..>, <..)
  - Inheritance (--|>, <|--)
  - Implementation (..|>, <|..)
  - Aggregation (--<>, <>--)
  - Composition (--*, *--)
- **Visibility Modifiers**: Support for standard visibility notation:
  - Public (+)
  - Private (-)
  - Protected (#)
  - Package (~)
- **Comments**: Support for single-line comments

## Language Syntax

### Class Definition
```
class MyClass {
    - privateField: string
    + publicMethod(param: string): void
}
```

### Relationships
```
ClassA --> ClassB                    // Simple association
ClassA "1" --> "many" ClassB         // Association with cardinality
ClassA --> ClassB: "contains"        // Association with label
```

## Project Structure

- `lexer/`: Contains the lexical analyzer
  - `lexer.go`: Core lexer implementation
  - `tokens.go`: Token definitions and types
  - `lexerHelperFunctions.go`: Helper functions for character classification
- `parser/`: Contains the parser implementation
  - `parser.go`: Core parser implementation
  - `ast.go`: Abstract Syntax Tree definitions
- `main.go`: Entry point of the parser

## Usage

### Basic Usage

1. Create a file with your class diagram definitions (e.g., `example.lang`)
2. Run the parser:
```bash
go run main.go
```

The parser will process the input file and output a JSON representation of the parsed Abstract Syntax Tree.

### Example Input

```
class Person {
    - name: string
    - age: int
    + getName(): string
    + setName(name: string): void
}

class Employee {
    - salary: float
}

Person <|-- Employee
```

## Building

To build the project:

```bash
go build
```

## Testing

The project includes comprehensive tests for both lexer and parser components. To run the tests:

```bash
go test ./...
```

## Error Handling

The parser provides detailed error messages for:
- Invalid syntax
- Unknown tokens
- Malformed class definitions
- Invalid relationships
- Missing closing brackets/parentheses
- Invalid parameter definitions

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the MIT License.