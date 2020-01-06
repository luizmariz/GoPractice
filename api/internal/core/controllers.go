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
	CsvPath     string `json:"csvPath"`
}

func decodeJSONResponse(r *http.Request) map[string]interface{} {
	decoder := json.NewDecoder(r.Body)
	var body map[string]interface{}

	err := decoder.Decode(&body)
	if err != nil {
		log.Println(err)
	}

	return body
}

func (c *Core) handleMandrakeSearch(w http.ResponseWriter, r *http.Request) {
	body := decodeJSONResponse(r)

	userID, ok := body["user"]
	if !ok {
		http.Error(w, "You should have a user field", http.StatusInternalServerError)
		return
	}

	searchURL, ok := body["searchUrl"]
	if !ok {
		http.Error(w, "You should send a valid url and search filters", http.StatusInternalServerError)
		return
	}

	// TODO: regex url validation

	description := body["description"] // value is nullable on db

	if _, err := c.Database.Exec(`
		INSERT INTO mandrakes(description, user_id, search_url)
		VALUES (?, ?, ?)`,
		description,
		userID,
		searchURL,
	); err != nil {
		http.Error(w, "An error ocurried while fetching our databases. Please contact an admin", http.StatusInternalServerError)
		return
	}

}

func (c *Core) handleSearchesByUser(w http.ResponseWriter, r *http.Request) {
	body := decodeJSONResponse(r)

	userID, ok := body["user"]
	if !ok {
		http.Error(w, "You should have a user field", http.StatusInternalServerError)
		return
	}

	rows, err := c.Database.Query(`
		SELECT
			m.id,
			COALESCE(description, '') AS description,
			status,
			COALESCE(csv_path, '') AS csv_path
		FROM mandrakes AS m
		JOIN users AS u ON m.user_id = u.id
		WHERE u.id = ?`,
		userID,
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "An error ocurried while fetching our databases. Please contact an admin", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var response []mandrakeData
	var mandrake *mandrakeData

	for rows.Next() {
		mandrake = new(mandrakeData)
		err := rows.Scan(&mandrake.ID, &mandrake.Description, &mandrake.Status, &mandrake.CsvPath)
		if err != nil {
			log.Println(err)
			http.Error(w, "An error ocurried. Please contact an admin", http.StatusInternalServerError)
			return
		}
		response = append(response, *mandrake)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "An error ocurried while parsing your data. Please contact admin", http.StatusInternalServerError)
		return
	}
}

func (c *Core) handleCsvDownload(w http.ResponseWriter, r *http.Request) {
	// TODO: finish it after integration
}

func (c *Core) handleUser(w http.ResponseWriter, r *http.Request) {
	body := decodeJSONResponse(r)

	email, ok := body["email"]
	if !ok {
		http.Error(w, "You should have a field email", http.StatusInternalServerError)
		return
	}

	emailValue, ok := email.(string)
	if !ok {
		http.Error(w, "Your email should be a string", http.StatusInternalServerError)
		return
	}

	user := &userData{Email: emailValue}

	if err := c.Database.QueryRow("SELECT id FROM users WHERE email=?", emailValue).Scan(&user.ID); err != nil {
		// if user not in db just register it

		res, err := c.Database.Exec("INSERT INTO users(email) VALUES (?)", emailValue)
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
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "An error ocurried while parsing your data. Please contact admin", http.StatusInternalServerError)
		return
	}
}
