// Define a suitable type
class Pair
    var _x: U32 = 0 // `_<variable>` indicates a private variable
    var _y: U32 = 0

    new create(x: U32, y: U32) => // `new` indicates we are creating a new object, can have multiple constructors
        _x = x
        _y = y

    // Define a + function
    fun add(other: Pair): Pair =>
        Pair(_x + other._x, _y + other._y)

    fun factorial(x: I32): I32 ? =>
        if x < 0 then error end
        if x == 0 then
            1
        else
            x * factorial(x - 1)
        end

    // This means the `Pair` object can be used in a `with` statement
    fun dispose() =>
        // Do something...
        _y + _x


class Foo
    var _f: F64 = 0

    fun ref addmul(add: F64, mul: F64): F64 =>
        _f = (_f + add) * mul // This line would return the original value of `_f`
        _f // This line will make sure we return the new value of `_f`


actor Main
    new create(env: Env) =>
        var name: String = "Jonathan"
        env.out.print("Hello, world! I am " + name) // Can only output String or Array[U8]

        /*
        Primitive types as Unions

        type Number is (Signed | Unsigned | Float)
        type Signed is (I8 | I16 | I32 | I64 | I128)
        type Unsigned is (U8 | U16 | U32 | U64 | U128)
        type Float is (F32 | F64)

        */

        // IF statements
        var a: I32
        a = 10
        a = 11

        let b: I32 = 9 // Constant
        // b = 10 // Should fail

        if a > b then
            var x = "A is bigger"
            env.out.print(x)
        elseif b > a then
            env.out.print("B is bigger")
        else
            env.out.print("They are the same")
        end

        // Tuples
        var t: (String, U64)
        t = ("hi", 3)
        t = ("bye", 7)

        env.out.print(t._1)

        // Class business
        // Infix overloading - https://tutorial.ponylang.org/expressions/infix-ops.html
        var x = Pair(1, 2)
        var y = Pair(2, 3)
        var z = x + y // OR `x.add(z)` they are the same

        // Everything is an expression
        let lots = true
        var expr_if: I32 = 1 + if lots then 100 else 2 end // will equal 101 is lots == true, else 3

        var expr_x: (String | Bool) =
        if lots then
            "Hello"
        else
            false
        end

        // Loops
        for name1 in ["Bob"; "Fred"; "Sarah"].values() do // .values() returns an Iterator
            env.out.print(name1)
        end

        // Exceptions... Why do these even still exist..!
        try
            env.out.print("Exception testing")
            if true then error end
            env.out.print("This should never get called")
        else
            env.out.print("Exception_Else gets called when there is an `error` raised")
        then
            env.out.print("This will always be called regardless!")
        end

        var facPair = Pair(1, 2)
        try
          // This raises a compile time error when there is no try-statement around it, cool!
          // Although just returning an error would be much easier to follow control wise
          facPair.factorial(-1)
        end

        // With statement: Similar to Pythons, except the object needs a `.dispose()` method
        with obj = Pair(1, 2) do
            env.out.print("With Statement")
            error
        else // Will raise a compiler error is no error is raised in the with statement
            env.out.print("Will only run when an error occurs")
        end

        // Equality testing
        var p1 = Pair(1, 1)

        if p1 is Pair(1, 1) then
            env.out.print("These are equal") // THIS SHOULD NEVER PRINT
        else
            env.out.print("These Pair's are not equal")
        end

        let p2 = p1

        if p2 is p1 then
            env.out.print("These Pair's are equal")
        end

        // Function partials
        let foo: Foo = Foo
        let foo_f = foo~addmul(3) // Creates a function partial
        if foo_f(4) == 12 then
            env.out.print("Function partial works")
        end
