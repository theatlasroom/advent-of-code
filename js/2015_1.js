'use strict'

const utils = require('../utils')
const UP = '('
const DOWN = ')'
const initialFloor = 0
const initialIndex = 0
const args = process.argv

function matchInstruction(input, index, currentFloor) {
    switch (input.charAt(index)) {
        case UP: {
            return currentFloor + 1
        }
        case DOWN: {
            return currentFloor - 1
        }
        default:
            return currentFloor
    }
}

function parseInput(input, index = initialIndex, floor = initialFloor) {
    while (index < input.length) {
        floor = matchInstruction(input, index, floor)
        index++
    }
    return floor
}

function parseFloorInstructions(data = '') {
    if (data.length < 1) return 0
    return parseInput(data, initialIndex, initialFloor)
}

if (args.length > 2 && args[2]) {
    const file = args[2]
    console.log(`Reading input file ${file}`)
    utils
        .readData(file)
        .then(parseFloorInstructions)
        .then(floor => console.log(`Santa is on floor ${floor}`))
        .catch(console.error)
}

module.exports = parseFloorInstructions
