const path = require('path')
const fs = require('fs')
const { promisify } = require('util')
const dataPath = path.join(__dirname, '../', 'data')
const readFile = promisify(fs.readFile)

const readData = async file => await readFile(path.join(dataPath, file), 'utf8')

module.exports = {
    dataPath,
    readData,
}
