use std::collections::HashMap;
use utils;

// abcdef contains no letters that appear exactly two or three times.
// bababc contains two a and three b, so it counts for both.
// abbcde contains two b, but no letter appears exactly three times.
// abcccd contains three c, but no letter appears exactly two times.
// aabcdd contains two a and two d, but it only counts once.
// abcdee contains two e.
// ababab contains three a and three b, but it only counts once.

// Todo: should return a tuple (repeat_2, repeat_3)
fn repeated_character(s: &str, c: &str, max_repeats: i32) -> bool {
  let mut has_repeat = false;
  let mut count = 0;
  for curr in s.chars() {
    if curr.to_string() == c {
      count += 1;
    }
    if count == max_repeats {
      has_repeat = true;
    } else {
      has_repeat = false;
    }
  }
  has_repeat
}

#[test]
fn will_check_for_repetitions_of_a_character() {
  let repeat_2 = vec![
    ("abcdef", "a", false),
    ("bababc", "b", true),
    ("abbcde", "b", true),
    ("aabcdd", "d", true),
  ];
  let repeat_3 = vec![
    ("abcdef", "a", false),
    ("bababc", "b", true),
    ("abbcde", "b", false),
    ("aabcdd", "d", false),
  ];

  for i in repeat_2 {
    let (q, c, a) = i;
    let res = repeated_character(q, c, 2);
    assert_eq!(res, a);
  }

  for i in repeat_3 {
    let (q, c, a) = i;
    let res = repeated_character(q, c, 3);
    assert_eq!(res, a);
  }
}

fn calculate_checksum(lines: &str) -> i32 {
  let mut repeat_2 = 0;
  let mut repeat_3 = 0;
  for line in lines.split("\n") {
    let mut has_2 = false;
    let mut has_3 = false;
    for c in line.trim().chars() {
      if repeated_character(&line, &c.to_string(), 2) {
        has_2 = true;
      }
      if repeated_character(&line, &c.to_string(), 3) {
        has_3 = true;
      }
    }
    if has_2 {
      repeat_2 += 1;
    }
    if has_3 {
      repeat_3 += 1;
    }
  }
  println!("rep2: {:?} rep3: {:?}", repeat_2, repeat_3);
  repeat_2 * repeat_3
}

pub fn solve() -> String {
  let data = utils::read_file("../data/2018_2.txt");
  let checksum = calculate_checksum(&data);
  format!("Box IDs checksum {:?}", checksum)
}
