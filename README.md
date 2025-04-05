## blunderbuss – A Simple Compiler in C++

A minimal compiler with custom, C-like syntax from scratch using C++.  
At the moment, it includes:

- A lexer that tokenizes source code
- A recursive descent parser that builds an abstract syntax tree (AST)


## Building and Running

### 🔧 Build the Compiler (Release)

```bash
make
```

### Build in Debug mode (pipe to gdb)

```bash
make debug
```

### Build and run

```bash
make run
```

## Run tests

All unit tests live in `src/test/` and are compiled and ran:

```
make test
```

## Clean

```bash
make clean
```

