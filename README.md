# chinrem

## Chinese remainder arithmetic (massively parallele computation for very large numbers)

This librairy allows to perfom arithmetic by internally encoding positive numbers using their remainers, modulo primes.

The theorem of the chinese remainder states that : _a number is uniquely defined, modulo the product of the primes used, by its remainder modulo different primes._

The benefit of that approach is that :
* any positive big.Int can be encoded/decoded in a unique manner, modulo a very large "limit", that can be arbitrily set.
* most arithmetic operations can be performed separately and easily on each remainer,
* limited to no memory allocation is required,
* _(future step) massive thread paralelization becomes possible, to maximize use of multi core cpu ..._


## How to use 

    // First, you need to create an engine that will be use to compute,
    // specifying the number of primes used for the computation.
    // CREngine are immutable, and safe for concurrent use.
    e:= NewCREngine(20) // create an engine with 20 primes.

    // Then, creates the number (called CRI), as needed.
    a:= NewCRIInt64(2577) // there are many ways to create a CRI. 

    // Do the maths ...
    a.Mul(a,a)
    a.Inv(a)

    // Print the result 
    fmt.Println(a)

## Benchmarks

Using big.Int package (from the go standard library) versus this package (chinrem).

    2023-02-27 17:52:38.3041445 +0100 CET m=+0.004099401
    goos: windows
    goarch: amd64
    pkg: github.com/xavier268/chinrem
    cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
    BenchmarkBigVersusChinrem/big.Int-mul-16         	 1000000	    102797 ns/op	   73156 B/op	       0 allocs/op
    BenchmarkBigVersusChinrem/chinrem.CRI-mul-16     	 1680429	       714.6 ns/op	       0 B/op	       0 allocs/op
    BenchmarkBigVersusChinrem/big.Int-inv-16         	 3178966	       435.5 ns/op	     560 B/op	       8 allocs/op
    BenchmarkBigVersusChinrem/chinrem.CRI-inv-16     	100000000	        10.53 ns/op	       0 B/op	       0 allocs/op
    PASS