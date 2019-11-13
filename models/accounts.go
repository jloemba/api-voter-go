package models

import (
	"fmt"
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
	db.Where("uuid = ?", uuid).Find(&checkAccount)
	db.Delete(&checkAccount)
	response := u.Message(true, "le compte a été supprimé")
	return response
}

// PutUserHandler is updating a user from the given uuid param.
func PutUserHandler(params string,json *Account) (map[string]interface{}) {
/*
	account.UUID = uuid
	var checkAccount Account
	checkAccount.UUID = uuid
	//err := db.Where("uuid = ?", uuid).Find(&checkAccount)
	err := GetDB().Table("accounts").Where("uuid = ?", uuid).First(account).Error


	fmt.Println(account)

	//si le sujet de vote n'existe pas
	if err == nil {
		return u.Message(false, "Il n'y a aucun user avec cette uuid")
	}

	if len(account.Email) > 0 {
		checkAccount.Email = account.Email
	}
	if !account.Birthdate.IsZero() {
		checkAccount.Birthdate = account.Birthdate
	}
	if len(account.Password) > 0 {
		checkAccount.Password = account.Password
	}
	if account.Accesslevel <= 0 {
		checkAccount.Accesslevel = account.Accesslevel
	}

	if resp, ok := checkAccount.Validate(); !ok {
		return resp
	}
	checkAccount.UpdatedAt = time.Now()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(checkAccount.Password), bcrypt.DefaultCost)
	checkAccount.Password = string(hashedPassword)
	
	if err := GetDB().Update(checkAccount); err != nil {
		response := u.Message(false, "update n'a pas fonctionner")
		return response
	}
	response := u.Message(true, "update a fonctionner")
	response["account"] = checkAccount
	return response


 */


	temp := &Account{}


	fmt.Println("aaaa")
	fmt.Println(params)
	fmt.Println("zzzzz")


	fmt.Println(temp)

	err := GetDB().Table("accounts").Where("ID = ?", params).First(temp).Error

	fmt.Println(err)

	fmt.Println(json.Email)
	fmt.Println(json.Password)
	fmt.Println(json.Birthdate)
	fmt.Println(json.Accesslevel)

	//si le sujet de vote n'existe pas
	if err != nil{
		return u.Message(false, "Il n'y a aucun user avec cette id")
	}else{
		//fmt.Println(err)
		if(json.Email != ""){
			temp.Email = json.Email
		}

		if(json.Password != ""){
			temp.Password = json.Password
		}

		if(!json.Birthdate.IsZero()){
			temp.Birthdate = json.Birthdate
		}

		if(json.Accesslevel <= 0 ){
			temp.Accesslevel = json.Accesslevel
		}

		if resp, ok := temp.Validate(); !ok {
			return resp
		}
		temp.UpdatedAt = time.Now()
		fmt.Println(temp)

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


