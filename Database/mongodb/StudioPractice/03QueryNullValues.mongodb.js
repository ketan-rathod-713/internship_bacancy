//db.inventory.insertMany([
//   { _id: 1, item: null },
//   { _id: 2 }
//])

// it will find all the null or missing vallues
db.inventory.find({
    item: {
        $eq: null
    }
})

// query documents based on their field existance in document.
db.inventory.find( { item : { $exists: false } } )
