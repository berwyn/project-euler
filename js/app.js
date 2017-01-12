const { p1 } = require('./p1')
const { p2 } = require('./p2')
const { p3 } = require('./p3')

function assert(expected, actual) {
    return expected === actual ? 'Correct' : `Incorrect -- ${actual}`
}

console.log(`P1: ${assert(233168, p1())}`)
console.log(`P2: ${assert(4613732, p2())}`)
console.log(`P3: ${assert(6857, p3())}`)