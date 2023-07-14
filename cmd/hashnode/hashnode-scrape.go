package hashnode

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ayushmehta651/scrape-cli/helpers"

	"github.com/PuerkitoBio/goquery"
	"github.com/alexeyco/simpletable"
)

var indexNumber = 1
var cells [][]*simpletable.Cell

func getResponse(url string) (*http.Response, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func hashnode(args []string, section string) {

	fmt.Println("Searching for: " + strings.Join(args, " "))

	url := "https://hashnode.com/search/" + args[1] + "?q=" + args[0]

	res, err := getResponse(url)
	if err != nil {
		log.Fatal(err)
	}

	hashnodeScrape(res)

	createTable()

}

func hashnodeScrape(res *http.Response) {
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.flex-1").Each(func(i int, item *goquery.Selection) {
		title := item.Find("h3").Text()
		popularity := item.Find("span").Text()

		title = helpers.Green(title)
		popularity = helpers.Red(popularity)
		indexString := helpers.Purple(strconv.Itoa(indexNumber))

		appendTableData(title, popularity, indexString)
		indexNumber++
	})
}

func appendTableData(title, popularity, indexNumber string) {

	cell := []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: indexNumber},
		{Align: simpletable.AlignLeft, Text: title},
		{Align: simpletable.AlignLeft, Text: popularity},
	}
	cells = append(cells, cell)
}

func createTable() {

	table := simpletable.New()

	// Set the headers
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "Popularity"},
		},
	}

	// Table Body
	table.Body = &simpletable.Body{Cells: cells}

	// Set the style
	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())

}
