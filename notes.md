# design 

- just back with github
- build a github hook, that is why we need an ec2 instance. we want to build a 
  hook responder

Here are the steps to (partially) build a program or a library.

1. get the list of packages that require to build; this is the input
2. load the packages, parse the package imports, and the build timestamps
3. if new packages are imported, also load them
4. figure out the packages that need to rebuild based on timestamps
5. load the headers/symbols of the packages that do not need to rebuild but need to load
6. build the packages based on its dependency order
7. cache the headers/sybmols built, save the built object archive and headers back to persistant storage

- the assembly compiler will just take a bunch of file openers,
  that will return the readcloser upon calling.
  it will also take an importer.
  an importer gives it back a readcloser (or error) upon calling for a package name
  the assembler will save the 