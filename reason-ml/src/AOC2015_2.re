/*
 -- Day 2: I Was Told There Would Be No Math ---
 The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

 Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

 For example:

 A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
 A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
 All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order? */

type present = {
  h: int,
  w: int,
  l: int,
};

type parse_paper = string => present;
let parse_paper = str => {
  let arr = Js.String.split("x", str);
  let mapped = arr |> Array.map(int_of_string) |> Array.to_list;
  if (List.length(mapped) >= 3) {
    let l = List.nth(mapped, 0);
    let w = List.nth(mapped, 1);
    let h = List.nth(mapped, 2);
    {h, w, l};
  } else {
    {h: 0, w: 0, l: 0};
  };
};

let int_comparator = (a: int, b: int) => a > b ? 1 : (-1);

type calculate_paper = present => int;
let calculate_paper = p => {
  let side_a = p.l * p.w;
  let side_b = p.l * p.h;
  let side_c = p.h * p.w;
  let sides: list(int) =
    List.sort(int_comparator, [side_a, side_b, side_c]);
  let smallest = List.hd(sides);
  List.fold_left((acc, item) => acc + 2 * item, 0, sides) + smallest;
};

/* use tail recursion to pop off the head and calculate, returning the total when we have no items left in the list */
type calculate_wrapping_paper = (list(string), int) => int;
let rec calculate_wrapping_paper = (~data, ~result) => 
	switch (data) {
  | [] => result
  | [dimensions, ...rest] =>
    let p = parse_paper(dimensions);
    let acc = result + calculate_paper(p);
    calculate_wrapping_paper(~data=rest, ~result=acc);
  };

/* convert the data string into a list */
let solve = rawdata => {
  let data = Js.String.split("\n", Js.String.trim(rawdata)) |> Array.to_list;
  Js.Promise.resolve(
    {
      let wrapping_paper_used = calculate_wrapping_paper(~data, ~result=0);
      wrapping_paper_used;
    },
  );
};
