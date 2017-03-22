# Monkey

This repository contains an implementation of the `Monkey` language interpreter developed in the book [Writing An Interpreter In Go](https://interpreterbook.com/).

## Example

The current implementation contains a REPL which lexes user input and prints a list of the lexed tokens.

```
$ monkey
Hello jamesroutley! This is the Monkey programming language!
Feel free to type in commands.
>> 10 + 10;
{Type:INT Literal:10}
{Type:+ Literal:+}
{Type:INT Literal:10}
{Type:; Literal:;}
>>
```

## TODO:

The interpreter is a work in progress.

- [x] Define tokens
- [x] Lexer
- [x] Basic REPL
- [ ] Parser
- [ ] Evaluator
