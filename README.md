# Monkey

This repository contains an implementation of the `Monkey` language interpreter developed in the book [Writing An Interpreter In Go](https://interpreterbook.com/).

## Example

The current implementation contains a REPL which parses user input and prints the parsed program.

```
$ monkey
Hello jamesroutley! This is the Monkey programming language!
Feel free to type in commands.
>> 5 + 5 * 10;
(5 + (5 * 10))
>>
```

## TODO:

The interpreter is a work in progress.

- [x] Define tokens
- [x] Lexer
- [x] Basic REPL
- [x] Parser
- [ ] Evaluator
