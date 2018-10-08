const deliveries = require('./2015_3')
const data = [['>', 2], ['^>v<', 4], ['^v^v^v^v^v', 2]]
data.forEach(([testCase, result]) => {
  test(`${testCase} returns ${result}`, () => {
    expect(deliveries(testCase)).toEqual(result)
  })
})
