package main

type Product struct {
	Id uint64 `gorm:"primaryKey"`
	Name string
	Price int64 
}

func main(){
	
}