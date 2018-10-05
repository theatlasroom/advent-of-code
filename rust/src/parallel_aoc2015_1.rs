/*
* TODO: Reimplement in parallel
* - split the input string
* - run pattern match code to calculate final floor
* - combine results (final result = LHS + RHS)
 */
use utils;

pub const INITIAL_FLOOR: i32 = 0;

struct TestData {
    input: String,
    output: i32,
}

#[derive(Debug)]
#[derive(PartialEq)]
enum Action {
    UP,
    DOWN,
    NONE,
}

fn parse_instruction(instruction: &str) -> Action {
    match instruction {
        "(" => Action::UP,
        ")" => Action::DOWN,
        _ => Action::NONE
    }
}

fn calculate_floor(input: &str, floor: i32) -> i32 {
    let action = parse_instruction(input);
    let next_floor = match action {
        Action::UP => floor + 1,
        Action::DOWN => floor - 1,
        _ => floor,
    };
    next_floor
}

fn exec(input: &str, init_floor: i32) -> i32 {
    let mut floor = init_floor;
    for c in input.chars() {
        floor = calculate_floor(&c.to_string(), floor);
    }
    floor
}

pub fn solve() -> i32 {
    let data = utils::read_file("../data/2015_1.txt");
    let floor = exec(&data, INITIAL_FLOOR);
    floor
}

#[test]
fn it_will_convert_instructions_into_enum_value(){
    for c in
        "!@#$%^&*{}[]1234567890abcdefghijklmnopqrstuvwxyz|+_=-`~';:?/.,<>"
        .to_string()
        .chars() {
            assert_eq!(parse_instruction(&c.to_string()), Action::NONE);
    }
}

#[test]
fn it_will_correctly_compute_test_data(){
    let td = [
        TestData { input: String::from("(())"), output: 0 },
        TestData { input: String::from("((("), output: 3 },
        TestData { input: String::from("(()(()("), output: 3 },
        TestData { input: String::from("))((((("), output: 3 },
        TestData { input: String::from("())"), output: -1 },
        TestData { input: String::from("))("), output: -1 },
        TestData { input: String::from(")))"), output: -3 },
        TestData { input: String::from(")())())"), output: -3 },
    ];
    for data in td.iter(){
        assert_eq!(exec(&data.input, INITIAL_FLOOR), data.output);
    }
}
