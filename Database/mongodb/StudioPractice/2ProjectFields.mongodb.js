
// Second argument is we can project fields.
db.inventory.find({}, {
    _id: 1
})

// We can also suppress specific fields
db.inventory.find({}, {
    _id: 0,
    "instock.qty": 0
})