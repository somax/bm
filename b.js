console.time('b')
let n = 0
for (let i = 0; i < 100000000; i++){
   n += 1+1-1*1/1
}
let t = Date.now()
console.timeEnd('b')
console.log(n)
