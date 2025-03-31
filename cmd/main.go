package main

import (
	"fmt"
	"sync"
	// "io"
	"net/http"
	// "os"
	"time"
)


type Emploue struct {
	Name string
	Age int
}


func main() {

	vas := &Emploue{"VAS",38}
	var wg sync.WaitGroup
	t := time.Now()

	
	go func ()  {
		fmt.Printf("Can emploue %s drink bear ? - %v.\t It's age = %v\n", vas.Name, canDrink(&vas.Age), *getAge(*vas))
		
	} ()
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go requestGoogle(i, &wg)
		
	}
	wg.Wait()

	time.Sleep(time.Second * 4)
	tt := time.Since(t)
	fmt.Printf("Total time %v\n", tt)
}


func canDrink(age *int) bool {
	return *age >=18
}

func getAge(e Emploue) *int{
	return &e.Age
}

func requestGoogle(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get("https://google.com") 
	if err != nil { 
			fmt.Println(err) 
			return
	} 
	defer resp.Body.Close()
	sC := resp.StatusCode 
	fmt.Printf("Try - %d status - %v \n", num, sC)
//    io.Copy(os.Stdout, resp.Body)
}
