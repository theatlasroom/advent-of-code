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

let solve = data =>
  Js.Promise.make((~resolve, ~reject) => {
    let result = calculate_floor(~floor=0, ~data, ~index=0);
    resolve(. result);
  });
