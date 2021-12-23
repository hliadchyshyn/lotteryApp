package main

import (
	"net/http"
	"testing"
)

func TestGetTicket(t *testing.T) {
	msg, _, httpStatus := getTicket("email@email.com", 1, false)

	if httpStatus != http.StatusOK {
		t.Errorf("Result was incorrect, got: %s, want: %s, statusCode: %d", msg, "", httpStatus)
	}

	msg, _, httpStatus = getTicket("email", 1, false)

	if httpStatus != http.StatusBadRequest {
		t.Errorf("Result was incorrect, got: %s, want: %s, statusCode: %d", msg, "Email is not valid", httpStatus)
	}

	msg, _, httpStatus = getTicket("email@email.com", 1, true)

	if httpStatus != http.StatusForbidden {
		t.Errorf("Result was incorrect, got: %s, want: %s, statusCode: %d", msg, "You've got ticket already", httpStatus)
	}

	msg, _, httpStatus = getTicket("email@email.com", 0, false)

	if httpStatus != http.StatusGone {
		t.Errorf("Result was incorrect, got: %s, want: %s, statusCode: %d", msg, "Sorry, but we are out of tickets", httpStatus)
	}
}
