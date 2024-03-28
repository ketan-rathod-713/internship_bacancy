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


