/* Santa is delivering presents to an infinite two-dimensional grid of houses.

   He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

   However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

   For example:

   > delivers presents to 2 houses: one at the starting location, and one to the east.
   ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
   ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses. */

type move =
  | NORTH
  | EAST
  | SOUTH
  | WEST
  | NOOP;

let total_deliveries: int = 0;

type house = {
  x: int,
  y: int,
};

module type Comparable = {
  type t;
  let equal: (t, t) => bool;
};

module MakeSet = (Item: Comparable) => {
  /* let's use a list as our naive backing data structure */
  type backingType = list(Item.t);
  let empty = [];
  let add = (currentSet: backingType, newItem: Item.t): backingType =>
    /* if item exists */
    if (List.exists(x => Item.equal(x, newItem), currentSet)) {
      currentSet; /* return the same (immutable) set (a list really) */
    } else {
      [
        newItem,
        ...currentSet /* prepend to the set and return it */
      ];
    };
};

module House = {
  type t = house;
  let equal = (a, b) => a.x == b.x && a.y == b.y;
  let create = (x, y) => {x, y};
};

/* create a custom set for our houses so we can properly model the unique houses delivered to */
/* use the Set.Make functor */
module HouseSet = MakeSet(House);

type move_house = (move, house) => house;
let move_house = (action, location) => {
  let {x: curr_x, y: curr_y} = location;
  switch (action) {
  | NORTH => {x: curr_x, y: curr_y + 1}
  | EAST => {x: curr_x + 1, y: curr_y}
  | SOUTH => {x: curr_x, y: curr_y - 1}
  | WEST => {x: curr_x - 1, y: curr_y}
  | NOOP => {x: curr_x, y: curr_y}
  };
};

/* match an instruction */
type match_instruction = string => move;
let match_instruction = instruction =>
  switch (instruction) {
  | "^" => NORTH
  | ">" => EAST
  | "v" => SOUTH
  | "<" => WEST
  | _ => NOOP
  };

type next_delivery = (~instruction: string, ~location: house) => house;
let next_delivery = (~instruction, ~location) => {
  let action = match_instruction(instruction);
  move_house(action, location);
};

type deliver_presents =
  (
    ~deliveries_made: list(House.t),
    ~instructions: string,
    ~index: int,
    ~location: house
  ) =>
  list(House.t);
let rec deliver_presents =
        (~deliveries_made, ~instructions, ~index, ~location) =>
  if (index < String.length(instructions)) {
    let instruction = String.make(1, instructions.[index]);
    let next_location = next_delivery(~instruction, ~location);
    let deliveries_made = HouseSet.add(deliveries_made, next_location);
    deliver_presents(
      ~deliveries_made,
      ~instructions,
      ~index=index + 1,
      ~location=next_location,
    );
  } else {
    deliveries_made;
  };

/* convert the data string into a list, use destructuring to 'pop' elements off the list */
let solve = data =>
  Js.Promise.resolve(
    {
      let deliveries_made = HouseSet.(empty);

      let result =
        deliver_presents(
          ~deliveries_made,
          ~instructions=data,
          ~index=0,
          ~location={x: 0, y: 0},
        );
      result |> List.length;
    },
  );