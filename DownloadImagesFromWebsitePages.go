package main

import (
	"github.com/gocolly/colly"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var num1 = 0
var links  = ""

func main(){

	collector := colly.NewCollector()
 	collector.OnHTML("img[src]",func(e *colly.HTMLElement){
 		 links = e.Attr("src")
		link := strings.Split(links,"\n")

		for index:=0;index<len(link) ;index++  {
			response,_ := http.Get("first Link"+link[index])
			defer response.Body.Close()
			num1 +=1

			outfile ,_ := os.Create("./data/"+strconv.Itoa(num1)+".jpg")
			defer  outfile.Close()
			io.Copy(outfile,response.Body)
		}

	})
	collector.Visit("your web site")

}

