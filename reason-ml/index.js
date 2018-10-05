// TODO: make a function to loop over src dir with *.bs.js globbing
const fs = require('fs')
// const path = require('path')
const utils = require('../utils')
const aoc2015_1 = require('./src/AOC2015_1.bs')

const problems = [
  {
    name: '2015 - 1',
    file: '2015_1.txt',
    solver: aoc2015_1,
  },
]

const solve = async ({ name, file, solver }) => {
  console.log(`\nProblem ${name}`)
  const data = await utils.readData(file)
  const result = await solver.solve(data)
  console.log(`Answer: ${result} 🔥\n`)
}

async function solveProblems() {
  console.log(
    '****************************************************************'
  )
  console.log('⚡️🎉 Advent of Code solutions 🎉⚡️')
  console.log('\nLanguage: ReasonML ⚡️⚡️⚡️⚡️⚡️⚡️⚡️⚡️')
  console.log(
    `\nThis is an attempt at solutions to the problems using reasonml,\nthis is purely as a exercise in learning reason, so i am sure \nthe solutions are not "optimal"`
  )
  console.log(
    '****************************************************************'
  )
  try {
    for (let problem of problems) {
      await solve(problem)
    }
  } catch (err) {
    console.error(err)
  }
  console.log(
    '****************************************************************'
  )
  console.log('🎉🎉🎉 Fin 🎉🎉🎉')
  console.log(
    '****************************************************************'
  )
}

async function exec() {
  await solveProblems()
}

exec()
