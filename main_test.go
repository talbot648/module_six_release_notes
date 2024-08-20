package main

import (
	"fmt"
	"testing"
)

func TestGetTicketSummaryInformation(t *testing.T) {
	//arrange
	_, got, _, err := getTicketInformation()
	fmt.Println(err)
	want := "Implement User Authentication"

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestGetTicketKeyInformation(t *testing.T) {
	//arrange
	got, _, _, _ := getTicketInformation()

	want := "ATT-1588"

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

func TestGetTicketDescriptionInformation(t *testing.T) {
	//arrange
	_, _, got, _ := getTicketInformation()

	want := "As a security team member I would like to implement a user suthentication to test application to enhance security"

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}
