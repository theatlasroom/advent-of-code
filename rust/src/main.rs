mod aoc2015_1;
mod aoc2015_2;
mod aoc2015_3;
mod aoc2015_4;
// mod aoc2015_5;
mod aoc2015_6;
mod aoc2018_1;
mod aoc2018_2;
mod utils;

#[macro_use]
extern crate lazy_static;
extern crate md5;
extern crate regex;

fn banner() {
    println!("****************************************************************");
    println!("⚡️🎉 Advent of Code solutions 🎉⚡️");
    println!("\nLanguage: Rust ⚡️⚡️⚡️⚡️⚡️⚡️⚡️⚡️⚡️⚡️");
    println!("\nThis is an attempt at solutions to the problems using rust,\nthis is purely as a exercise in learning rust, so i am sure \nthe solutions are not \"optimal\"");
    println!("****************************************************************");
}

fn closing_banner() {
    println!("****************************************************************");
    println!("🎉🎉🎉 Fin 🎉🎉🎉");
    println!("****************************************************************");
}

fn main() {
    banner();
    println!("\n2015 - 🎄🎄🎄🎄🎄🎄🎄🎄🎄");
    println!("Puzzle 1: {}", aoc2015_1::solve());
    println!("Puzzle 2: {}", aoc2015_2::solve());
    println!("Puzzle 3: {}", aoc2015_3::solve());
    // ignored because its really slow to run
    // println!("Puzzle 4: {}", aoc2015_4::solve());
    // println!("Puzzle 5: {}", aoc2015_5::solve());
    // println!("Puzzle 6: {}", aoc2015_6::solve());
    println!("\n2018 - 🎄🎄🎄🎄🎄🎄🎄🎄🎄");
    println!("Puzzle 1: {}", aoc2018_1::solve());
    println!("Puzzle 2: {}", aoc2018_2::solve());
    println!("\n");
    closing_banner();
}
