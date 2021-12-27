package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/leyiqiang/home-library-server/models"
	"github.com/leyiqiang/home-library-server/utils"
	"net/http"
)

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	hashedPwd, err := utils.GetHash([]byte(user.Password))

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	user.Password = hashedPwd
	
	validate = validator.New()
	err = validate.Struct(user)
	if err != nil {
		utils.ErrorJSON(w, err)
		return

	}
	err = c.Repo.Register(user)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, "", "user")

}
