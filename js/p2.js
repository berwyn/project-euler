function fib(x) {
    if(x < 2) {
        return 1
    } else {
        return fib(x - 2) + fib(x - 1)
    }
}

exports.p2 = function p2() {
    let items = []

    for(let i = 0; true; i++) {
        let a = fib(i)
        if(a <= 4000000) {
            items.push(a)
        } else {
            break
        }
    }
    
    return items.filter(i => i%2 === 0).reduce((prev, cur) => prev += cur)
}