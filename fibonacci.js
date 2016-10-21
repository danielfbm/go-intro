

function fibonacci(n) {
    if (n > 1) {
        return fibonacci(n-1) + fibonacci(n-2)
    } 
    return n
}


start = new Date()
for (var i = 0; i < 40; i++) {
    console.log(i, fibonacci(i))
}
end = new Date() - start

console.log(end/1000)