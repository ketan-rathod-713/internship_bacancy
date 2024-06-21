use("mongodb_documentation")

// db.students.insertMany( [
//     { _id: 1, test1: 95, test2: 92, test3: 90, modified: new Date("01/05/2020") },
//     { _id: 2, test1: 98, test2: 100, test3: 102, modified: new Date("01/05/2020") },
//     { _id: 3, test1: 95, test2: 110, modified: new Date("01/04/2020") }
//  ] )

//  db.students.find()

// Update using aggregation pipeline

// db.students.updateOne( { _id: 3 }, [ { $set: { "test3": 98, modified: "$$NOW"} } ] )

// db.students.find().pretty()

// IMPORTANT
// db.students.updateMany({}, [
//     {
//         $replaceRoot: {
//             newRoot: { $mergeObjects: [{quiz1: 0, quiz2: 0, quiz3: 0}, "$$ROOT"]}
//         }
//     },
//     {$set: {modified: "$$NOW"}}
// ])

db.students.find({})

// DOCS $replaceRoot
// Replaces the input document with the specified document. replace all fields including _id
// you can promote existing embedded document to the top level or create new document for promotion.

// SYNTAX : { $replaceRoot: { newRoot: <replacementDocument> } }

// if replacementDocument is not a document then fails and error.

// db.collection.aggregate([
//     { $replaceRoot: { newRoot: "$name" } }
// ])
// This fails because one of document does not have name field. 

