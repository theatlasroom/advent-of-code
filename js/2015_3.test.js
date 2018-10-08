const deliveries = require('./2015_3')
const data = [['>', -1], ['^>v<', 5], ['^v^v^v^v^v', 2]]
data.forEach(([testCase, result]) => {
  test(`${testCase} returns ${result}`, () => {
    expect(deliveries(testCase)).toEqual(result)
  })
})
