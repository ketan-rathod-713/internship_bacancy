package main

import (
	"csv/jsontwocsv"
)

func main() {
	// encodingcsv.ReadCsv()
	// encodingcsv.WriteToCsv()

	// The GoCSV package aims to provide easy serialization and deserialization functions to use CSV in Golang

	// MarhalFile Method takes data and saves it to csv file
	// gocsvpackage.Marshaling()
	// gocsvpackage.Unmarshalling()

	// only change file mode to append mode and it will work fine
	// gocsvpackage.AppendingToFile()

	// Other functions of gocsv

	// MarshalString to marhsal content to string
	// customizable csv interfface methods

	// MarshalCSV and UnmarshalCSV method on particular struct.

	// CSV TO MAPS
	// gocsvpackage.CsvToMaps()
	// /* Same with csvToMap but it can be used if we have two columns only */
	// gocsvpackage.Unmarshal() // TODO it takes input of io.Reader whilte UnmarshalCSV takes input of CSVReader
	// TODO same thing when it comes to marshal // It outputs io.Writter while UnmarshalCSV will give output of CSVWritter // MarshalFile for writting to file

	// ? Different Methods
	/*
		? 1. Marshal : out : io.Writter
		? 2. MarshalBytes : out is []bytes
		? 3. MarshalCSV : out type is CSVWriter
		? 4. MarshalFile : out type is os.File
		? 5. MarshalString: out type is string

		! SetCSV Reader ?? and writter
		! SetHeaderNormalizer

		? Same thing for unmarshal

		? 1.Unmarshal ( in io.Reader)
		? 2.UnmarshalBytes
		? 3.UnmarshalCSV
		? 4.UnmarshalCSVToMap
		? 5.UnmarshalFile
		? 6. UnmarshalString also there

		! UnmarshalDecoder
		! UnmarshalMultipartFile





	*/

	// gocsvinchunks.ReadCsvFileInChunks()
	// gocsvinchunks.ReadCsvLineByLine()
	// json to csv
	jsontwocsv.JsonToCsv()

	// same for csv to json
	// csv to json
}
