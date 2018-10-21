const { isNiceString, improvedNiceString } = require('./2015_5')
// ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
// aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
// jchzalrnumimnmhp is naughty because it has no double letter.
// haegwjzuvuyypxyu is naughty because it contains the string xy.
// dvszwmarrgswjxmb is naughty because it contains only one vowel.
const partOneData = [
  ['ugknbfddgicrmopn', true],
  ['aaa', true],
  ['jchzalrnumimnmhp', false],
  ['haegwjzuvuyypxyu', false],
  ['dvszwmarrgswjxmb', false],
]
partOneData.forEach(([testCase, result]) => {
  test(`${testCase} returns ${result}`, () => {
    expect(isNiceString(testCase)).toEqual(result)
  })
})

const partTwoData = [
  ['qjhvhtzxzqqjkmpb', true],
  ['xxyxx', true],
  ['uurcxstgmygtbstg', false],
  ['ieodomkazucvgmuy', false],
]
partTwoData.forEach(([testCase, result]) => {
  test(`${testCase} returns ${result}`, () => {
    expect(improvedNiceString(testCase)).toEqual(result)
  })
})
