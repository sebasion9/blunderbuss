#!/bin/zsh
source ~/.zshrc
antlr4 blunderbuss.g4
javac blunderbuss*.java
grun blunderbuss program -gui test.bs
