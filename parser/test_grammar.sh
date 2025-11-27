#!/bin/zsh
source ~/.zshrc
antlr4 Blunderbuss.g4
javac Blunderbuss*.java
grun Blunderbuss program -gui ../bbuss/print.bbuss
