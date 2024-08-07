package studentship

import(
	

)

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


