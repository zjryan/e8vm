# todo

- build a github hook that visualizes the code
- asm importing

# library layer

a library has three layers.

the bottom layer is the linking layer, it has a symbol table of variables
and functions (or small data segments and small code segments), data segments
has alignments

the upper layer is the header layer, it has the table of available functions,
variables, types and constants.

the last layer has the actual data bytes for the data segments and code segments.
the code segments also has linking places.

the linker only cares about the bottom layer and last layer

imported package only looks at the second layer