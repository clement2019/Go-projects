package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
)

//lets try to do another data collections for students and subjects using structs

type student struct {
	name    string
	email   string
	class   string
	subject subject
	dob     dob
}

type subject struct {
	subj1 string
	subj2 string
	subj3 string
}

// Now for us to count the even and odds in a string

func countpassandfail(p string) int {
	passs := 0
	for _, c := range p {
		if c%2 == 3 {
			passs += 1
		}

	}
	return passs
}

func findOddsandEven(u string, evens int, odds int) (int, int) {

	for _, c := range u {

		if c%2 == 0 {
			evens++
		} else {

			odds++
		}

	}
	return evens, odds
}

// crteating struts and map for our =data type
// we can create struct of dop and strcut of people ehile creating a mpa of key int aand type people

type dob struct {
	day   int
	month int
	year  int
}

type people struct {
	name  string
	email string
	dob   dob
}

// Now lets create the map for our students data collections

var records map[int]student

// create the map
var memembers map[int]people

// Now to calculate the squareroot of an element and make sure it doesnt have zero input
func squareroot(s float64) (float64, error) {
	if s < 0 {

		return 0, errors.New("undefined the value of s cannot be less than zero")
	} else {
		return math.Sqrt(s), nil
	}

}

func main() {

	var john = "hello and welcome Ayenco and dami"
	fmt.Println(john)
	if result, err := squareroot(-2); err == nil {
		fmt.Println(result)

	} else {
		fmt.Println(err)

	}

	records = make(map[int]student)

	records[1] = student{
		name:  "seyi",
		email: "seyi@yahoo.com",
		class: "ss2",
		subject: subject{
			subj1: "math",
			subj2: "english",
			subj3: "economics",
		},
		dob: dob{
			day:   3,
			month: 5,
			year:  2012,
		},
	}
	records[2] = student{
		name:  "femi johnson",
		email: "femi@yahoo.com",
		class: "ss2",
		subject: subject{
			subj1: "math",
			subj2: "english",
			subj3: "government",
		},
		dob: dob{
			day:   13,
			month: 6,
			year:  2012,
		},
	}
	records[3] = student{
		name:  "bunmi oluwa",
		email: "bunmi@yahoo.com",
		class: "ss2",
		subject: subject{
			subj1: "math",
			subj2: "english",
			subj3: "arabic",
		},
		dob: dob{
			day:   31,
			month: 4,
			year:  2011,
		},
	}
	memembers = make(map[int]people)
	memembers[1] = people{
		name:  "Bolajhi",
		email: "bolaji@yahoo.com",

		dob: dob{
			day:   7,
			month: 9,
			year:  2003,
		},
	}
	memembers[2] = people{
		name:  "kunle",
		email: "kunle@yahoo.com",

		dob: dob{
			day:   17,
			month: 2,
			year:  2001,
		},
	}
	memembers[3] = people{
		name:  "john",
		email: "john@yahoo.com",

		dob: dob{
			day:   2,
			month: 4,
			year:  1993,
		},
	}
	memembers[4] = people{
		name:  "clement",
		email: "clement@yahoo.com",

		dob: dob{
			day:   2,
			month: 2,
			year:  2012,
		},
	}

	// collate the key value and data contentst of the students records
	for t, q := range records {

		fmt.Println(t, q.name, q.email, q.class, q.subject, q.dob)
	}
	johnJSON, err := json.Marshal(records)
	if err == nil {
		fmt.Println(johnJSON)
	} else {
		fmt.Println(err)
	}

	//to print all the key value and data contents of memembers and store in k and v
	for k, v := range memembers {
		fmt.Println(k, v.name, v.email, v.dob)

	}
	//call trhe evenandddsfunctions and you dont need the key values
	w := countpassandfail("4666789")
	fmt.Println(w)

	//call the function sumoftwonumbers here
	k, v := findOddsandEven("45678912", 0, 0)
	fmt.Println(k, v)

}

// Hello returns a greeting for the named person.
