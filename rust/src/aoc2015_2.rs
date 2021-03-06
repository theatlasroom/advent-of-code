use utils;

const SEPARATOR: char = 'x';

#[derive(Debug, PartialEq)]
struct SantaPresent {
    width: i32,
    height: i32,
    length: i32,
}

fn parse_line(line: &str, separator: char) -> SantaPresent {
    // TODO: LOOOOOL this seems bad, very very bad
    let mut chars = line.split(separator);
    let length = chars.next().unwrap().parse::<i32>().unwrap();
    let width = chars.next().unwrap().parse::<i32>().unwrap();
    let height = chars.next().unwrap().parse::<i32>().unwrap();
    SantaPresent {
        length,
        width,
        height,
    }
}

#[test]
fn it_will_correctly_parse_a_line_of_input() {
    assert_eq!(
        parse_line("29x13x26", SEPARATOR),
        SantaPresent {
            width: 13,
            length: 29,
            height: 26
        }
    )
}

// TODO: should be implemented for the struct...
fn compute_surface_area(cuboid: &SantaPresent) -> i32 {
    // compute all the faces
    // sort to find the smallest face
    let mut faces = vec![
        cuboid.length * cuboid.width,
        cuboid.width * cuboid.height,
        cuboid.height * cuboid.length,
    ];
    faces.sort();
    let min = faces[0];
    ((2 * faces[0]) + (2 * faces[1]) + (2 * faces[2]) + min)
}

#[test]
fn it_will_compute_surface_area() {
    let b = SantaPresent {
        length: 2,
        width: 3,
        height: 4,
    };
    assert_eq!(compute_surface_area(&b), 58);
    let b2 = SantaPresent {
        length: 1,
        width: 1,
        height: 10,
    };
    assert_eq!(compute_surface_area(&b2), 43);
}

// TODO: should be implemented for the struct...
fn compute_bow_ribbon_length(cuboid: &SantaPresent) -> i32 {
    // compute all the faces
    // sort to find the smallest face
    let mut faces = vec![cuboid.length, cuboid.width, cuboid.height];
    faces.sort();
    2 * (faces[0] + faces[1])
}

#[test]
fn it_will_compute_bow_ribbon_length_required() {
    let b = SantaPresent {
        length: 2,
        width: 3,
        height: 4,
    };
    assert_eq!(compute_bow_ribbon_length(&b), 10);
    let b2 = SantaPresent {
        length: 1,
        width: 1,
        height: 10,
    };
    assert_eq!(compute_bow_ribbon_length(&b2), 4);
}

// TODO: should be implemented for the struct...
fn compute_wrapping_ribbon_length(cuboid: &SantaPresent) -> i32 {
    // compute all the faces
    // sort to find the smallest face
    let mut faces = vec![cuboid.length, cuboid.width, cuboid.height];
    faces.sort();
    faces[0] * faces[1] * faces[2]
}

#[test]
fn it_will_compute_wrapping_ribbon_length_required() {
    let b = SantaPresent {
        length: 2,
        width: 3,
        height: 4,
    };
    assert_eq!(compute_wrapping_ribbon_length(&b), 24);
    let b2 = SantaPresent {
        length: 1,
        width: 1,
        height: 10,
    };
    assert_eq!(compute_wrapping_ribbon_length(&b2), 10);
}

fn calculate_wrapping_paper(_data: &str) -> i32 {
    let lines = utils::read_file_by_lines("../data/2015_2.txt");
    let mut sum: i32 = 0;
    for line in lines {
        let b = parse_line(&line, SEPARATOR);
        sum += compute_surface_area(&b);
    }
    sum
}

fn calculate_total_ribbon_needed(_data: &str) -> i32 {
    let lines = utils::read_file_by_lines("../data/2015_2.txt");
    let mut sum: i32 = 0;
    for line in lines {
        let b = parse_line(&line, SEPARATOR);
        sum += compute_wrapping_ribbon_length(&b) + compute_bow_ribbon_length(&b);
    }
    sum
}

pub fn solve() -> String {
    let data = utils::read_file("../data/2015_1.txt");
    let wrapping_paper = calculate_wrapping_paper(&data);
    let ribbon = calculate_total_ribbon_needed(&data);
    format!(
        "{:?} square feet of wrapping paper and {:?} feet of ribbon",
        wrapping_paper, ribbon
    )
}
