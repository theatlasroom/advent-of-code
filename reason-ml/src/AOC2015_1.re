type action =
  | NONE
  | UP
  | DOWN;

type resolve_action = string => action;

let resolve_action = str =>
  switch (str) {
  | "(" => UP
  | ")" => DOWN
  | _ => NONE
  };

type calculate_floor = (string, int, int) => int;

let rec calculate_floor = (~data, ~floor, ~index) => {
  let len = String.length(data);
  if (index < len) {
    let action = resolve_action(String.make(1, data.[index]));
    let next_floor =
      switch (action) {
      | UP => floor + 1
      | DOWN => floor - 1
      | _ => floor
      };
    let next_index = index + 1;
    calculate_floor(~data, ~floor=next_floor, ~index=next_index);
  } else {
    floor;
  };
};

type exec = (string, int) => int;

let exec = (~initial_floor=0, ~data) =>
  calculate_floor(~data, ~floor=initial_floor, ~index=0);

type solve = string => int;

let solve = data => exec(~data);
