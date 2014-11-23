The EnableInt/DisableInt member functions of the Interrupt struct might be
redundant.  Programs running on CPU will just directly manipulate the memory.
And programs outside CPU do not need to modify that.

But let's just leave it there for a while. If they have no use, we will remove
them later.


