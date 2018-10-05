mod utils;
mod aoc2015_1;
mod aoc2015_2;

fn banner(){
    println!("****************************************************************");
    println!("⚡️🎉 Advent of Code solutions 🎉⚡️");
    println!("\nLanguage: Rust ⚡️⚡️⚡️⚡️⚡️⚡️⚡️⚡️⚡️⚡️");
    println!("\nThis is an attempt at solutions to the problems using reasonml,\nthis is purely as a exercise in learning reason, so i am sure \nthe solutions are not \"optimal\"");
    println!("****************************************************************");
}

fn closing_banner(){
    println!("****************************************************************");
    println!("🎉🎉🎉 Fin 🎉🎉🎉");
    println!("****************************************************************");
}

fn main() {
    banner();
    println!("\n2015");
    println!("Puzzle 1: {}", aoc2015_1::solve());
    println!("Puzzle 2: {}", aoc2015_2::solve());
    println!("\n");
    closing_banner();
}
