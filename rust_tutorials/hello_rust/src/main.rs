fn variable_bindings() {

    // "!" means we are calling a Macro
    println!("Hello, world!");

    // Let pattern(?)
    let (x, y) = (5, 6);

    // Type Annotations
    let z: i32 = 7;
    // Below does not work as all bindings are immutable
    // z += 1;

    // Make variable binding mutable
    let mut a = 5;
    a += 1;

    println!("z={}, a={}", z, a);

    println!("{} + {} = {}", x, y, x + y);

    {
        let b = 100;
        println!("{}", b);
    }

    // Below wont' work due to Scope
    //println!("{}", b)

    // Variable shadowing
    let x: i32 = 8;
    {
        println!("{}", x); // Prints "8".
        let x = 12;
        println!("{}", x); // Prints "12".
    }
    println!("{}", x); // Prints "8".
    let x =  42;
    println!("{}", x); // Prints "42".
}

// Rust functions return exactly one value
// The last line of a function determines what it returns
fn add(x: i32, y: i32) -> i32 {
     x + y
}

// The "-> !" implies that this function never returns
fn diverges() -> ! {
    panic!("This function never returns!");
}

fn func_pointers() {
    // Function pointers
    fn plus_one(i: i32) -> i32 {
        i + 1
    }

    // Without type inference:
    let f: fn(i32) -> i32 = plus_one;

    // With type inference:
    let g = plus_one;

    println!("f(2) = {}", f(2));
    println!("g(2) = {}", g(2));
}

fn primitive_types() {
    // Booleans
    let x = true;
    let y: bool = false;
    println!("x={} y={} (x==y)={}", x, y, x==y);

    // Chars
    let x = 'P';
    let y = '£';
    println!("{}{}", y, x);

    // Numeric Types
    let x = 42; // `x` has type `i32`.
    let y = 1.0; // `y` has type `f64`.
    println!("x={} y={}", x, y);
    // List of numeric types
    /*
    i8
    i16
    i32
    i64
    u8
    u16
    u32
    u64
    isize
    usize
    f32
    f64
    */

    // Arrays
    let a = [1, 2, 3];
    println!("a has {} elements", a.len());
    println!("a[1] == {}", a[1]);

    // Slices
    let a = [0, 1, 2, 3, 4];
    let complete = &a[..]; // A slice containing all of the elements in `a`.
    let middle = &a[1..4]; // A slice of `a`: only the elements `1`, `2`, and `3`.
    println!("complete.len={}", complete.len());
    println!("middle.len={}", middle.len());

    // Tuples
    let x = (1, "hello");
    let x: (i32, &str) = (1, "hello");

    let mut x = (1, 2); // x: (i32, i32)
    let y = (2, 3); // y: (i32, i32)
    // You can assign one tuple into another, if they have the same contained types and arity. Tuples have the same arity when they have the same length.
    x = y;
    let (z, _) = x;
    println!("z={}", z);
    // Indexing
    let tuple = (1, 2, 3);

    let x = tuple.0;
    let y = tuple.1;
    let z = tuple.2;

    println!("x is {}", x);
}

fn expressions() {
    let x = 5;
    if x == 5 {
        println!("It is 5");
    } else if x == 6 {
        println!("It is 6");
    } else {
        println!("Something else!");
    }

    // If statements are expressions...
    // An if without an else always results in () as the value.
    let y = if x == 5 { 10 } else { 15 };
    println!("{}", y);

    // Loops
    // `loop` loops forever
    // loop { println!("x"); }

	'outer: for x in 0..10 {
		'inner: for y in 0..10 {
			if x % 2 == 0 { continue 'outer; } // Continues the loop over `x`.
			if y % 2 == 0 { continue 'inner; } // Continues the loop over `y`.
			println!("x: {}, y: {}", x, y);
		}
	}

}

fn vectors() {
	let v = vec![1, 2, 3, 4, 5];

	// Type inferance works
	// `x` needs to be a usize
	let x = 0;
	println!("v[0] is {}", v[x]);
}

fn ownership() {

	fn skip_prefix(line: &str, prefix: &str) -> &str {
		let n = "hello";
		return &n;
	}

	let line = "lang:en=Hello World!";
	let lang = "en";

	let v;
	{
		let p = format!("lang:{}=", lang);  // -+ `p` comes into scope.
		v = skip_prefix(line, p.as_str());  //  |
	}                                       // -+ `p` goes out of scope.
	println!("{}", v);
}

fn main() {

    //variable_bindings();
    //println!("{}", add(4, 5));
    //func_pointers();
    //primitive_types();
	//expressions();
	//vectors()
	ownership();
}
