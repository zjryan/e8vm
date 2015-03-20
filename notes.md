# todo

- new import semantics, clean up import structure

- let lexer/parser also save a token array
- formatter: rearrange token position
- printer: print the file back with the rearranged token position

- build/make system, track timestamp

## importing

- only one file can have the import header
- this file will be automatically imported in all the other files in the directory, so it must not rely on any other file in the directory to make
sure there is no circular dependency