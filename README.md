# chinrem

## Chinese remainder arithmetic (massively parallele computation for very large numbers)

This librairy allows to perfom arithmetic by rinternally encoding positive numbers using their remainers, modulo the sucessive primes.

The benefit of that approach is that :
* any positive big.Int can be encoded/decoded in a unique manner, modulo a very large "limit", that can be arbitrily set.
* most arithmetic operations can be performed separately and easily on each remainer
* (future step) massive thread parralelization becomes possible, to maximize use of multi core cpu ...


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