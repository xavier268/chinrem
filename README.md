# chinrem

## Chinese remainder arithmetic (massively parallele computation for very large numbers)

This librairy allows to perfom arithmetic by internally encoding positive numbers using their remainers, modulo primes.

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

    2023-02-27 22:45:44.8145345 +0100 CET m=+0.004601701
    === RUN   BenchmarkBigVersusChinrem/big.Int-mul
    BenchmarkBigVersusChinrem/big.Int-mul
    BenchmarkBigVersusChinrem/big.Int-mul-16
    1000000            103266 ns/op           73156 B/op          0 allocs/op
    === RUN   BenchmarkBigVersusChinrem/chinrem.CRI-mul
    BenchmarkBigVersusChinrem/chinrem.CRI-mul
    BenchmarkBigVersusChinrem/chinrem.CRI-mul-16
    1679671               714.1 ns/op             0 B/op          0 allocs/op
    === RUN   BenchmarkBigVersusChinrem/big.Int-inv
    BenchmarkBigVersusChinrem/big.Int-inv
    BenchmarkBigVersusChinrem/big.Int-inv-16
    2051143               517.8 ns/op           560 B/op          8 allocs/op
    === RUN   BenchmarkBigVersusChinrem/chinrem.CRI-inv
    BenchmarkBigVersusChinrem/chinrem.CRI-inv
    BenchmarkBigVersusChinrem/chinrem.CRI-inv-16
    100000000               10.58 ns/op            0 B/op          0 allocs/op
    PASS