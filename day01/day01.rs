// Advent of Code 2021, Day 1 (Rust version)
//
// Read a text file of numbers, report the number of increases,
// first between subsequent lines, then between subsequent sums
// of 3 lines.
//
// To compile this Rust program: rustc day01.rs
// To run: ./day01 (or use cargo to build & run)
// 
// AK, 3/01/2022

use std::fs;

fn main() {

    // Read file as one string
    let fname = "input.txt".to_string(); 
    println!("Reading {}", fname);
    let data = fs::read_to_string(&fname)
        .expect("Error reading file");

    // Split by newlines
    let lines: Vec<&str> = data.split("\n").collect();
    println!("{} lines", lines.len());

    // Convert to list of numbers
    let mut nums: Vec<i32> = Vec::new();
    for l in lines {
        let n = parse_int(l);
        if n  >= 0 {  // i.e., no parse error
            nums.push(n);
        }
    }
    println!("{} numbers", nums.len());

    // Part 1: Count the number of increases
    let mut prev_n = -1;    // to avoid counting the first number as an increase
    let mut increases = 0;
    for n in &nums {                     // need to use & on vector of ints!!!
        if prev_n >= 0 && n > &prev_n {  // need & for second use of prev_n!!! wtf???
            increases += 1;
        }
        prev_n = *n;        //  need to use pointer here, bizarre!
    }

    println!("Part 1: there were {} increases (should be 1832)", increases);

    // Part 2: Count the number of increases between consecutive 3-value windows
    let mut prev_w = -1;  // to avoid counting the first number as an increase
    let mut w_increases = 0;
    let mut i = 2;  // end of first window
    while i < nums.len() {
        let w = nums[i] + nums[i-1] + nums[i-2]; 
        if prev_w >= 0 && w > prev_w {
            w_increases += 1;
        }
        prev_w = w;
        i += 1;
    }

    println!("Part 2: there were {} window increases (should be 1858)", w_increases);
}

// Parse an integer string, return -1 if there was an error
fn parse_int(s: &str) -> i32 {
    match s.parse::<i32>() {
        Ok(n) => return n,
        Err(_) => return -1,
    }
}
