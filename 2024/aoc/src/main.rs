use std::env;

mod day1;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() != 2 {
        eprintln!("Usage: cargo run <day_number>");
        std::process::exit(1);
    }

    let day_number: usize = args[1].parse().expect("Invalid day number");

    match day_number {
        1 => {
            let input = std::fs::read_to_string("input/day1.txt").unwrap();
            let result1 = day1::solve_part1(&input);
            let result2 = day1::solve_part2(&input);
            println!("Day {}, Part 1: {}", day_number, result1);
            println!("Day {}, Part 2: {}", day_number, result2);
        }
        _ => {
            eprintln!("Invalid day number");
            std::process::exit(1);
        }
    }
}
