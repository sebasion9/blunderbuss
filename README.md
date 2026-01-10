# blunderbuss

## Usage

### Requirements for usage

- compiler binary
- standard library

### CLI

```
./bbuss --input <source.bbuss> --output <out.asm>
nasm -f elf64 <out.asm> -o <out.o>
gcc -no-pie <out.o> <path/to/stdlib> -o <out>
```

