#!/bin/bash


if [ "$#" -ne 1 ]; then
    echo "usage: $0 <file.bbuss>"
    exit 1
fi

filename="$1"

export TERM="xterm"
# preprocessor step and compile to asm
echo "[BBUSS] compiler step..."
gcc -E -P -x c $1 | go run ./cmd

echo "[BBUSS] building std lib..."
echo "[WARN] should remove this step for realese as it is static library"
go build -buildmode=c-archive -ldflags=-linkmode=external -o target/std.a ./lib

# to bin
echo "[BBUSS] writing to obj file..."
nasm -f elf64 target/out.asm -o target/out.o

echo "[BBUSS] linker step..."
# gcc -no-pie target/out.o -lncurses -o target/out
# gcc -no-pie target/out.o -lraylib -o target/out
# gcc -no-pie target/out.o bbuss/lib/raylib.o  -lraylib -lm -ldl -lpthread -lGL -lfreetype -lrt -lX11 -o target/out

gcc -no-pie target/out.o target/std.a -o target/out -lpthread -lm -ldl

echo "[BBUSS] running executable..."
./target/out


