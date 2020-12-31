package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
	"os"
	"github.com/Jeffail/gabs"
)

func main() {
    response, err := http.Get("https://api-in-dev.shortlyst.com/shopalyst-service/v1/products")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
	}
	
	jsonParsed, err := gabs.ParseJSON([]byte(data))
	if err != nil {
		panic(err)
	}

	fmt.Println()
	// To get all the required titles
	fmt.Println("Title of all items")
	fmt.Println()
	for _, child := range jsonParsed.Search("productList").Children() {
		fmt.Println(child.Path("title").Data())
	}

	fmt.Println()
	// To get all the required merchants
	fmt.Println("Merchants of all items")
	fmt.Println()
	for _, child := range jsonParsed.Search("productList").Children() {
		fmt.Println(child.Path("merchant").Data())
	}

	fmt.Println()
	// To get all the salePrice values 
	fmt.Println("SalePrice values of all items")
	fmt.Println()
	for _, child := range jsonParsed.Search("productList").Children() {
		fmt.Println(child.Path("salePrice").Data())
	}

	fmt.Println()
	// To get the result of only first list
	children := jsonParsed.Search("productList").Children()

	fmt.Println("title of first item: ", children[0].Path("title").Data())
	fmt.Println("merchant of first item: ", children[0].Path("merchant").Data())
	fmt.Println("salePrice of first item: ", children[0].Path("salePrice").Data())
	


}