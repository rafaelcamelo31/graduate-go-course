# Go Internals

## Multitask and its Timeline

- Before Multitask - (1940 - 1960)
- Time-Sharing System - (1960 - 1970)
- OS's Multitask - (1980)
- Hyper-Threading - (1990 - 2000)
- Multi-Core - (2000)
- Optimization for Cloud, AI and etc (2010)

## Processes

- An instance of a program in execution
- Components
  - Addressing (part in memory dedicated for a process)
  - Contexts
    - Datas and informations OS saves to manage processes
      - Ex.:Program Counter(PC) or Instruction Pointer(IP)
      - Have the address for next instruction processor will execute
      - Assist in Context Switching

## Process Register

- Area where datas and addresses are saved temporary in CPU to be executed
- Datas
  - Ex.: Arithmetic operations and logics
- Register Address
  - Memory allocation including stack pointers
  - Ex.: When accessing a variable, CPU has a registry of a memory to save its value

## Heap and Stack

[Heap and Stack as memory](https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d)

- Heap

  - Dynamic memory allocation at runtime

- Stack

  - Used for storing local variables, function arguments, and return addresses during function calls. (LIFO)

## Status and Flags in Registers

- Gives recent status from operations executed by CPU
- Uses specific bits (flags)
- Ex.:
  - Flag Zero (Z): Result of the operation which returned 0.
  - Flag Signal (S) or Negative(N): Indicates whether the result of operation is positive or negative
  - Flag Overflow: Overcapacity in produced result

## Lifecycle of a process

- Creation

  - UNIX/Linux uses fork() to create new process when a program requests for process execution

- Execution

  - The process is being executed by CPU. Can be "executing" or "ready"

- Waiting/Blocked
  - The process is suspended until external event finishes. Common in I/O, where process waits for disc reading or waits for networkd input
- Termination
  - The process completes it execution or terminated forcefuly
  - Exit: Success after concluding instructions
  - Killed: Interruption by error execution or terminated by other process (such as command "kill")

## Creating Process

- UNIX/Linux

  - fork()
  - Clones current process
  - Creates a child process
  - fork() returns diffent value to the parent process (PID)
  - Parent and child processes are identical. However, the values in memory are copied to the different address, separate and independent

<img src="image/pid.png" height="300">

## Managing Process

- Scheduler
  - Decides which processe will be executed
  - Alternate between processes
  - Has several algorithms in attempt to maximize CPU use
  - It selects processes from queue with "ready queue"
  - Allocates CPU: Switch states, Ready to Running
  - Frees CPU: I/O when not needed
- Types of Schedulers
  - Non-preemptive: The process in execution has control on when to free CPU for other process. It has a drawback of one process using CPU exclusively.
  - Preemptive: The OS is capable to interrupt a process in execution and give CPU different process. It operates "fairly". Lots of "Context Switching".

## Threads

[Program vs Process vs Thread](https://bytebytego.com/guides/what-is-the-difference-between-process-and-thread/)

- **Program** is an executable file containing a set of instructions. One program can have multiple processes.
- **Process** is an instance of program in execution. When a program is loaded into the memory and becomes active, the program becomes a process.
- **Thread** is the smallest unit of execution within a process.

## Memory in Threads

- Thread uses less memory than process.
- Each thread has own stack independently and isolated.
- Each thread (OS) uses 2MB. (Golang thread uses 2kb)

## Runtime Architecture

- Scheduler
- Goroutines
- Channels
- Garbage Collector
- Memory Allocation
- Stack Management
- Network Poller
- Reflection

## M:N Threading Model

[Introducing M:N Hybrid Threading in Go: Unveiling the Power of Goroutines](https://medium.com/@rezauditore/introducing-m-n-hybrid-threading-in-go-unveiling-the-power-of-goroutines-8f2bd31abc84)

- User-level vs kernel-level threads
- Allows for a dynamic allocation of M user-level threads (goroutines) to N kernel-level threads.
- Allows flexibility and efficiency to concurrent programming.

## Goroutines

- Functions or methods executed concurrently.
- Threads managed by Go runtime.
- Much cheaper than kernel threads (2kb).
- Easier to create and destroy.
- Has own stack. Shares same address of memory in Go program.

## M:P:G Model

[Revealing Golang’s Secret Sauce: A Deep Dive into Its Internals](https://meetsoni15.medium.com/unveiling-golangs-hidden-internals-discover-the-hidden-mechanics-that-optimize-performance-8f946f784041)

<img src="image/mpg.png" height="300">

## runtime.GOMAXPROCS()

[GOMAXPROCS](https://go.dev/blog/container-aware-gomaxprocs)

- Go creates a P (Processor) per CPU cores.
- Go tries to create a M (Machine -Threads) for each P.

## Scheduler

[Scheduler Structures](https://go.dev/src/runtime/HACKING)

> The scheduler’s job is to match up a G (the code to execute), an M (where to execute it), and a P (the rights and resources to execute it). When an M stops executing user Go code, for example by entering a system call, it returns its P to the idle P pool. In order to resume executing user Go code, for example on return from a system call, it must acquire a P from the idle pool.

- The scheduler in Go runtime:

  - Manages task
  - Load balancing
  - Manages concurrencies

- Go version >= 1.14 works in preemptive scheduler
  - Which means Go is capable to allocate and deallocate processes depending on the situation

<img src="image/scheduler.png" height="420">

## Memory Management

- Fast access

  - L1 - 64kb
  - L2 - 0.5mb
  - L3 - 8mb
  - Used in CPU as a cache with same chip

- Slow access
  - RAM sticks, DDR - Double Data Rate. Clock can access twice per cicle
  - Connected with memory bus (communication channel between CPU and memory)
  - Memories are referenced in Hexadecimal format

<img src="image/memory_management.png" height="300">

## Memory Access Cost

### Memory management

<img src="image/metadata.png" height="300">

### Stack

<img src="image/stack.png" height="300">

### Heap

<img src="image/heap.png" height="300">

## Memory Fragmentation

Strategy to avoid ineffienct memory management

- Memory Arenas

<img src="image/memory_arenas.png" height="300">

- Memory Allocations

  - malloc(C std library)
  - dlmalloc(Doug Lea's Malloc) - Doesn't support multithreading efficiently
  - ptmalloc/ptmalloc2 (pthreads Malloc) - First memory arenas appearence
  - jemalloc (Jason Evans) - Facebook, Rust, Postgres
  - TCMalloc (Thread-Caching Malloc) - Google

## Memory Allocation in Go

[A visual guide to Go memory allocator](https://medium.com/@ankur_anand/a-visual-guide-to-golang-memory-allocator-from-ground-up-e132258453ed)

- mallocgc

  <img src="image/span.png" height="150">
  <img src="image/mcentral_mcache.png" height="300">
