package entity

// Here binding tag is used to validate the struct fields

// when to use validate field ??

// What if we want to have a custom validator
// What if i write validate:"is-cool" and my custom function validate it ha ha let's do it

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int    `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=2,max=10"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}

/*
	{
	  "title": "good",
	  "description": "good",
	  "url":"http://localhost.com",
	  "author": {
	    "firstname": "ketan",
	    "lastname": "rathod",
	    "age": 10,
		"email":"required with email type"
	  }
	}

*/
