package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	u "github.com/api-projet/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
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
	First_name    string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
	Birthdate  time.Time `json:"birthdate"`
	Accesslevel uint `json:"access_level"`
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	//Email must be unique
	temp := &Account{}

	//check for errors and duplicate emails
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "erreur de connection"), false
	}
	if temp.Email != "" {
		return u.Message(false, "addresse mail utilisé par un autre utilisateur"), false
	}

	for _, r := range account.Last_name {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return u.Message(false, "le lastname ne contient pas de bonne valeur"), false
		}
	}

	for _, r := range account.First_name {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return u.Message(false, "le firstname ne contient pas de bonne valeur"), false
		}
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !re.MatchString(account.Email) {
		return u.Message(false, "l'email ne contient pas de bonne valeur"), false
	}

	return u.Message(false, "validé"), true
}

func (account *Account) Create() (map[string]interface{}) {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	account.UUID = uuid.NewV4().String()

	GetDB().Create(account)

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: uuid.NewV4().String()}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Le compte a était crée")
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
	db.Where("uuid = ?", uuid).Find(&checkAccount)
	db.Delete(&checkAccount)
	response := u.Message(true, "le compte a été supprimé")
	return response
}

// PutUserHandler is updating a user from the given uuid param.
func PutUserHandler(params string,json *Account) (map[string]interface{}) {
	temp := &Account{}
	err := GetDB().Table("accounts").Where("UUID = ?", params).First(temp).Error


	//si le sujet de vote n'existe pas
	if err != nil{
		return u.Message(false, "Il n'y a aucun user avec cette uuid")
	}else{

		if(json.Email != ""){
			temp.Email = json.Email
		}

		if(json.Password != ""){
			temp.Password = json.Password
		}

		if(json.First_name != ""){
			temp.First_name = json.First_name
		}

		if(json.Last_name != ""){
			temp.Last_name = json.Last_name
		}


		temp.UpdatedAt = time.Now()
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(temp.Password), bcrypt.DefaultCost)
		temp.Password = string(hashedPassword)

		GetDB().Model(&temp).Update(temp)
	}

	response := u.Message(true, "L'user a été édité")
	response["user"] = temp
	return response
}



func Login(email, password string) (map[string]interface{}) {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email non trouvé")
		}
		return u.Message(false, "erreur de connection")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "login invalide")
	}


	//Create JWT token
	id := uuid.NewV4().String()
	tk := &Token{UserId: id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "connecté")
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


