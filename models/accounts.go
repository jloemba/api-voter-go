package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	u "github.com/api-projet/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	uuid "github.com/satori/go.uuid"
	"time"

)

/*
JWT claims struct
*/
type Token struct {
	UserId  string `json:"uuid" gorm:"primary_key"`
	//UUID      string `json:"uuid" gorm:"primary_key"`
	jwt.StandardClaims
}

//a struct to rep user account
type Account struct {
	gorm.Model
	UUID      string `json:"uuid" gorm:"primary_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
	Birthdate  time.Time `json:"birthdate"`
	Accesslevel uint `json:"access_level"`
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &Account{}

	//check for errors and duplicate emails
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string]interface{}) {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	account.UUID = uuid.NewV4().String()

	GetDB().Create(account)
	/*if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}*/

	//Create new JWT token for the newly registered account
	//u.UUID = uuid.NewV4().String()
	tk := &Token{UserId: uuid.NewV4().String()}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}



// DeleteUserHandler is deleting a user from the given uuid param.
func DeleteUserHandler(uuid string) (map[string]interface{}) {
	account := &Account{}
	account.UUID = uuid
	err := GetDB().Table("accounts").Where("uuid = ?", uuid).First(account).Error
	if err != nil {
		response := u.Message(false, "Le compte n'existe pas")
		return response
	}
	var checkAccount Account
	db.Where("name = ?", uuid).Find(&checkAccount)
	db.Delete(&checkAccount)
	response := u.Message(true, "le compte a été supprimé")
	return response
}



func Login(email, password string) (map[string]interface{}) {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	id := uuid.NewV4().String()
	tk := &Token{UserId: id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}


