package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/leyiqiang/home-library-server/config"
	"github.com/leyiqiang/home-library-server/models"
	"github.com/leyiqiang/home-library-server/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var cfg config.Config

func init() {
	cfg.Read()
}

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
	user.IsAdmin = false

	validate = validator.New()
	err = validate.Struct(user)
	if err != nil {
		utils.ErrorJSON(w, err)
		return

	}
	err = c.Repo.AddUser(user)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, "", "user")

}

func (c *Controller) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		utils.ErrorJSON(w, errors.New("unauthorized"), 401)
	}

	found, err := c.Repo.FindUserByUsername(user.Username)

	//hashed, err := utils.GetHash([]byte(user.Password))
	err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(user.Password))
	if err != nil {
		utils.ErrorJSON(w, errors.New("unauthorized"), 401)
	}

	var claims jwt.StandardClaims
	claims.Subject = fmt.Sprint(user.ID)
	claims.IssuedAt = time.Now().Unix()
	claims.NotBefore = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
	claims.Issuer = "home-library"
	claims.Audience = "home-library"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfg.Auth.Secret))

	if err != nil {
		utils.ErrorJSON(w, err)
	}
	utils.WriteJSON(w, http.StatusOK, tokenString, "response")

}
