use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;

pub fn read_file(filename: &str) -> String {
    let mut f = File::open(filename).expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    String::from(contents)
}

pub fn read_file_by_lines(filename: &str) -> Vec<String> {
    let f = File::open(filename).expect("file not found");
    let file = BufReader::new(&f);
    let mut data = vec![];
    for line in file.lines() {
        let l = line.unwrap();
        data.push(l.to_string());
    }
    data
}
