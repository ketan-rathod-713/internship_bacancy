# CSV package

## Reading csv file

r := csv.NewReader(file)
r.Comma = '|' // if we are using delimitor other then comma
r.Comment = '#' // if we don't know what is comment

records, err := r.ReadAll()

- Define newreader
- read, readall will return all records slice of slice


## Writting csv file

1. create csv file using os
2. write to csv file

The Write function writes a single CSV record to writer. A record is a slice of strings with each string being one field. Writes are buffered, so Flush must be called to ensure that the record is written to the underlying writer.

The WriteAll function writes multiple CSV records to the writer using Write and then calls Flush.