exports.p1 = function p1() {
    let iter = []
    for(let a of list(3)) iter.push(a)
    for(let b of list(5)) iter.push(b)
    return iter.filter((v, i, arr) => arr.indexOf(v) === i).reduce((prev, cur) => prev += cur)
}

/**
 * @param {number} multiple
 */
function* list(multiple) {
    let val = multiple
    while(val < 1000) {
        yield val
        val += multiple
    }
}