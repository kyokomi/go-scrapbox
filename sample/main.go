package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/kyokomi/go-scrapbox"
)

func main() {
	log.SetFlags(log.Llongfile)
	token := flag.String("t", "aaaaaaaaa", "scrapbox connect.sid")
	projectName := flag.String("p", "kyokomi", "scrapbox project name")
	flag.Parse()

	offset := uint(0)
	limit := uint(5)

	client := scrapbox.NewClient(*token)

	ctx := context.Background()
	res, err := client.Page.List(ctx, *projectName, offset, limit)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	for _, page := range res.Pages {
		fmt.Println(page.Title, page.Descriptions[0])

		detail, err := client.Page.Get(ctx, *projectName, page.Title)
		if err != nil {
			log.Fatalf("%+v", err)
		}

		fmt.Println("# links")
		fmt.Println(detail.Links)
		fmt.Println("# link1hop")
		for _, linkHop := range detail.RelatedPages.Links1Hop {
			fmt.Println(linkHop.Title)
		}
		fmt.Println("# link2hop")
		for _, linkHop := range detail.RelatedPages.Links2Hop {
			fmt.Println(linkHop.Title)
		}

		found, iconURL, err := client.Page.IconURL(ctx, *projectName, page.Title)
		if err != nil {
			log.Fatalf("%+v", err)
		}

		if found {
			fmt.Println(iconURL.String())
		}

		text, err := client.Page.Text(ctx, *projectName, page.Title)
		if err != nil {
			log.Fatalf("%+v", err)
		}
		fmt.Println(text)

		fmt.Println()
	}
}
