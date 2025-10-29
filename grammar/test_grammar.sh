#!/bin/zsh
source ~/.zshrc
antlr4 main.g4
javac main*.java
grun main program -gui test.bs
