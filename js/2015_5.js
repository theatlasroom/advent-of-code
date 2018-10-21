'use strict'

// Santa needs help figuring out which strings in his text file are naughty or nice.

// A nice string is one with all of the following properties:

// It contains at least three vowels(aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
// It contains at least one letter that appears twice in a row, like xx, abcdde(dd), or aabbccdd(aa, bb, cc, or dd).
// It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.

const utils = require('../utils')
const args = process.argv

const PARSER_ONE = 'PARSER_ONE'
const PARSER_TWO = 'PARSER_TWO'

// const threeVowels = str, pattern

// isNiceString = string => bool
function isNiceString(str = '') {
  if (!str || !str.length) return false
  const flags = 'gi'
  const THREE_VOWELS = new RegExp('([a|e|i|o|u].*){3,}')
  // capture any of our bad strings
  const BAD_STRINGS = new RegExp('(ab|cd|pq|xy)')
  // capture any letter, then repeat the captured text once
  const DOUBLE_CHARACTER = new RegExp('([A-Za-z])\\1')
  return (
    THREE_VOWELS.test(str) &&
    DOUBLE_CHARACTER.test(str) &&
    !BAD_STRINGS.test(str)
  )
}

// isNiceString = string => bool
function improvedNiceString(str = '') {
  if (!str || !str.length) return false
  const flags = 'gi'
  // a letter, repeated, with a different letter in between
  // ie xyx
  const REPEATED_LETTER_SEPARATED_BY_ANOTHER = new RegExp('([A-Za-z]).\\1')
  // two letter sequence, repeated
  // ie xyfkdkelxy
  const REPEATED_TWO_LETTERS = new RegExp('([A-Za-z]{2}).*\\1')
  return (
    REPEATED_TWO_LETTERS.test(str) &&
    REPEATED_LETTER_SEPARATED_BY_ANOTHER.test(str)
  )
}

// parseInput = Array => Number
function parseInput(dataArr, parser) {
  const parsingFunction =
    parser == 'PARSER_ONE' ? isNiceString : improvedNiceString
  return dataArr
    .map(parsingFunction)
    .reduce((acc, next) => (!!next ? acc + 1 : acc), 0)
}

function countNiceStrings(data = [], parser = PARSER_ONE) {
  const dataArr = Array.isArray(data) ? data : [data]
  if (!data || dataArr.length < 1) return 0
  return parseInput(data, parser)
}

if (args.length > 2 && args[2]) {
  const file = args[2]
  console.log(`Reading input file ${file}`)
  utils
    .readData(file)
    .then(inputData => {
      const p1 = countNiceStrings(inputData.split('\n'), PARSER_ONE)
      const p2 = countNiceStrings(inputData.split('\n'), PARSER_TWO)
      return Promise.all([p1, p2])
    })
    .then(([part1, part2]) => {
      console.log(`Part 1: there are ${part1} nice strings in the text`)
      console.log(`Part 2: there are ${part2} nice strings in the text`)
    })
    .catch(console.error)
}

module.exports = {
  countNiceStrings,
  isNiceString,
  improvedNiceString,
}
