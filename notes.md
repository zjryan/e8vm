# todo

- add web hook
- new import semantics, clean up import structure
- printer: print the file back with their token position
- formatter: rearrange token position
- build/make system, track timestamp

## importing

- only one file can have the import header
- this file will be automatically imported in all the other files in the directory, so it must not rely on any other file in the directory to make
sure there is no circular dependency

## some things needs to fix

- files might need more information embeded, to avoid misplaced filenames
- recursive building might not be the best way to do it. should try
  build the real buidling graph
