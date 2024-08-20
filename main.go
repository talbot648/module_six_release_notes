package main

import (
	"encoding/csv"
	"errors"
	"os"
)

type Ticket struct {
	TicketKey   string
	Summary     string
	Description string
}

func getTicketInformation() (string, string, string, error) {
	//opening csv file
	file, err := os.Open("ticket.csv")
	if err != nil {
		return "", "", "", errors.New("error when opening ticket page") //returns no value for the strings and returns specific error for opening page error
	}
	defer file.Close()

	//reading csv file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return "", "", "", errors.New("error reading csv file") //returns specific error for not reading csv file
	}

	for _, record := range records[1:] {
		ticket := Ticket{
			TicketKey:   record[0],
			Summary:     record[1],
			Description: record[2],
		}
		return ticket.TicketKey, ticket.Summary, ticket.Description, nil
	}
	return "all tickets have been processed", "", "", nil
}
