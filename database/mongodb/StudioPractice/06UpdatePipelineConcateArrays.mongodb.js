use("mongodb_documentation")

// db.students.updateMany(
//     {},
//     [
//         {
//             $set: {tests: [12, 30, 40, 50]}
//         },
//         {
//             $set: {average : {$trunc: [{$avg: "$tests"}, 0]}, modified: "$$NOW"}
//         },
//         {
//             $set: {
//                 grade : {
//                     $switch: {
//                         branches: [
//                             { case: {$gte: ["$average", 90]}, then: "A"},
//                             { case: {$gte: ["$average", 80]}, then: "B"},
//                             { case: {$gte: ["$average", 70]}, then: "C"},
//                             { case: {$gte: ["$average", 60]}, then: "D"}
//                         ],
//                         default: "F"
//                     }
//                 }
//             }
//         }
//     ]
// )

// db.students.find({})

// Switch case
// switch ke andar branches and defaults field he
// branches is array of case conditions 
// for defaults directly set value


// db.students.updateOne( { _id: 2 },
//     [ { $set: { tests: { $concatArrays: [ "$tests", [ 200, 100 ]  ] } } } ]
//   )


// db.students.find({})

// What can be the different methods like $concatArrays, $trunc, $avg of arrays.

