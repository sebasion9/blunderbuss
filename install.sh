#!/usr/bin/env bash
set -euo pipefail

PREFIX=${PREFIX:-/opt/bbuss}

echo "[*] Building binaries and standard library..."
go build -o target/bbuss ./cmd
go build -buildmode=c-archive -ldflags=-linkmode=external -o target/libstd.a ./lib

echo "[*] Installing blunderbuss to $PREFIX"

install -d "$PREFIX/bin"
install -d "$PREFIX/lib"
install -d "$PREFIX/include/bbuss"


install -m 0755 target/bbuss.bin "$PREFIX/bin/"

install -m 0644 target/libstd.a "$PREFIX/lib"

cp bbuss/lib/* "$PREFIX/include/bbuss"

echo "[*] Installation complete"
echo "Add to path: export PATH=$PREFIX/bin:\$PATH"

