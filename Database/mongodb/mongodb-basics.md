# Mongodb

## Insert Documents

- InsertOne
- InsertMany etc.

## Query Documents

- to specify equality conditon <field>:<value>

```
cursor, err := coll.Find(
	context.TODO(),
	bson.D{{"status", "D"}},
)
```

### Specify Condition using query operators

```
{ status: { $in: [ "A", "D" ] } }
```

Is equivalent to 

`SELECT * FROM inventory WHERE status in ("A", "D")`

https://www.mongodb.com/docs/manual/reference/operator/query/

### Specify and, or conditions

```
   $or: [
     { qty: { $lt: 30 } }, { item: { $regex: '^p' } }
   ]
```

### Select nested structures

```
cursor, err := coll.Find(
	context.TODO(),
	bson.D{
		{"size.h", bson.D{
			{"$lt", 15},
		}},
		{"size.uom", "in"},
		{"status", "D"},
	})
```
### Query an array

1. Match an array



2. Get all documents which has exactly two elements red and black in any order

```
bson.D{
		{"tags", bson.D{{"$all", bson.A{"red", "blank"}}}},
	})
```

3. for specific order

```
bson.D{{"tags", bson.A{"red", "blank"}}},
Above is same as "tags":["red", "black"] in node js
```

4. Query array for an element

if atleast one element then will work such queries

```
bson.D{
		{"dim_cm", bson.D{
			{"$gt", 15},
			{"$lt", 20},
		}},
	}
```

dim_cm is array here.

### Query an array of embedded documents

Continue from here


# Mongodb Docs Important


## Insert document 

db.collection.insertOne() // Inserts a single document into a
collection.
db.collection.insertMany() // Inserts multiple documents into a collection.
Additional Methods for Inserts
The following methods can also add new documents to a collection:

db.collection.updateOne() when used with the upsert: true option.

db.collection.updateMany() when used with the upsert: true option.

db.collection.findAndModify() when used with the upsert: true option.

db.collection.findOneAndUpdate() when used with the upsert: true option.

db.collection.findOneAndReplace() when used with the upsert: true option.

db.collection.bulkWrite().

## Query documents

db.inventory.find( { status: "D" } )

## Specify condition using query operators
db.inventory.find( { status: { $in: [ "A", "D" ] } } )

And condtion
db.inventory.find( { status: "A", qty: { $lt: 30 } } )

Or condition ke liye bahar $or likhna padega
we can do both nested.

Query embedded nested documents
db.inventory.find( { "size.uom": "in" } )

Specify match on nested fields.
db.inventory.find( { size: { h: 14, w: 21, uom: "cm" } } )

Find documents with exact this array.
db.inventory.find( { tags: ["red", "blank"] } )

Find Array with this two values, Array which contains both the elements.
db.inventory.find( { tags: { $all: ["red", "blank"] } } )


Quering array for an element
db.inventory.find( { tags: "red" } )
we can same do with all others...

Compound filter operations
db.inventory.find( { dim_cm: { $gt: 15, $lt: 20 } } )


