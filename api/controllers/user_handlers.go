package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MudassirDev/all-in-one-business-tool/api/auth"
	Json "github.com/MudassirDev/all-in-one-business-tool/api/json"
	"github.com/MudassirDev/all-in-one-business-tool/internal/database"
	"github.com/MudassirDev/all-in-one-business-tool/models"
)

func CreateUserRegisterHandler(apiCfg *models.APICfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type RequestParams struct {
			Username  string `json:"username"`
			Email     string `json:"email"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Password  string `json:"password"`
		}

		var params RequestParams

		defer r.Body.Close()

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			Json.RespondWithError(w, http.StatusBadRequest, "invalid parameters", err)
			return
		}

		pass, err := auth.HashPassword(params.Password)
		if err != nil {
			Json.RespondWithError(w, http.StatusInternalServerError, "failed to encode password", err)
			return
		}

		dbUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			Username:     params.Username,
			Email:        params.Email,
			FirstName:    params.FirstName,
			LastName:     params.LastName,
			PasswordHash: pass,
		})

		if err != nil {
			Json.RespondWithError(w, http.StatusInternalServerError, "failed to create user", err)
		}

		Json.RespondWithJson(w, http.StatusCreated, models.UserStruct{
			ID:        dbUser.ID,
			Email:     dbUser.Email,
			Username:  dbUser.Username,
			FirstName: dbUser.FirstName,
			LastName:  dbUser.LastName,
			CreatedAt: dbUser.CreatedAt,
			UpdatedAt: dbUser.UpdatedAt,
		})
	}
}

func CreateUserLoginHandler(apiCfg *models.APICfg) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
