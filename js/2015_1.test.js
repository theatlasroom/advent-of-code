const p = require('./2015_1');
const data = [
 ["(())", 0],
 ["(((", 3 ],
 ["(()(()(", 3],
 ["))(((((", 3],
 ["())", -1],
 ["))(", -1],
 [")))", -3],
 [")())())", -3]
];
data.forEach(([testCase, result]) => {
 test(`${testCase} returns ${result}`, () => {
  expect(p(testCase)).toEqual(result);
 });
});
