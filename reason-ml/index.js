// TODO: make a function to loop over src dir with *.bs.js globbing
const fs = require('fs')
// const path = require('path')
const utils = require('../utils')

const problems = [
  {
    year: 2015,
    day: 1,
  },
  {
    year: 2015,
    day: 2,
  },
  {
    year: 2015,
    day: 3,
  },
  {
    year: 2015,
    day: 4,
  },
  // {
  //   year: 2018,
  //   day: 1,
  // },
  {
    year: 2018,
    day: 2,
  },
]

const solve = async ({ year, day, file }) => {
  console.log(`\nProblem ${year} - ${day}`)
  const solver = require(`./src/AOC${year}_${day}.bs`)
  const dataFile = `${year}_${day}.txt`
  const data = await utils.readData(dataFile)
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
