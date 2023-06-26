package main

import (
	"log"
	"net/http"
)

func main() {
	/*
		// create file server handler
		fs := http.FileServer(http.Dir("/bundles"))

		// handle `/bundles` route
		http.Handle("/bundles/", http.StripPrefix("/bundles/", fs))
	*/

	// This works too, but "/static2/" fragment remains and need to be striped manually
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/gzip")
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	printPrimeNumber(5,19)
	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":9000", nil))

}

func printPrimeNumbers(num1, num2 int){
   if num1<2 || num2<2{
      fmt.Println("Numbers must be greater than 2 for this to work.")
      return
   }
   for num1 <= num2 {
      isPrime := true
      for i:=2; i<=int(math.Sqrt(float64(num1))); i++{
         if num1 % i == 0{
            isPrime = false
            break
         }
      }
      if isPrime {
         fmt.Printf("Found Prime number: %d ", num1)
      }
      num1++
   }
   fmt.Println()
}
