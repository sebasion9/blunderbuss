#!/bin/bash
set -e

if [ "$#" -ne 2 ]; then
    echo "usage: $0 <input.bbuss> <output>"
    exit 1
fi

filename="$1"

export TERM="xterm"
# preprocessor step and compile to asm
echo "[BBUSS] compiler step..."
# gcc -E -P -x c $1 | go run ./cmd
go run ./cmd --input $1 --output $2

echo "[BBUSS] building std lib..."
echo "[WARN] should remove this step for realese as it is static library"
go build -buildmode=c-archive -ldflags=-linkmode=external -o target/std.a ./lib

# to bin
echo "[BBUSS] writing to obj file..."
nasm -f elf64 $2 -o $2.o

echo "[BBUSS] linker step..."
# gcc -no-pie target/out.o -lncurses -o target/out
# gcc -no-pie target/out.o -lraylib -o target/out
# gcc -no-pie target/out.o bbuss/lib/raylib.o target/std.a  -lraylib -lm -ldl -lpthread -lGL -lfreetype -lrt -lX11 -o target/out
gcc -no-pie $2.o target/std.a -lncurses -o $2.bin

echo "[BBUSS] running executable..."
./$2.bin


