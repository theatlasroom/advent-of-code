/* Santa is delivering presents to an infinite two-dimensional grid of houses.

   He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

   However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

   For example:

   > delivers presents to 2 houses: one at the starting location, and one to the east.
   ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
   ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses. */
use std::collections::HashMap;
use std::fmt;
use std::ops::Add;
use utils;

#[derive(Debug, Copy, Clone)]
struct Location {
  x: i32,
  y: i32,
}

impl Location {
  fn zero() -> Location {
    Location { x: 0, y: 0 }
  }

  fn move_location(self, direction: &Direction) -> Location {
    let displacement = match direction {
      Direction::North => Location { x: 0, y: 1 },
      Direction::South => Location { x: 0, y: -1 },
      Direction::East => Location { x: 1, y: 0 },
      Direction::West => Location { x: -1, y: 0 },
    };
    self + displacement
  }
}

// implement the `+` symbol the type `Location`
impl Add for Location {
  type Output = Location;

  fn add(self, displacement: Location) -> Location {
    Location {
      x: self.x + displacement.x,
      y: self.y + displacement.y,
    }
  }
}

// lets us use to_string
impl fmt::Display for Location {
  fn fmt(&self, fmt: &mut fmt::Formatter) -> fmt::Result {
    fmt.write_str(&format!("{:?}-{:?}", self.x, self.y))?;
    Ok(())
  }
}

#[derive(Debug, PartialEq)]
enum Direction {
  North,
  South,
  East,
  West,
}

fn parse_instruction(s: &str) -> Option<Direction> {
  match s {
    "^" => Some(Direction::North),
    "v" => Some(Direction::South),
    ">" => Some(Direction::East),
    "<" => Some(Direction::West),
    _ => None,
  }
}

#[test]
fn it_will_parse_any_valid_instruction() {
  let cases = [
    ("^", Direction::North),
    ("v", Direction::South),
    (">", Direction::East),
    ("<", Direction::West),
  ];

  for case in cases.iter() {
    let (i, result) = case;
    assert_eq!(&parse_instruction(i).unwrap(), result);
  }
}

fn calculate_deliveries_made(input: &str) -> usize {
  let mut deliveries = HashMap::new();
  let mut loc = Location::zero();
  // add the first location
  deliveries.insert(loc.to_string(), loc);
  for c in input.chars() {
    let dir = parse_instruction(&c.to_string()).unwrap();
    loc = loc.move_location(&dir);
    deliveries.insert(loc.to_string(), loc);
  }
  deliveries.len()
}

#[test]
fn it_will_calculate_our_test_cases() {
  let cases = [(">", 2), ("^>v<", 4), ("^v^v^v^v^v", 2)];
  for case in cases.iter() {
    let (problem, solution) = case;
    assert_eq!(&calculate_deliveries_made(problem), solution);
  }
}

pub fn solve() -> String {
  let data = utils::read_file("../data/2015_3.txt");
  let deliveries = calculate_deliveries_made(&data);
  format!("Santa delivers {:?} presents", deliveries)
}
