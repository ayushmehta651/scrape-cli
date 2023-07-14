package moneycontrol

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
var headers []string

func getResponse(url string) (*http.Response, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func moneycontrol(args []string) {

	fmt.Println("Searching for: " + strings.Join(args, " "))

	url := "https://www.moneycontrol.com/stocks/marketstats/nsegainer/index.php"

	res, err := getResponse(url)
	if err != nil {
		log.Fatal(err)
	}

	moneycontrolScrape(res)

	createTable()
}

func moneycontrolScrape(res *http.Response) {
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.bsr_table").Find("table").Each(func(i int, item *goquery.Selection) {
		if i == 0 {
			item.Find("th").Each(func(j int, sele *goquery.Selection) {
				if sele != nil && j <= 5 {
					headers = append(headers, sele.Text())
				}
			})
			item.Find("tbody").Find("tr").Each(func(i int, sel *goquery.Selection) {
				var rows []string
				sel.Find("td").Each(func(j int, se *goquery.Selection) {
					if j == 0 {
						rows = append(rows, se.Find("h3").Text())
					} else {
						if j <= 5 {
							rows = append(rows, se.Text())
						}
					}
				})
				if len(rows) > 2 {
					indexString := helpers.Purple(strconv.Itoa(indexNumber))
					appendTableData(rows, indexString)
					indexNumber++
				}
			})
		}
	})
}

func appendTableData(values []string, indexNumber string) {

	cell := []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: indexNumber},
		{Align: simpletable.AlignLeft, Text: values[0]},
		{Align: simpletable.AlignLeft, Text: values[1]},
		{Align: simpletable.AlignLeft, Text: values[2]},
		{Align: simpletable.AlignLeft, Text: values[3]},
		{Align: simpletable.AlignLeft, Text: values[4]},
		{Align: simpletable.AlignLeft, Text: values[5]},
	}
	cells = append(cells, cell)
}

func createTable() {
	table := simpletable.New()

	// Set the headers
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "#"},
			{Align: simpletable.AlignLeft, Text: "Company Name"},
			{Align: simpletable.AlignLeft, Text: "High"},
			{Align: simpletable.AlignLeft, Text: "Low"},
			{Align: simpletable.AlignLeft, Text: "Last Price"},
			{Align: simpletable.AlignLeft, Text: "Prev Close"},
			{Align: simpletable.AlignLeft, Text: "Change"},
		},
	}

	// Table Body
	table.Body = &simpletable.Body{Cells: cells}

	// Set the style
	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())

}
