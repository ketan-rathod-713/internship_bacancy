//db.inventory.insertMany( [
//   { item: "journal", instock: [ { warehouse: "A", qty: 5 }, { warehouse: "C", qty: 15 } ] },
//   { item: "notebook", instock: [ { warehouse: "C", qty: 5 } ] },
//   { item: "paper", instock: [ { warehouse: "A", qty: 60 }, { warehouse: "B", qty: 15 } ] },
//   { item: "planner", instock: [ { warehouse: "A", qty: 40 }, { warehouse: "B", qty: 5 } ] },
//   { item: "postcard", instock: [ { warehouse: "B", qty: 15 }, { warehouse: "C", qty: 35 } ] }
//]);

//db.inventory.find({})

// Find documents whose size is 1
//db.inventory.find({
//    instock: {
//        $size: 1
//    }
//})

// Query embedded documents
//db.inventory.find({
//    "instock.qty": 5
//})

// i want ki instock array ka 0th element only match hona chahiye then do this
//db.inventory.find({
//    "instock.0.qty": 5
//})

// Specify multiple conditions for array of documents

// Use $elemMatch operator to specify multiple criteria on an array of embedded documents such that at least one embedded document satisfies all the specified criteria.
//db.inventory.find({
//    instock: {
//        $elemMatch : {
//            qty: 5,
//            warehouse: "A"
//        }
//    }
//})



