use("mongodb_documentation")

db.temperatures.insertMany( [
    { "_id" : 1, "date" : ISODate("2019-06-23"), "tempsC" : [ 4, 12, 17 ] },
    { "_id" : 2, "date" : ISODate("2019-07-07"), "tempsC" : [ 14, 24, 11 ] },
    { "_id" : 3, "date" : ISODate("2019-10-30"), "tempsC" : [ 18, 6, 8 ] }
  ] )

//   In this case addFields is not working may be version is low.
db.temperatures.updateMany({},{ $addFields: {
    "tempsF": {
        $map: {
            input: "$tempsC",
            as: "celsius",
            in: {$add: [{$multiply: ["$$celsius",9/5]}, 32]}
        }
    }
}
})

db.temperatures.find({})