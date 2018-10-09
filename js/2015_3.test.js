const deliveries = require('./2015_3')
const data = [['>', 2], ['^>v<', 4], ['^v^v^v^v^v', 2]]
data.forEach(([testCase, result]) => {
  test(`${testCase} returns ${result}`, () => {
    expect(deliveries(testCase)).toEqual(result)
  })
})

const deliveriesPartTwo = require('./2015_3_2')
const dataPartTwo = [['^v', 3], ['^>v<', 3], ['^v^v^v^v^v', 11]]
dataPartTwo.forEach(([testCase, result]) => {
  test(`${testCase} returns ${result}`, () => {
    expect(deliveriesPartTwo(testCase)).toEqual(result)
  })
})
