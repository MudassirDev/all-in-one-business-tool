package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/MudassirDev/all-in-one-business-tool/models"
	"github.com/google/uuid"
)

func CreateIndexFileHandler(apiCfg *models.APICfg) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uidStr, ok := r.Context().Value(apiCfg.AuthCookieName).(string)
		var user interface{}

		if ok {
			uid, err := uuid.Parse(uidStr)
			if err != nil {
				log.Printf("Invalid UUID: %v", err)
				http.Error(w, "Invalid user ID", http.StatusUnauthorized)
				return
			}

			dbUser, err := apiCfg.DB.GetUserByID(r.Context(), uid)
			if err != nil {
				log.Printf("Error fetching user: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			user = models.UserStruct{
				Username:  dbUser.Username,
				Email:     dbUser.Email,
				ID:        dbUser.ID,
				CreatedAt: dbUser.CreatedAt,
				UpdatedAt: dbUser.UpdatedAt,
				FirstName: dbUser.FirstName,
				LastName:  dbUser.LastName,
			}
		}

		tmpl, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"User": user,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}
