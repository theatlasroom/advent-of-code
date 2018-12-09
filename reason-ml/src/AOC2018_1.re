open Utils;

type action =
  | INC
  | DEC
  | NOOP;

type instruction = {
  action,
  amount: int,
};

type match_instruction = string => instruction;
let match_instruction = input => {
  let charlist = Utils.str_to_list(~delimiter="", input);
  switch (charlist) {
  | [] => {action: NOOP, amount: 0}
  | [act, ...value] =>
    let action = act == "+" ? INC : DEC;
    let amount = value |> String.concat("") |> int_of_string;
    {action, amount};
  };
};

type calculate_next_value = (int, instruction) => int;
let calculate_next_value = (~current_value=0, ins) =>
  switch (ins.action) {
  | INC => current_value + ins.amount
  | DEC => current_value - ins.amount
  | _ => current_value
  };

type calibrate = (~data: list(string), ~result: int) => int;
let rec calibrate = (~data, ~result) =>
  switch (data) {
  | [] => result
  | [head, ...tail] =>
    let next =
      calculate_next_value(~current_value=result, match_instruction(head));
    calibrate(~data=tail, ~result=next);
  };

type first_repeated_frequency =
  (
    ~curr_freq: int,
    ~curr_index: int,
    ~data: list(string),
    ~freq_table: Hashtbl.t(int, int)
  ) =>
  int;

let rec first_repeated_frequency =
        (~curr_freq, ~curr_index, ~data, ~freq_table) =>
  switch (Hashtbl.find(freq_table, curr_freq)) {
  | exception Not_found =>
    Hashtbl.add(freq_table, curr_freq, curr_freq);
    let next =
      calculate_next_value(
        ~current_value=curr_freq,
        match_instruction(List.nth(data, curr_index)),
      );
    let next_index = curr_index < List.length(data) - 1 ? curr_index + 1 : 0;
    first_repeated_frequency(
      ~curr_freq=next,
      ~curr_index=next_index,
      ~data,
      ~freq_table,
    );
  | freq => freq
  };

let solve = (input: string) =>
  Js.Promise.resolve(
    {
      let data = input |> Utils.str_to_list(~delimiter="\n");
      let resulting_frequency = calibrate(~data, ~result=0);
      let initial_size = 2 * List.length(data);
      let repeated_frequency =
        first_repeated_frequency(
          ~curr_freq=0,
          ~curr_index=0,
          ~data,
          ~freq_table=Hashtbl.create(initial_size),
        );
      (resulting_frequency, repeated_frequency);
    },
  );