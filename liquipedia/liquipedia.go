package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL of the Valorant Liquipedia tournament page you want to scrape
	url := "https://liquipedia.net/valorant/VCT/2023/Game_Changers/EMEA/Stage_3"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// Make an HTTP GET request to the URL
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Extract tournament information
	tournamentName := doc.Find(".infobox-header").First().Text()
	tournamentName = strings.TrimPrefix(tournamentName, "[e][h]")
	tournamentDates := doc.Find(".infobox-valorant infobox-cell-2:contains('Start Date')").Text()
	participatingTeams := []string{}

	// Extract participating teams (modify selector as needed)
	doc.Find(".teamcard-template-standard").Each(func(index int, teamCard *goquery.Selection) {
		teamName := teamCard.Find(".teamcard-template-text").Text()
		participatingTeams = append(participatingTeams, teamName)
	})

	// Extract sponsors (modify selector as needed)
	sponsors := doc.Find(".infobox-valorant .infobox-cell-2:contains('Sponsors')").Text()
	// Extract tournament format, prize pool, and any other information as needed
	tournamentFormat := doc.Find(".mw-headline b:contains('Format')").Parent().Next().Text()
	prizePool := doc.Find(".infobox-valorant .infobox-cell-2 b:contains('Prize Pool')").Next().Text()

	// Clean up and format the dates
	tournamentDates = strings.TrimSpace(tournamentDates)
	tournamentDates = strings.ReplaceAll(tournamentDates, " – ", " - ")

	// Print tournament information
	fmt.Printf("Tournament Name: %s\n", tournamentName)
	fmt.Printf("Tournament Dates: %s\n", tournamentDates)
	fmt.Println("Participating Teams:")
	for i, team := range participatingTeams {
		fmt.Printf("%d. %s\n", i+1, team)
	}
	fmt.Printf("Sponsors: %s\n", sponsors)
	fmt.Printf("Tournament Format: %s\n", tournamentFormat)
	fmt.Printf("Prize Pool: %s\n", prizePool)
}
