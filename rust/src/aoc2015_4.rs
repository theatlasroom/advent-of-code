#![allow(dead_code)]
// Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts for all the economically forward-thinking little girls and boys.

// To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least five zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.

// For example:

// If your secret key is abcdef, the answer is 609043, because the MD5 hash of abcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest such number to do so.
// If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....
// Your puzzle input is ckczppom.
use md5;
use regex::Regex;
use utils;

enum AdventCodePattern {
  One,
  Two,
}

fn is_valid_adventcoin_hash(pattern: &AdventCodePattern, s: &str) -> bool {
  let re = match pattern {
    AdventCodePattern::One => Regex::new(r"^00000[A-Fa-f0-9]{27}$").unwrap(),
    AdventCodePattern::Two => Regex::new(r"^000000[A-Fa-f0-9]{26}$").unwrap(),
  };
  re.is_match(s)
}

#[test]
fn it_will_validate_correct_hashes() {
  let cases_one = [
    ("48fbdf1af6eb206e65ef98bf8a78ad85", false),
    ("ab1cf84209ffe088ac7822af3eb8b533", false),
    ("a00001dbbfa3a5c83a2d506429c7b00e", false),
    ("", false),
    ("00000XXXXXXXXXX", false), // should be 32 digits long
    ("000001dbbfa3a5c83a2d506429c7b00e", true),
    ("000006136ef2ff3b291c85725f17325c", true),
  ];

  for elem in cases_one.iter() {
    let (input, result) = elem;
    let computed = &is_valid_adventcoin_hash(&AdventCodePattern::One, &input);
    assert_eq!(computed, result);
  }

  let cases_two = [
    ("48fbdf1af6eb206e65ef98bf8a78ad85", false),
    ("ab1cf84209ffe088ac7822af3eb8b533", false),
    ("a00001dbbfa3a5c83a2d506429c7b00e", false),
    ("", false),
    ("00000XXXXXXXXXX", false), // should be 32 digits long
    ("000001dbbfa3a5c83a2d506429c7b00e", false),
    ("0000006136ef2ff3b291c85725f1732c", true),
  ];

  for elem in cases_two.iter() {
    let (input, result) = elem;
    let computed = &is_valid_adventcoin_hash(&AdventCodePattern::Two, &input);
    assert_eq!(computed, result);
  }
}

fn calculate_first_hash(pattern: &AdventCodePattern, input: &str) -> i32 {
  // println!("calculate_first_hash for {:?}", input);
  let mut seed = 0;
  loop {
    let digest = md5::compute([String::from(input), seed.to_string()].join(""));
    let hash = format!("{:x}", digest);
    // println!("seed: {:?} digest: {:?}", seed, digest);
    if is_valid_adventcoin_hash(&pattern, &hash) {
      break;
    }
    seed += 1;
  }
  seed
}

// TODO: reimplement using parallel processing, or maybe some kind of divide / conquer algorithms?
// #[test]
// fn it_will_compute_our_test_cases() {
//   let cases = [("abcdef", 609043), ("pqrstuv", 1048970)];

//   for case in cases.iter() {
//     let (input, result) = case;
//     let computed = &calculate_first_hash(&input);
//     assert_eq!(computed, result);
//   }
// }

fn part_one(data: &str) -> i32 {
  calculate_first_hash(&AdventCodePattern::One, &data.trim())
}

fn part_two(data: &str) -> i32 {
  calculate_first_hash(&AdventCodePattern::Two, &data.trim())
}

pub fn solve() -> String {
  let data = utils::read_file("../data/2015_4.txt");
  let result_one = part_one(&data);
  let result_two = part_two(&data);
  format!(
    "{:?} produces our valid AdventCoin for input {:?} with 5 leading 0's
     \n{:?} produces our valid AdventCoin for input {:?} with 6 leading 0's",
    result_one,
    &data.trim(),
    result_two,
    &data.trim(),
  )
}
