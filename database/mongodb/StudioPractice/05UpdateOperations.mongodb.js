use("updateDocumentsMongodb2")

// db.inventory.insertMany( [
//     { item: "canvas", qty: 100, size: { h: 28, w: 35.5, uom: "cm" }, status: "A" },
//     { item: "journal", qty: 25, size: { h: 14, w: 21, uom: "cm" }, status: "A" },
//     { item: "mat", qty: 85, size: { h: 27.9, w: 35.5, uom: "cm" }, status: "A" },
//     { item: "mousepad", qty: 25, size: { h: 19, w: 22.85, uom: "cm" }, status: "P" },
//     { item: "notebook", qty: 50, size: { h: 8.5, w: 11, uom: "in" }, status: "P" },
//     { item: "paper", qty: 100, size: { h: 8.5, w: 11, uom: "in" }, status: "D" },
//     { item: "planner", qty: 75, size: { h: 22.85, w: 30, uom: "cm" }, status: "D" },
//     { item: "postcard", qty: 45, size: { h: 10, w: 15.25, uom: "cm" }, status: "A" },
//     { item: "sketchbook", qty: 80, size: { h: 14, w: 21, uom: "cm" }, status: "A" },
//     { item: "sketch pad", qty: 95, size: { h: 22.85, w: 30.5, uom: "cm" }, status: "A" }
//  ] );

db.inventory.find({})

// To update a document, MongoDB provides update operators, such as $set, to modify field values.
// Some update operators will create field if it is not present such as $set

db.inventory.updateOne({
    item: "paper"
}, {
    $set: {"size.uom": "cm", status: "p"},
    $currentDate : {lastModified: true}
})

// uses the $currentDate operator to update the value of the lastModified field to the current date. If lastModified field does not exist, $currentDate will create the field. See $currentDate for details.

// updateMany will update multiple documents based on the condition

// To replace the entire document except for the _id . use replaceOne function

// during replacement it should not contain any update operators.

db.inventory.replaceOne(
    { item: "paper" },
    { item: "paper", instock: [ { warehouse: "A", qty: 60 }, { warehouse: "B", qty: 40 } ] }
 )


//  Upsert Option
// If updateOne(), updateMany(), or replaceOne() includes upsert : true and no documents match the specified filter, then the operation creates a new document and inserts it.

// Updates with aggregation pipeline

// various operators are 
// $addFields, $set, $project, $unset, $replaceRoot, $replaceWith