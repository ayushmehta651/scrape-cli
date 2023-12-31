package ebay

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

var cells [][]*simpletable.Cell
var indexNumber int = 1

func getResponse(url string) (*http.Response, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func ebay(args []string, pages int) {

	fmt.Println("Searching for: " + strings.Join(args, " "))
	queryKeyword := strings.Join(args, "+")

	currentPage := 1
	totalPages := pages

	for currentPage <= totalPages {

		url := "https://www.ebay.com/sch/i.html?_from=R40&_nkw=" + queryKeyword + "&_pgn=" + strconv.Itoa(totalPages)

		res, err := getResponse(url)
		if err != nil {
			log.Fatal(err)
		}

		ebayScrape(res)

		fmt.Printf("Scraping page %d of %d\n", currentPage, totalPages)
		currentPage++
	}

	createTable()

}

func ebayScrape(res *http.Response) {

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("ul.clearfix").Find("li.s-item__pl-on-bottom").Each(func(i int, item *goquery.Selection) {
		title := item.Find("div.s-item__title").Find("span").Text()
		price := item.Find("span.s-item__price").Text()

		// Truncate the title if it is too long
		if len(title) > 80 {
			title = title[:80] + "..."
		}

		price = helpers.Green(price)
		indexString := helpers.Purple(strconv.Itoa(indexNumber))

		appendTableData(title, price, indexString)
		indexNumber++

	})

}

func appendTableData(title, price, indexNumber string) {

	cell := []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: indexNumber},
		{Align: simpletable.AlignLeft, Text: title},
		{Align: simpletable.AlignLeft, Text: price},
	}
	cells = append(cells, cell)

}

func createTable() {

	table := simpletable.New()

	// Set the headers
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Product Name"},
			{Align: simpletable.AlignCenter, Text: "Price"},
		},
	}

	// Table Body
	table.Body = &simpletable.Body{Cells: cells}

	// Set the style
	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())

}
