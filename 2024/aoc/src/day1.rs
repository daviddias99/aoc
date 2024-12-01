use std::collections::HashMap;

pub fn solve_part1(input: &str) -> u32 {
    let mut left_list: Vec<u32> = vec![];
    let mut right_list: Vec<u32> = vec![];
    let mut total: u32 = 0;

    let lines = input.split('\n');

    // No need to use "first" and "last" since we know there will be
    // to elements in each Vec.

    for line in lines {
        let values: Vec<u32> = line
            .split_whitespace()
            .map(|num_str| num_str.parse().unwrap())
            .collect();
        left_list.push(values[0]);
        right_list.push(values[1]);
    }

    left_list.sort();
    right_list.sort();

    for i in 0..left_list.len() {
        if left_list[i] > right_list[i] {
            total += left_list[i] - right_list[i];
        } else {
            total += right_list[i] - left_list[i];
        }
    }

    return total;
}

pub fn solve_part2(input: &str) -> u32 {
    let mut left_list: Vec<u32> = vec![];
    let mut right_inventory: HashMap<u32, u32> = HashMap::new();
    let mut similarity_score: u32 = 0;

    let lines = input.split('\n');

    for line in lines {
        let values: Vec<u32> = line
            .split_whitespace()
            .map(|num_str| num_str.parse().unwrap())
            .collect();

        // No need to use "first" and "last" since we know there will be
        // to elements in each Vec.
        let left = values[0];
        let right = values[1];

        left_list.push(left);
        right_inventory.insert(right, right_inventory.get(&right).unwrap_or(&0) + 1);
    }

    for val in left_list {
        similarity_score += right_inventory.get(&val).unwrap_or(&0) * val;
    }

    return similarity_score;
}

#[cfg(test)]
mod tests {
    use crate::day1::*;

    #[test]
    fn d1_p1_ex1() {
        let input = std::fs::read_to_string("input/day1_example_1.txt").unwrap();
        assert_eq!(solve_part1(&input), 11);
    }

    #[test]
    fn d1_p1() {
        let input = std::fs::read_to_string("input/day1.txt").unwrap();
        assert_eq!(solve_part1(&input), 1666427);
    }

    #[test]
    fn d1_p2_ex1() {
        let input = std::fs::read_to_string("input/day1_example_1.txt").unwrap();
        assert_eq!(solve_part2(&input), 31);
    }

    #[test]
    fn d1_p2() {
        let input = std::fs::read_to_string("input/day1.txt").unwrap();
        assert_eq!(solve_part2(&input), 24316233);
    }
}
