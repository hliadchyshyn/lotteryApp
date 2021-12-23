package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	port := getPort()
	ticketsLeft = getTicketsAmount()
	http.HandleFunc("/ticket", getTicketHandler)
	//  Start HTTP
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal("Failed starting http server: ", err)
	}

}

func getPort() string {

	port := goDotEnvVariable("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	return port

}

func getTicketsAmount() int {
	s := goDotEnvVariable("num_of_tickets")
	ticketsLeft, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Failed converting tickets left env variable: ", err)
	}
	return ticketsLeft
}

type User struct {
	Email string
}

//Global variables to simplify the solution for the demonstration purpose
//in real project should be realised in storage
var (
	users       = make(map[string]int)
	ticketsLeft int
)

func getTicketHandler(w http.ResponseWriter, r *http.Request) {
	var u User

	// Validate and Parse JSON APi body
	err := decodeJSONBody(w, r, &u)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	_, found := users[u.Email]
	msg, ok, httpStatus := getTicket(u.Email, ticketsLeft, found)
	if !ok {
		http.Error(w, msg, httpStatus)
		return
	}

	n, err := fmt.Fprintf(w, "Email: %+v", u.Email)
	if err != nil {
		// Printing the number of bytes written
		fmt.Print(n, " bytes written.\n")

		// Printing if any error encountered
		fmt.Print(err)
	}
}

func getTicket(email string, ticketsLeft int, found bool) (msg string, ok bool, httpStatus int) {
	//Check if email valid
	if !valid(email) {
		return "Email is not valid", false, http.StatusBadRequest
	}

	//Check if user already get ticket
	if found {
		return "You've got ticket already", false, http.StatusForbidden
	}

	if ticketsLeft < 1 {
		return "Sorry, but we are out of tickets", false, http.StatusGone
	}

	// Add user as get ticket
	users[email] = ticketsLeft

	// Decrement ticket amount
	ticketsLeft--

	return "", true, http.StatusOK
}
