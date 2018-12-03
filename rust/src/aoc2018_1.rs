use regex::Regex;
use utils;

#[derive(Debug)]
enum FrequencyChange {
  Plus,
  Minus,
  Noop,
}

#[derive(Debug)]
struct Instruction {
  operator: FrequencyChange,
  amount: i32,
}

#[test]
fn will_extract_a_set_of_instructions_from_a_string() {
  let a1 = vec![Instruction {
    operator: FrequencyChange::Plus,
    amount: 1,
  }];
  let a2 = vec![
    Instruction {
      operator: FrequencyChange::Plus,
      amount: 1,
    },
    Instruction {
      operator: FrequencyChange::Minus,
      amount: 2,
    },
    Instruction {
      operator: FrequencyChange::Noop,
      amount: 0,
    },
  ];
  let strs: Vec<(&str, Vec<Instruction>)> = vec![("+1", a1), ("+1\n-2\nasdf", a2)];
  for t in strs.iter() {
    let (q, a) = t;
    assert_eq!(a, &extract_instructions(q));
  }
}

fn extract_instructions(lines: &str) -> Vec<Instruction> {
  lazy_static! {
    static ref re: Regex = Regex::new(r"(?P<operator>(\+|-))(?P<amount>(\d+))").unwrap();
  }

  let mut v: Vec<Instruction> = Vec::new();

  for line in lines.trim().split("\n") {
    let i = match re.captures(line) {
      Some(caps) => {
        let operator = caps
          .name("operator")
          .map_or(FrequencyChange::Noop, |m| match m.as_str() {
            "+" => FrequencyChange::Plus,
            "-" => FrequencyChange::Minus,
            _ => FrequencyChange::Noop,
          });

        let amount = caps
          .name("amount")
          .map_or(0, |m| m.as_str().parse::<i32>().unwrap());

        Instruction { operator, amount }
      }
      None => Instruction {
        operator: FrequencyChange::Noop,
        amount: 0,
      },
    };
    v.push(i);
  }
  v
}

fn calculate_resulting_frequency(instructions: &Vec<Instruction>, initial: i32) -> i32 {
  instructions
    .iter()
    .fold(initial, |acc, i| match i.operator {
      FrequencyChange::Plus => acc + i.amount,
      FrequencyChange::Minus => acc - i.amount,
      _ => acc,
    })
}

#[test]
fn will_correctly_sum_instructions() {
  let i = vec![
    Instruction {
      operator: FrequencyChange::Plus,
      amount: 1,
    },
    Instruction {
      operator: FrequencyChange::Minus,
      amount: 3,
    },
    Instruction {
      operator: FrequencyChange::Noop,
      amount: 0,
    },
  ];
  assert_eq!(calculate_resulting_frequency(&i, 0), -2);
  assert_eq!(calculate_resulting_frequency(&i, 5), 3);
}

pub fn solve() -> String {
  let data = utils::read_file("../data/2018_1.txt");
  let instructions = extract_instructions(&data);
  let resulting_frequency = calculate_resulting_frequency(&instructions, 0);
  format!("{:?} Final frequency", resulting_frequency)
}
