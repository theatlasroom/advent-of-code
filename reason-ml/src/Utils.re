module Utils = {
  type str_to_list = (string, string) => list(string);
  let str_to_list = (~delimiter="\n", input) =>
    Js.String.split(delimiter, Js.String.trim(input)) |> Array.to_list;
};