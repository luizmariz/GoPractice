package core

import (
	"encoding/json"
	"log"
	"net/http"
)

// TODO: Refact with a better MVC pattern
type userData struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type mandrakeData struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	UserID      int    `json:"userId"`
}

func (c *Core) handleMandrakeSearch(w http.ResponseWriter, r *http.Request) {
	log.Println("start search")
}

func (c *Core) handleSearchesByUser(w http.ResponseWriter, r *http.Request) {
	log.Println("get user searches")
}

func (c *Core) handleCsvDownload(w http.ResponseWriter, r *http.Request) {
	// TODO: finish it after integration
}

func (c *Core) handleUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body map[string]interface{}

	err := decoder.Decode(&body)
	if err != nil {
		log.Println(err)
	}

	email, ok := body["email"]
	if !ok {
		http.Error(w, "You should have a field email", http.StatusInternalServerError)
		return
	} else {
		_, ok := email.(string)
		if !ok {
			http.Error(w, "Your email should be a string", http.StatusInternalServerError)
			return
		}
	}

	user := &userData{Email: email.(string)}

	err = c.Database.QueryRow("SELECT id FROM users WHERE email=?", email).Scan(&user.ID)
	if err != nil {
		log.Println("User not in database, creating...")

		res, err := c.Database.Exec("INSERT INTO users(email) VALUES (?)", email)
		if err != nil {
			log.Println(err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			log.Println(err)
		}

		user.ID = int(id)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		panic(err)
	}
}
