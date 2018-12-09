/* To make sure you didn't miss any, you scan the likely candidate boxes again, counting the number that have an ID containing exactly two of any letter and then separately counting those with exactly three of any letter. You can multiply those two counts together to get a rudimentary checksum and compare it to what your device predicts.

   For example, if you see the following box IDs:

   abcdef contains no letters that appear exactly two or three times.
   bababc contains two a and three b, so it counts for both.
   abbcde contains two b, but no letter appears exactly three times.
   abcccd contains three c, but no letter appears exactly two times.
   aabcdd contains two a and two d, but it only counts once.
   abcdee contains two e.
   ababab contains three a and three b, but it only counts once.
   Of these box IDs, four of them contain a letter which appears exactly twice, and three of them contain a letter which appears exactly three times. Multiplying these together produces a checksum of 4 * 3 = 12.

   What is the checksum for your list of box IDs? */
open Utils;
type repeatedCharacters = {
  two: int,
  three: int,
};

type count_occurrences = (string, string) => int;
let count_occurrences = (str, delimiter) =>
  (str |> Utils.str_to_list(~delimiter) |> List.length) - 1;

type sum_repeated_characters_in_string = string => repeatedCharacters;
let sum_repeated_characters_in_string = curr_str =>
  List.fold_left(
    (acc, curr_char) => {
      let {two, three} = acc;
      switch (count_occurrences(curr_str, curr_char)) {
      | 2 => {two: 1, three}
      | 3 => {two, three: 1}
      | _ => {two, three}
      };
    },
    {two: 0, three: 0},
    curr_str |> Utils.str_to_list(~delimiter=""),
  );

type calculate_checksum = list(string) => int;
let calculate_checksum = data => {
  let counts =
    data
    |> Utils.str_to_list
    |> List.map(sum_repeated_characters_in_string)
    |> List.fold_left(
         (acc, curr) => {
           two: acc.two + curr.two,
           three: acc.three + curr.three,
         },
         {two: 0, three: 0},
       );
  counts.two * counts.three;
};

let solve = (input: string) =>
  Js.Promise.resolve(
    {
      let checksum = calculate_checksum(input);
      checksum;
    },
  );