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
go run ./cmd --input $1 --output $2 --include ./bbuss/lib

echo "[BBUSS] building std lib..."
echo "[WARN] should remove this step for realese as it is static library"
go build -buildmode=c-archive -ldflags=-linkmode=external -o target/libstd.a ./lib

# to bin
echo "[BBUSS] writing to obj file..."
nasm -f elf64 $2 -o $2.o

echo "[BBUSS] linker step..."

#raylib
#gcc -no-pie $2.o -L./bbuss/lib target/libstd.a bbuss/lib/raylib.o -lraylib -lm -ldl -lpthread -lGL -lfreetype -lrt -lX11 -o $2.bin

# ncurses
gcc -no-pie $2.o -L./bbuss/lib target/libstd.a -lncurses -o $2.bin

# standard
#gcc -no-pie $2.o -L./bbuss/lib target/libstd.a -o $2.bin

echo "[BBUSS] running executable..."
./$2.bin


