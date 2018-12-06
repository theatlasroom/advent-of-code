use std::collections::HashMap;
use utils;

// abcdef contains no letters that appear exactly two or three times.
// bababc contains two a and three b, so it counts for both.
// abbcde contains two b, but no letter appears exactly three times.
// abcccd contains three c, but no letter appears exactly two times.
// aabcdd contains two a and two d, but it only counts once.
// abcdee contains two e.
// ababab contains three a and three b, but it only counts once.

fn repeated_character(s: &str, max_repeats: i32) -> bool {
  let mut has_repeat = false;
  let mut set: HashMap<char, i32> = HashMap::new();
  for c in s.chars() {
    if !set.contains_key(&c) {
      set.insert(c, 1);
    } else {
      let mut count = match set.get(&c) {
        Some(&v) => v,
        _ => 0,
      };
      count += 1;
      if count >= max_repeats {
        has_repeat = true;
        break;
      } else {
        set.insert(c, count);
      }
    }
  }
  has_repeat
}

#[test]
fn will_check_for_repetitions_of_a_character() {
  let repeat_2 = vec![
    ("abcdef", false),
    ("bababc", true),
    ("abbcde", true),
    ("aabcdd", true),
  ];
  let repeat_3 = vec![
    ("abcdef", false),
    ("bababc", true),
    ("abbcde", false),
    ("aabcdd", false),
  ];

  for i in repeat_2 {
    let (q, a) = i;
    let res = repeated_character(q, 2);
    assert_eq!(res, a);
  }

  for i in repeat_3 {
    let (q, a) = i;
    let res = repeated_character(q, 3);
    assert_eq!(res, a);
  }
}

fn calculate_checksum(lines: &str) -> i32 {
  let mut repeat_2 = 0;
  let mut repeat_3 = 0;
  for line in lines.split("\n") {
    if repeated_character(&line, 2) {
      if repeated_character(&line, 3) {
        repeat_3 += 1
      } else {
        repeat_2 += 1
      }
    }
  }
  repeat_2 * repeat_3
}

pub fn solve() -> String {
  let data = utils::read_file("../data/2018_2.txt");
  let checksum = calculate_checksum(&data);
  format!("Box IDs checksum {:?}", checksum)
}
