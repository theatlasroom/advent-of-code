/* Santa is delivering presents to an infinite two-dimensional grid of houses.

   He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

   However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

   For example:

   > delivers presents to 2 houses: one at the starting location, and one to the east.
   ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
   ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses. */

/* naive: use a regex to scan the input for all occurrences of the strings we want to find */
/* "better" create a parser */

'use strict'

const utils = require('../utils')
const Actions = {
  UP: '^',
  DOWN: 'v',
  LEFT: '<',
  RIGHTT: '>',
}

const Point = function(opts = { x: 0, y: 0 }) {
  const that = this
  that.coords = { x: 0, y: 0 }

  function init(opts) {
    that.coords = { ...opts }
  }

  init(opts)

  return {
    add: pt => ({ x: pt.x + that.coords.x, y: pt.y + that.coords.y }),
    get: () => that.coords,
    generateKey: () => `x${that.coords.x}y${that.coords.y}`,
  }
}

function matchDirection(action = null) {
  switch (action) {
    case Actions.UP:
      return { x: 0, y: 1 }
    case Actions.DOWN:
      return { x: 0, y: -1 }
    case Actions.LEFT:
      return { x: -1, y: 0 }
    case Actions.RIGHT:
      return { x: 1, y: 0 }
    default:
      return { x: 0, y: 0 }
  }
}

function nextHouse(currentPosition, action) {
  const delta = matchDirection(action)
  return currentPosition.add(delta)
}

function followDirection(directions) {
  const houses = new Set()
  const cursor = 0
  let position = new Point()
  console.log(position)
  while (cursor < directions.length) {
    const action = directions[cursor]
    position = nextHouse(position, action)
    console.log('position', position)
    const key = position.generateKey()
    houses.add(key)
    console.log('key', key, 'size', houses.size)
    cursor++
  }
  return houses.size
}

function calculateDeliveries(directions) {
  return followDirection(directions, 0)
}

const args = process.argv

if (args.length > 2 && args[2]) {
  const file = args[2]
  console.log(`Reading input file ${file}`)
  utils
    .readData(file)
    .then(calculateDeliveries)
    .then(houses => console.log(`Santa delivered to ${houses} unique houses`))
    .catch(console.error)
}

module.exports = calculateDeliveries
