# mnky-interpreter

An interpreter written in Go using only the standard library. Thank you Thorsten Ball!

## Current Step

- 4.6 - Finale

- p.197

## Features

### Core Features

- **Arithmetic Operations**: Supports basic arithmetic operations like addition (`+`), subtraction (`-`), multiplication (`*`), and division (`/`).
- **Boolean Logic**: Includes support for boolean values (`true`, `false`) and logical operators (`!`, `==`, `!=`, `<`, `>`).
- **Variable Declarations**: Allows variable declarations using the `let` keyword (e.g., `let x = 5;`).
- **Return Statements**: Supports returning values from blocks or functions using the `return` keyword.
- **Conditionals**: Includes `if` and `if-else` expressions for conditional execution of code.
- **First-Class Functions**: Functions are first-class citizens, meaning they can be assigned to variables, passed as arguments, and returned from other functions.
- **Closures**: Functions can capture and use variables from their surrounding environment.

### Data Types

- **Integers**: Supports 64-bit integer values.
- **Strings**: Supports string literals and concatenation.
- **Booleans**: Includes `true` and `false` as boolean values.
- **Null**: Represents the absence of a value with `null`.
- **Arrays**: Supports array literals (e.g., `[1, 2, 3]`) and indexing (e.g., `myArray[0]`).

### Expressions

- **Prefix Expressions**: Supports prefix operators like `!` (logical NOT) and `-` (negation).
- **Infix Expressions**: Supports infix operators for arithmetic, comparison, and string concatenation (e.g., `5 + 5`, `1 < 2`, `"Hello" + " World!"`).
- **Grouped Expressions**: Allows grouping of expressions using parentheses (e.g., `(1 + 2) * 3`).
- **Index Expressions**: Supports accessing elements in arrays using index expressions (e.g., `myArray[1 + 1]`).

### Statements

- **Let Statements**: Allows variable declarations (e.g., `let x = 10;`).
- **Expression Statements**: Evaluates expressions as standalone statements (e.g., `5 + 5;`).
- **Return Statements**: Returns a value from a block or function (e.g., `return 10;`).

### Functions

- **Function Literals**: Define functions inline using the `fn` keyword (e.g., `fn(x, y) { x + y; }`).
- **Function Calls**: Call functions with arguments (e.g., `add(1, 2);`).
- **Closures**: Functions can capture variables from their surrounding scope (e.g., `let adder = fn(x) { fn(y) { x + y; }; };`).
- **Built-in Functions**: Includes built-in functions like:
  - `len`: Returns the length of a string or array (e.g., `len("hello")` = `5`).

### Scoping

- **Lexical Scoping**: Variables are scoped to the block or function in which they are defined.
- **Nested Environments**: Supports nested environments for managing variable scopes in functions and blocks.

### Error Handling

- **Runtime Errors**: Provides error messages for invalid operations (e.g., type mismatches, undefined variables, or unsupported operations).

### REPL (Read-Eval-Print Loop)

- Interactive REPL for executing Mnky code line by line.
- Displays evaluated results or error messages.

## Example Code

```mnky
let add = fn(x, y) {
    x + y;
};

let result = add(5, 10);
result; // Outputs: 15

if (result > 10) {
    return true;
} else {
    return false;
}

let myArray = [1, 2, 3];
myArray[0]; // Outputs: 1
```

## TODO

- Add tracing functions for parser debugging.
- Extend support for additional data types (e.g., hash maps).
- Add support for loops and iteration.
- Implement more built-in functions.
