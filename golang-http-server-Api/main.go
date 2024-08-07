package main

import (
	"fmt"
	"log"
	"net/http"
)


func calcuatemyage(){



	var age string



	fmt.Print("Please enter your age:") 
	fmt.Scanln(&age)
	
	if age, err := strconv.ParseUint(age); err{
		fmt.Println(err)
		fmt.Printf("%d" age)
		
	} else
	{
		
		fmt.printf(err)
	
		
	}
}



func evenandoddscombine(u string)(int, int){

evens,odds:=0,0
for _,c := range u{

	if c %2==0{
		evens++
	}else{

		odds++
	}
	
}
return evens,odds
}

func sumtwonumbers(s string,evens int,odds int) (int, int){

	
	for _,c := range s{
	
		if c %2==0{
			evens++
		}else{
	
			odds++
		}
		
	}
	return evens,odds
	}
	

//coding two functions together

func raining() bool {

	fmt.printf("Check if its raining now")
	return true
}
func snowing() bool{

	fmt.printf("check if its snowing now")
	return true
}


// lets do another functional procedures

func dosomething() (int, bool){

	return int , bool


}
func checkparity(num int) string{

	if num % 2 ==0{

		return "even"
	}
	return "odd"
}

//now lets deal with slices which can grow in sizes and lenth unlike arrays

func delete(oring []int,index int ) ([]int ,error){
if index<0 || index>=len(oring){

	return nil,error.New("Eooer out of bound")
}
orig = append(orig[:index], orig[index+1:]...)
    return orig, nil

}





func httphandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, my name is ayeni clement , I love you all %s!", r.URL.Path[1:])
}


func main() {
	http.HandleFunc("/", httphandler)

	log.Fatal(http.ListenAndServe(":2020", nil))
	if raining() || snowing(){
		fmt.printf("please stay in doors")
		
	}
	calcuatemyage()
	//now call the function countOddEven
	o, e := countOddEven("12345") 
	fmt.Println(o,e) // 3 2

	//If there are parts of the result that you want to ignore, use the blank identifier (_):
	_, e := countOddEven("12345") 
	fmt.Println(e) // 3 2
	u,k := sumtwonumbers("12345",0,0)
	fmt.Println(u,k) // 3 2

	//test the delete function
	t := []int {1,2,3,4,5}
	if t,err=delete(t,2); err{
			fmt.Println(err)
		} else {

		fmt.Println(t)
		}
		// [1 2 4 5]
	}

