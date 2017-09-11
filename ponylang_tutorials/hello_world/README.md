# Ponyland Hello World

Some example code testing out Ponylang features https://tutorial.ponylang.org.
Ponylang compiles down the binary code, it apparently has a concurrency model using actors.

https://tutorial.ponylang.org/capabilities/

A semi example on running a concurrent application with a producer and consumer that shares an internal buffer https://github.com/ponylang/ponyc/tree/master/examples/producer-consumer

And an example on running http get https://github.com/ponylang/ponyc/blob/master/examples/httpget/httpget.pony

## Running
```
$ ponyc # This will produce an executable
```

## Notes
```
- Runs on LLVM; kind of slow compile time for small examples
- No global values
- Everything is an expression
- Code can only be written in functions
- Mutable and Immutable variables via "var" and "let" respectively
- Immutable and Mutable functions via "fun" and "fun ref" respectively
    - An Immutable function can't update any of the objects fields
- Classes don't need a "self" or reference to the current instance in it's functions, it seems to automatically add the current instances variables to the functions local scope
- Private class methods using an underscore at the start of the field
- Has some weird infix operator functions that allow any class to define any infix operator: https://tutorial.ponylang.org/expressions/infix-ops.html
- Every infix operator is a function written on the object, so calling "+" calls the right side's .add(left side) method
- Annoyingly has exceptions, but there is only one type of exception `error` and a compiler error occurs if an error will be raised without being caught
- Has "with" statements similar to Python's
- Has function partials
- Has both "is" and "==" equality testing
- Does it have a string formatting function?
```
