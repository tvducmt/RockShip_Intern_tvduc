package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"

	"github.com/go-chi/chi"
	"github.com/robfig/cron"
)

type match struct {
	Venue             string `json:"venue"`
	Location          string `json:"location"`
	Status            string `json:"status"`
	Time              string `json:"time"`
	Fifaid            string `json:"fifa"`
	Datetime          string `json:"datetime"`
	Lasteventupdateat string `json:"lasteventupdateat"`
	Hometeam          string `json:"hometeam"`
	Awayteam          string `json:"awayteam"`
	Winner            string `json:"winner"`
}

func sendMail(matchcurrent string) {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "1510819@hcmut.edu.vn", "Password", "smtp.gmail.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"tvducmt@gmail.com"}
	msg := []byte("This is the email body.\r\n")
	err := smtp.SendMail("smtp.gmail.com:25", auth, "1510819@hcmut.edu.vn", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	r := chi.NewRouter()

	r.Get("/Matchs", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:3000/match")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//responseString := string(responseData)

		Matchcurrent := match{}
		json.Unmarshal(responseData, &Matchcurrent)

		c := cron.New()
		c.AddFunc("1 * * * * *", func() {
			if Matchcurrent.Winner != "" {
				fmt.Println(Matchcurrent)
				responseString := string(responseData)
				sendMail(responseString)
			}

		})
		fmt.Println("Start...")
		c.Run()

	})

	http.ListenAndServe(":8888", r)
}
