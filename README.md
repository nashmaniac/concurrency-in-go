# Concurrency in Go

# Examples to practice

## Building Blocks
* Goroutine
    * Calling criteria
        - Function call
        - Inline function call
        - Anonymous function call
    * Memory space
        - Sharing same memory space
        - How to differentiate memory space of each goroutine
    * Initialize memory size
        - Find out the memory required for a goroutine to start
    * Context Switching
        - Benchmark for context switch

* WaitGroup
    - Waitgroup in a single function
    - Waitgroup passing as params to different function
* Mutex & RWMutex
    - Example of mutex
    - Benchmark of mutex and RWMutex
* Cond
    - basic cond example
    - Signal example
    - Broadcast example
* Once 
    - Example
* Pool
    - Basic example Pool
    - Advance pool example
    - Benchmarking of pool and non-pool implementation
* Channels
    - Buffered & Unbuffered channel
* Select statement
    - Composition of different channels

## Concurrency Patterns
* Confinement
    * Adhoc example
    * Lexical example

* For-select Statement
    * for-select example

* Goroutine Leaks
    * Basic example
    * From producer to consumer example
    * From consumer to producer example

* Or-Done channel
    * Or channel

* Error handling