'use strict'

const utils = require('../utils')
const INC = '+'
const DEC = '-'
const initialFreq = 0
const args = process.argv

function matchInstruction(input, freq) {
  const amount = Number(input.slice(1))
  switch (input.charAt(0)) {
    case INC: {
      return freq + amount
    }
    case DEC: {
      return freq - amount
    }
    default:
      return freq
  }
}

const stringToArray = (str, delim = '\n') => str.split(delim)

const calculateResultingFrequency = (freq, input) =>
  matchInstruction(input, freq)

const calculateFirstRepeatedFrequency = (data, currFreq = 0) => {
  const visited = new Set([])
  let i = 0
  while (!visited.has(currFreq)) {
    visited.add(currFreq)
    currFreq = matchInstruction(data[i], currFreq)
    i = i < data.length - 1 ? i + 1 : 0
  }
  return currFreq
}

if (args.length > 2 && args[2]) {
  const file = args[2]
  utils
    .readData(file)
    .then(stringToArray)
    .then(arr => [
      arr.reduce(calculateResultingFrequency, initialFreq),
      calculateFirstRepeatedFrequency(arr, initialFreq),
    ])
    .then(([freq, firstRepeat]) =>
      console.log(
        `Resuting freq is ${freq}\nFirst repeated frequency is ${firstRepeat}`
      )
    )
    .catch(console.error)
}
