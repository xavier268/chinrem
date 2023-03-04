# chinrem


[![Go Reference](https://pkg.go.dev/badge/github.com/xavier268/chinrem.svg)](https://pkg.go.dev/github.com/xavier268/chinrem) 
[![Go Report Card](https://goreportcard.com/badge/github.com/xavier268/chinrem)](https://goreportcard.com/report/github.com/xavier268/chinrem)

## Chinese remainder arithmetic (massively parallele computation for very large numbers)

This librairy does arithmetic by internally encoding positive numbers using their remainers, modulo a set of primes.

The theorem of the chinese remainder states that : _a number is uniquely defined by its remainders modulo different primes, modulo the product of all the primes used._

The benefit of that approach is that :
* any positive big.Int can be encoded/decoded in a unique manner, modulo a very large "limit", that can be arbitrily set.
* most arithmetic operations can be performed separately and easily on each remainer,
* no memory allocation is required (as opposed to golang's big.Int)
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

Benchmark shows very significant gains (x50) for simple operations (multiply, inverse), but some disadvantage for exponentiation (x2)

    2023-03-04 13:31:48.816278621 +0100 CET m=+0.008978675
    goos: linux
    goarch: amd64
    pkg: github.com/xavier268/chinrem
    cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
    BenchmarkBvC/big.mul-8         	  357344	     57094 ns/op	   26296 B/op	       0 allocs/op
    BenchmarkBvC/chinrem.mul-8     	 1621789	       726.9 ns/op	       0 B/op	       0 allocs/op
    BenchmarkBvC/big.inv-8         	 2579451	       467.1 ns/op	     560 B/op	       8 allocs/op
    BenchmarkBvC/chinrem.inv-8     	100000000	        10.55 ns/op	       0 B/op	       0 allocs/op
    BenchmarkBvC/big.Exp-8         	  120085	      9820 ns/op	    1608 B/op	      12 allocs/op
    BenchmarkBvC/chinrem.Exp-8     	   51651	     23104 ns/op	     903 B/op	       1 allocs/op
    PASS
    ok  	github.com/xavier268/chinrem	27.868s