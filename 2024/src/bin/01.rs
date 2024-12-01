advent_of_code::solution!(1);

pub fn part_one(input: &str) -> Option<i32> {
    let lines = input.lines().into_iter();
    let mut list_one = lines
        .clone()
        .map(|line| line.split_whitespace().nth(0).unwrap())
        .map(|line| line.parse::<i32>().unwrap())
        .collect::<Vec<i32>>();
    let mut list_two = lines
        .map(|line| line.split_whitespace().nth(1).unwrap())
        .map(|line| line.parse::<i32>().unwrap())
        .collect::<Vec<i32>>();

    list_one.sort();
    list_two.sort();

    let mut total_difference = 0;
    for (i, j) in list_one.iter().zip(list_two.iter()) {
        total_difference += i32::abs(i - j);
    }

    return Some(total_difference);
}

pub fn part_two(input: &str) -> Option<i32> {
    let lines = input.lines().into_iter();
    let mut list_one = lines
        .clone()
        .map(|line| line.split_whitespace().nth(0).unwrap())
        .map(|line| line.parse::<i32>().unwrap())
        .collect::<Vec<i32>>();
    let mut list_two = lines
        .map(|line| line.split_whitespace().nth(1).unwrap())
        .map(|line| line.parse::<i32>().unwrap())
        .collect::<Vec<i32>>();

    list_one.sort();
    list_two.sort();

    let mut total: i32 = 0;
    let mut latest: i32 = 0;
    let mut latest_apps: i32 = 0;

    for i in list_one.iter() {
        if *i == latest {
            total += latest_apps;
        } else {
            let apps = list_two.iter().filter(|&n| *n == *i).count();
            latest = *i;
            latest_apps = apps as i32 * *i;
            total += latest_apps;
        }
    }

    return Some(total);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, None);
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, None);
    }
}
