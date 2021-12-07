# Advent of Code 2021

This repository contains my entries for the Advent of Code 2021.

* **Day 1** (C, Python): Count number of increases in a text file of lines
* **Day 2** (Go, Python): Process the state of a submarine by interpreting 
    instructions "forward", "up", "down", etc.
* **Day 3** (Julia, Python): Read list of binary numbers to establish 
    codes based on column-wise bit frequency
* **Day 4** (Go): Simulate game of bingo, by applying list of called numbers
    until first board wins (for part 2, also find the last board
    that wins)
* **Day 5** (Go): Trace lines on a grid, given starting and ending 
    coordinates of each line, and count up where 2 or more lines cross 
    (part 1 only includes horizontal and vertical lines, part 2 diagonal 
    as well).
* **Day 6** (Python, Go): Simulate growth of a population of "lanternfish",
    tricky because performance crawls if you try to simulate individual
    creatures (used dictionary/array of fish by age for speed).
* **Day 7** (Julia): Find target position that minimizes fuel use for 
    a population of crabs that need to move there, under different 
    assumptions about fuel use per each crab's distance moved 
    (optimization problem, but used brute force).

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
