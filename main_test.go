package main

import (
	"testing"
)

func TestGetTicketSummaryInformation(t *testing.T) {
	//arrange
	got, _, _ := getTicketInformation()
	want := "Implement User Authentication"

	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}

}

/*
func testGetTicketDescriptionInformation(t *testing.T) {
	//arrange
	ticketDescription := getTicketInformation()

	descriptionWant := "As a security team member I would like to implement a user suthentication to test application to enhance security"

}
*/
