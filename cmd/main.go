package main

import (
	"fmt"
	"sync"
	// "io"
	"net/http"
	// "os"
	"time"
)


// type Emploue struct {
// 	Name string
// 	Age int
// }
// func canDrink(age *int) bool {
// 	return *age >=18
// }

// func getAge(e Emploue) *int{
// 	return &e.Age
// }

func main() {

	
	
	// vas := &Emploue{"VAS",38}
	// go func ()  {
	// 	fmt.Printf("Can emploue %s drink bear ? - %v.\t It's age = %v\n", vas.Name, canDrink(&vas.Age), *getAge(*vas))
	// } ()

	var wg sync.WaitGroup
	t := time.Now()
	sites := generateSite()


	for site := range sites {
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				resp := <- statusSite(site.URL)
				fmt.Println(resp)
				wg.Done()			
			}()
		}
		
		go func() {
			wg.Wait()
			close(sites)
			
		}()
		

	}

	tt := time.Since(t)
	fmt.Printf("Total time %v\n", tt)
}

func statusSite(site string) <- chan string {
		ch := make(chan string, 1)
		resp, err := http.Get(site) 
		if err != nil { 
			fmt.Println(err) 
			ch <- err.Error()
			return ch
				
		} 
		defer resp.Body.Close()
		info := fmt.Sprintf("Try to - %s  status - %d.",site, resp.StatusCode) 
		ch <- info
		return ch
//    io.Copy(os.Stdout, resp.Body)
}

func generateSite() chan SiteInfo{
	ch := make(chan SiteInfo)
	sites := []SiteInfo{
		SiteInfo{"Distrowatch","https://distrowatch.com/"},
		SiteInfo{"Vanilla","https://vanillaos.org/"},
		SiteInfo{"AlmaLinux","https://almalinux.org/"},
		SiteInfo{"Nobara","https://nobaraproject.org/"},
		SiteInfo{"POP.os","https://system76.com/pop/"},
		// SiteInfo{"Fedora","https://fedoraproject.org/"},
		SiteInfo{"OpenSuse","https://www.opensuse.org/"},
		// SiteInfo{"Manjaro","https://manjaro.org/"},
		// SiteInfo{"Debian","https://www.debian.org/"},
		SiteInfo{"Google","https://google.com"}}
		go func ()  {
			for _, v := range sites {
				ch <- v
			}	
		}()
	return ch
}

type SiteInfo struct {
	Name string
	URL string
}
