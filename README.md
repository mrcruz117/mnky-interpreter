# mnky-interpreter

An interpreter written in Go using only the standard library. Thank you Thorsten Ball!

## Current Step

- 5 - Macro system

## Features

### Core Features

- **Arithmetic Operations**: Supports basic arithmetic operations like addition (`+`), subtraction (`-`), multiplication (`*`), and division (`/`).
- **Boolean Logic**: Includes support for boolean values (`true`, `false`) and logical operators (`!`, `==`, `!=`, `<`, `>`).
- **Variable Declarations**: Allows variable declarations using the `let` keyword (e.g., `let x = 5;`).
- **Return Statements**: Supports returning values from blocks or functions using the `return` keyword.
- **Conditionals**: Includes `if` and `if-else` expressions for conditional execution of code.
- **First-Class Functions**: Functions are first-class citizens, meaning they can be assigned to variables, passed as arguments, and returned from other functions.
- **Closures**: Functions can capture and use variables from their surrounding environment.
- **Built-in Functions**: Includes built-in functions like:
  - `len`: Returns the length of a string or array (e.g., `len("hello")` = `5`).
  - `first`: Returns the first element of an array (e.g., `first([1, 2, 3])` = `1`).
  - `last`: Returns the last element of an array (e.g., `last([1, 2, 3])` = `3`).
  - `rest`: Returns all elements of an array except the first (e.g., `rest([1, 2, 3])` = `[2, 3]`).
  - `push`: Adds an element to the end of an array (e.g., `push([1, 2], 3)` = `[1, 2, 3]`).
  - `puts`: Prints arguments to the console.

### Data Types

- **Integers**: Supports 64-bit integer values.
- **Strings**: Supports string literals and concatenation.
- **Booleans**: Includes `true` and `false` as boolean values.
- **Null**: Represents the absence of a value with `null`.
- **Arrays**: Supports array literals (e.g., `[1, 2, 3]`) and indexing (e.g., `myArray[0]`).
- **Hashes**: Supports hash literals with string, integer, and boolean keys (e.g., `{"key": "value", 1: true}`).

### Expressions

- **Prefix Expressions**: Supports prefix operators like `!` (logical NOT) and `-` (negation).
- **Infix Expressions**: Supports infix operators for arithmetic, comparison, and string concatenation (e.g., `5 + 5`, `1 < 2`, `"Hello" + " World!"`).
- **Grouped Expressions**: Allows grouping of expressions using parentheses (e.g., `(1 + 2) * 3`).
- **Index Expressions**: Supports accessing elements in arrays or hashes using index expressions (e.g., `myArray[1 + 1]`, `myHash["key"]`).

### Statements

- **Let Statements**: Allows variable declarations (e.g., `let x = 10;`).
- **Expression Statements**: Evaluates expressions as standalone statements (e.g., `5 + 5;`).
- **Return Statements**: Returns a value from a block or function (e.g., `return 10;`).

### Functions

- **Function Literals**: Define functions inline using the `fn` keyword (e.g., `fn(x, y) { x + y; }`).
- **Function Calls**: Call functions with arguments (e.g., `add(1, 2);`).
- **Closures**: Functions can capture variables from their surrounding scope (e.g., `let adder = fn(x) { fn(y) { x + y; }; };`).

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
// Variable declarations
let x = 10;
let y = 20;

// Arithmetic operations
let sum = x + y;

// Conditional statements
if (sum > 20) {
    return true;
} else {
    return false;
}

// Arrays and indexing
let myArray = [1, 2, 3];
let firstElement = myArray[0];

// Hashes
let myHash = {"key": "value", 1: true};
let hashValue = myHash["key"];

// Functions and closures
let add = fn(a, b) {
    return a + b;
};
let result = add(5, 10);

// Built-in functions
len(myArray); // Outputs: 3
first(myArray); // Outputs: 1
push(myArray, 4); // Outputs: [1, 2, 3, 4]
```

## TODO

- Add support for loops and iteration.
- Implement macros for advanced metaprogramming.
- Add more built-in functions.
