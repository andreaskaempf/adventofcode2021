# Advent of Code 2021

This repository contains my entries for the Advent of Code 2021.

* **Day 1** (C, Python): Count number of increases in a text file of lines (easy)
* **Day 2** (Go, Python): Process the state of a submarine by interpreting 
    instructions "forward", "up", "down", etc. (easy)
* **Day 3** (Julia, Python): Read list of binary numbers to establish 
    codes based on column-wise bit frequency (easy)
* **Day 4** (Go): Simulate game of bingo, by applying list of called numbers
    until first board wins (for part 2, also find the last board
    that wins -- easy)
* **Day 5** (Go): Trace lines on a grid, given starting and ending 
    coordinates of each line, and count up where 2 or more lines cross 
    (part 1 only includes horizontal and vertical lines, part 2 diagonal 
    as well -- easy).
* **Day 6** (Python, Go): Simulate growth of a population of "lanternfish",
    tricky because performance crawls if you try to simulate individual
    creatures (used dictionary/array of fish by age for speed -- medium).
* **Day 7** (Julia): Find target position that minimizes fuel use for 
    a population of crabs that need to move there, under different 
    assumptions about fuel use per each crab's distance moved 
    (optimization problem, but used brute force -- easy).
* **Day 8** (Julia): Decipher scrambled digits on a digital display (very hard,
    failed using logic so solved using brute force)
* **Day 9** (Julia): Find lowest points in a 2-dimensional space, and the 
    areas around these low points, bounded by edges or 9s (medium).
* **Day 10** (Python): Identify and resolve different types of mis-matched
    brackets in a file of strings (easy).
* **Day 11** (Julia): Simulate energy growth of a field of "octopusses", with each
    exploding when they reach a level, infecting those around them; for part 2,
    identify when they all explode in the same iteration (easy).
* **Day 12** (Julia): Find the number of paths from start to end in a graph, allowing repeat
    visits for "large caves" (capitalized), but either no (part 1) or at
    most 2 (part 2) visits for "large caves" (lower case) (hard).
* **Day 13** (Julia): Repeatedly "fold" a matrix of 1/0 values along 
    horizontal or vertical lines, and (part 1) count up the number of 
    points remaining after one fold, and (part 2) visualize the matrix
    after all folds to yield a dot matrix alphabetic code.
* **Day 14** (Python): Apply character insertion rules to a sequence of characters,
    by managing list of possible transitions rather than building up string
    of characters (otherwise runs out of memory in Part 2).

To compile and run a Go program
* Change into the directory with the program
* go build day02.go
* ./day02

To run a Julia program
* Change into the directory with the program
* julia day02.jl

To compile and run a C program
* Change into the directory with the program (e.g., cd day01)
* gcc day01.c -o day01
* ./day01

And of course, to run a Python program
* Change into the directory with the program
* python day06.py

AK, Dec 2021
