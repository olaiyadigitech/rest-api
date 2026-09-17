package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "modernc.org/sqlite"
)

type Portfolio struct {
	Name         string   `json:"name"`
	Title        string   `json:"title"`
	Experience   string   `json:"experience"`
	Skills       []string `json:"skills"`
	Technologies []string `json:"technologies"`
}

type Skill struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var portfolio = Portfolio{
	Name:       "Ajayi Oluwafemi Olaiya",
	Title:      "Network Engineer | NOC Engineer | Infrastructure & Automation Enthusiast | Go Backend Developer",
	Experience: "1 Year",
	Skills: []string{
		"Network Engineering",
		"NOC Operations",
		"Routing & Switching",
		"TCP/IP",
		"Fibre Optic Transmission",
		"Network Troubleshooting",
	},
	Technologies: []string{
		"Cisco",
		"Go",
		"Linux",
		"Git",
		"GitHub",
		"REST API",
	},
}

var db *sql.DB

func portfolioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(portfolio); err != nil {
		log.Printf("Error encoding portfolio: %v", err)
	}
}

func skillsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		rows, err := db.Query("SELECT id, name FROM skills ORDER BY id")
		if err != nil {
			http.Error(w, "Failed to retrieve skills", http.StatusInternalServerError)
			log.Printf("Database query error: %v", err)
			return
		}
		defer rows.Close()

		skills := []Skill{}

		for rows.Next() {
			var skill Skill

			if err := rows.Scan(&skill.ID, &skill.Name); err != nil {
				http.Error(w, "Failed to read skills", http.StatusInternalServerError)
				log.Printf("Database scan error: %v", err)
				return
			}

			skills = append(skills, skill)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, "Failed to read skills", http.StatusInternalServerError)
			log.Printf("Database rows error: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(skills); err != nil {
			log.Printf("Error encoding skills: %v", err)
		}

	case http.MethodPost:
		var newSkill Skill

		if err := json.NewDecoder(r.Body).Decode(&newSkill); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if newSkill.Name == "" {
			http.Error(w, "Skill name is required", http.StatusBadRequest)
			return
		}

		result, err := db.Exec(
			"INSERT INTO skills (name) VALUES (?)",
			newSkill.Name,
		)

		if err != nil {
			http.Error(w, "Failed to create skill", http.StatusInternalServerError)
			log.Printf("Database insert error: %v", err)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to get skill ID", http.StatusInternalServerError)
			log.Printf("LastInsertId error: %v", err)
			return
		}

		newSkill.ID = int(id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(newSkill); err != nil {
			log.Printf("Error encoding new skill: %v", err)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func skillByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/skills/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid skill ID", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodGet:
		var skill Skill

		err := db.QueryRow(
			"SELECT id, name FROM skills WHERE id = ?",
			id,
		).Scan(&skill.ID, &skill.Name)

		if err == sql.ErrNoRows {
			http.Error(w, "Skill not found", http.StatusNotFound)
			return
		}

		if err != nil {
			log.Printf("Database query error: %v", err)
			http.Error(w, "Failed to retrieve skill", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(skill)

	case http.MethodPut:
		var updatedSkill Skill

		if err := json.NewDecoder(r.Body).Decode(&updatedSkill); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if updatedSkill.Name == "" {
			http.Error(w, "Skill name is required", http.StatusBadRequest)
			return
		}

		result, err := db.Exec(
			"UPDATE skills SET name = ? WHERE id = ?",
			updatedSkill.Name,
			id,
		)

		if err != nil {
			log.Printf("Database update error: %v", err)
			http.Error(w, "Failed to update skill", http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("RowsAffected error: %v", err)
			http.Error(w, "Failed to update skill", http.StatusInternalServerError)
			return
		}

		if rowsAffected == 0 {
			http.Error(w, "Skill not found", http.StatusNotFound)
			return
		}

		updatedSkill.ID = id

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedSkill)

	case http.MethodPatch:
		var updates map[string]interface{}

		if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		name, ok := updates["name"].(string)

		if !ok || name == "" {
			http.Error(w, "Skill name is required", http.StatusBadRequest)
			return
		}

		result, err := db.Exec(
			"UPDATE skills SET name = ? WHERE id = ?",
			name,
			id,
		)

		if err != nil {
			log.Printf("Database patch error: %v", err)
			http.Error(w, "Failed to update skill", http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("RowsAffected error: %v", err)
			http.Error(w, "Failed to update skill", http.StatusInternalServerError)
			return
		}

		if rowsAffected == 0 {
			http.Error(w, "Skill not found", http.StatusNotFound)
			return
		}

		var skill Skill

		err = db.QueryRow(
			"SELECT id, name FROM skills WHERE id = ?",
			id,
		).Scan(&skill.ID, &skill.Name)

		if err != nil {
			http.Error(w, "Failed to retrieve updated skill", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(skill)

	case http.MethodDelete:
		result, err := db.Exec(
			"DELETE FROM skills WHERE id = ?",
			id,
		)

		if err != nil {
			log.Printf("Database delete error: %v", err)
			http.Error(w, "Failed to delete skill", http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("RowsAffected error: %v", err)
			http.Error(w, "Failed to delete skill", http.StatusInternalServerError)
			return
		}

		if rowsAffected == 0 {
			http.Error(w, "Skill not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func initDatabase() {
	var err error

	db, err = sql.Open("sqlite", "rest-api.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS skills (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		)
	`)

	if err != nil {
		log.Fatal("Failed to create skills table:", err)
	}

	log.Println("SQLite database connected")
}

func main() {
	initDatabase()
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/portfolio", portfolioHandler)
	mux.HandleFunc("/skills", skillsHandler)
	mux.HandleFunc("/skills/", skillByIDHandler)

	server := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	log.Println("REST API server listening on :8090")
	log.Println("GET    http://localhost:8090/portfolio")
	log.Println("GET    http://localhost:8090/skills")
	log.Println("POST   http://localhost:8090/skills")
	log.Println("GET    http://localhost:8090/skills/{id}")
	log.Println("PUT    http://localhost:8090/skills/{id}")
	log.Println("PATCH  http://localhost:8090/skills/{id}")
	log.Println("DELETE http://localhost:8090/skills/{id}")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
