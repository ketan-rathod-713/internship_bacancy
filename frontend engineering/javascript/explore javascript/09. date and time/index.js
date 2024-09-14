// date
// milliseconds refers to the time from January 1, 1970 UTC

let mydate = new Date()
console.log(mydate)

// conversion is important in date

console.log(mydate.toString())
console.log(mydate.toDateString()) // only date
console.log(mydate.toISOString())
console.log(mydate.toLocaleDateString()) // eg. 8/1/2024

console.log(typeof mydate) // object

// IMP: month start from 0 in javascript ha ha
let createdDate = new Date(2023, 0, 23) // year, month, date, hours, minutes, seconds, milliseconds
console.log(createdDate.toDateString())
console.log(createdDate.toISOString())

// year-month-date
// or we can do date-month-year
let anotherDate = new Date("2023/01/12")
console.log(anotherDate.toDateString())

let myTimeStamp = Date.now()
console.log(myTimeStamp) // milliseconds from 1 january 1970
console.log(createdDate.getTime())

// how to convert milliseconds to convert to seconds or hours
// converting to seconds
console.log(Math.floor(Date.now() / 1000))

const newDate = new Date()

// to customize the formate
console.log(newDate.toLocaleString('default', {
    weekday: "long"
}))