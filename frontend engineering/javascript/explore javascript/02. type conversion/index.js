// in js things gets automatically converted to desired types. so we should be very careful with that.

console.log("conversion in javascript")

console.log(1 + 3)
console.log(3 + "5")
console.log("1"+ "3")

console.log(+true) // will output 1

console.log(true + 1) // will output 2

console.log(1 + "true") // will output "1true"


let num1, num2, num3

num1 = num2 = num3 = 3 + 2
console.log(num1, num2, num3)

let gameCounter = 100
gameCounter++

console.log(gameCounter)