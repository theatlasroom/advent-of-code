const p = require('./2015_1')
const partTwo = require('./2015_1_2')
const data = [
  ['(())', 0],
  ['(((', 3],
  ['(()(()(', 3],
  ['))(((((', 3],
  ['())', -1],
  ['))(', -1],
  [')))', -3],
  [')())())', -3],
]
data.forEach(([testCase, result]) => {
  test(`${testCase} returns ${result}`, () => {
    expect(p(testCase)).toEqual(result)
  })
})

const dataPartTwo = [[')', 1], ['()())', 5], ,]
dataPartTwo.forEach(([testCase, result]) => {
  test(`${testCase} returns ${result}`, () => {
    expect(partTwo(testCase)).toEqual(result)
  })
})
