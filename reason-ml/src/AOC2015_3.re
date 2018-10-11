/* Santa is delivering presents to an infinite two-dimensional grid of houses.

   He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

   However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

   For example:

   > delivers presents to 2 houses: one at the starting location, and one to the east.
   ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
   ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses. */

/* naive: use a regex to scan the input for all occurrences of the strings we want to find ?? */
/* track (x,y) positions */


type action =
  | DELIVER
  | DELIVER_AND_MOVE
  | MOVE;

let total_deliveries: int = 0;

/*
 string: Pattern to match
 int: # of presents delivered
  */
let patterns: list((string, int)) = [
  (">", 2),
  (">v<", 5),
  ("^v^v^v^v^v", 2),
];

Js.String.match_

let parse_regex = (pattern, text) => {

}

type calculate_deliveries: string => int;
let calculate_deliveries = input => {
  List.map
  let deliveries_made = List.map((pattern) => parse_regex(pattern, input)), patterns);
  /* reduce the deliveries and add  */
  /* total_deliveries = reduce(deliveries_made) */
  /* total_deliveries + 1 */
  0
}

let solve = data =>
  Js.Promise.resolve(
    {
      let result = calculate_deliveries();
      result;
    },
  );