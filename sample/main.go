package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/kyokomi/go-scrapbox"
)

func main() {
	token := flag.String("t", "aaaaaaaaa", "scrapbox connect.sid")
	projectName := flag.String("p", "kyokomi", "scrapbox project name")
	flag.Parse()

	offset := uint(0)
	limit := uint(5)

	client := scrapbox.NewClient(*token)

	ctx := context.Background()
	res, err := client.PageService.List(ctx, *projectName, offset, limit)
	if err != nil {
		log.Fatal(err)
	}

	for _, page := range res.Pages {
		fmt.Println(page.Title, page.Descriptions[0])
	}
}
