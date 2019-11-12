package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
	u "github.com/api-projet/utils"

)


//a struct to rep user account
type Vote struct {
	gorm.Model
	UUID      string `json:"uuid" gorm:"primary_key"`
	Title    string `json:"title"`
	Description string `json:"description"`
	UUIDVote    string `json:"uuidvote"`
	StartDate  time.Time `json:"start_date"`
	EndDate  time.Time `json:"end_date"`

}


/*
**Vote** : ID (int), UUID (string), Title (string), Description (text), UUIDVote (collection), StartDate (time.Time), EndDate (time.Time), CreatedAt (time.Time), UpdatedAt (time.Time), DeletedAt (*time.Time).
L'UUID vote est une collection d'UUID d'utilisateurs ayant voté.
*/

func (vote *Vote) Validate() (map[string]interface{}, bool) {
	fmt.Println(vote)
	//Si il y a un titre
	if vote.Title == "" {
		return u.Message(false, "Votre titre ne peut pas être vide."), false
	} 


	//Si il y a une description
	if vote.Description == "" {
		return u.Message(false, "Votre description ne peut pas être vide."), false
	} 

	//Si la date de début < .. de fin && S'elles ne sont pas vides
	/*if vote.StartDate  == nil {
		return u.Message(false, "Vous ne pouvez pas avoir une date de départ supérieur à la date de fin."), false
	} */

		//Email must be unique
		temp := &Vote{}

		//check for errors and duplicate emails
		err := GetDB().Table("vote").Where("uuid = ?", temp.UUID).First(temp).Error
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Vote non trouvé"), false
		}
		if temp.Title != "" {
			return u.Message(false, "Le titre de ce sujet de vote a déjà été pris."), false
		}

		if temp.Description != "" {
			return u.Message(false, "La description de ce sujet de vote a déjà été pris."), false
		}

	return u.Message(false, "Recquis respectés"), true
}


func (vote *Vote) Create() (map[string]interface{}) {

	if resp, ok := vote.Validate(); !ok {
		return resp
	}

	vote.UUID = uuid.NewV4().String()

	GetDB().Create(vote)

	response := u.Message(true, "Le sujet de vote a été créé")
	response["vote"] = vote
	return response
}

func UpdateVote(params interface{},json *Vote) (map[string]interface{}) {

	temp := &Vote{}
	err := GetDB().Table("vote").Where("ID = ?", params).First(temp).Error
	
	//fmt.Println(json)

	//si le sujet de vote n'existe pas
	if err == nil || err == gorm.ErrRecordNotFound {
		return u.Message(false, "Il n'y a aucun sujet de vote avec cet titre")
	}else{
		//fmt.Println(err)
		if(json.Title != ""){
			temp.Title = json.Title
		}

		if(json.Description != ""){
			temp.Description = json.Description
		}

		if(json.StartDate.IsZero()){
			temp.StartDate = json.StartDate
		}

		if(json.EndDate.IsZero()){
			temp.EndDate = json.StartDate
		}
		

		temp.UpdatedAt = time.Now()

		GetDB().Update(temp)
	}

	response := u.Message(true, "Le sujet de vote a été édité")
	response["vote"] = temp
	return response
}

func DeleteVote(params interface{}, json *Vote)  (map[string]interface{}) {
	temp := &Vote{}
	err := GetDB().Table("vote").Where("ID = ?", params).First(temp).Error
	
	//si le sujet de vote n'existe pas
	if err == nil || err == gorm.ErrRecordNotFound {
		return u.Message(false, "Il n'y a aucun sujet de vote avec cet titre")
	}
		//var checkVote Vote
		//db.Where("ID = ?", params).Find(&checkVote)
	db.Delete(&temp)
	

	response := u.Message(true, "Le sujet de vote a été supprimé")
	response["vote"] = temp
	return response

} 


