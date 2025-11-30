#!/bin/bash


if [ "$#" -ne 1 ]; then
    echo "usage: $0 <file.bbuss>"
    exit 1
fi

filename="$1"

export TERM="xterm"
# preprocessor step and compile to asm
gcc -E -P -x c $1 | go run .

# to bin
nasm -f elf64 target/out.asm -o target/out.o
# gcc -no-pie target/out.o -lncurses -o target/out
# gcc -no-pie target/out.o -lraylib -o target/out
gcc -no-pie target/out.o bbuss/lib/raylib.o  -lraylib -lm -ldl -lpthread -lGL -lfreetype -lrt -lX11 -o target/out

./target/out


