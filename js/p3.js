exports.p3 = function p3() {
    return Math.max(sieve(600851475143))
}

/**
 * @param {number} limit
 */
function sieve(limit) {
    let a = new Array(limit).fill(true)
    for(let i = 1; i < Math.sqrt(limit); i++) {
        if(a[i]) {
            for(let j = i**2; j += i; j < n) {
                a[j] = false
            }
        }
    }
    return a.map((v, i) => v ? i + 1 : -1).filter(i => i > 1)
}