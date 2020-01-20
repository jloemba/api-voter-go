package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
	u "github.com/api-projet/utils"
	"github.com/lib/pq"
)


//a struct to rep user account
type Vote struct {
	gorm.Model
	UUID      string `json:"uuid" gorm:"primary_key"`
	Title    string `json:"title"`
	Description string `json:"description"`
	UUIDVote    pq.StringArray `json:"uuid_votes" gorm:"type:varchar(64)[]"`
	StartDate  time.Time `json:"start_date"`
	EndDate  time.Time `json:"end_date"`

}


/*
**Vote** : ID (int), UUID (string), Title (string), Description (text), UUIDVote (collection), StartDate (time.Time), EndDate (time.Time), CreatedAt (time.Time), UpdatedAt (time.Time), DeletedAt (*time.Time).
L'UUID vote est une collection d'UUID d'utilisateurs ayant voté.
*/

func (vote *Vote) Validate() (map[string]interface{}, bool) {

	//Si il y a un titre
	if vote.Title == "" {
		return u.Message(false, "Votre titre ne peut pas être vide."), false
	} 


	//Si il y a une description
	if vote.Description == "" {
		return u.Message(false, "Votre description ne peut pas être vide."), false
	} 

		//Email must be unique
		temp := &Vote{}

		//check for errors and duplicate emails
		err := GetDB().Table("votes").Where("uuid = ?", temp.UUID).First(temp).Error

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

	/*if resp, ok := vote.Validate(); !ok {
		return resp
	}*/


	vote.UUID = uuid.NewV4().String()

	GetDB().Create(vote)

	response := u.Message(true, "Le sujet de vote a été créé")
	response["vote"] = vote
	return response
}

func UpdateVote(params string,json *Vote) (map[string]interface{}) {

	row := &Vote{}
	err := GetDB().Table("votes").Where("UUID = ?", params).First(row).Error
	fmt.Println(err)

	//si le sujet de vote n'existe pas
	if row.Title == "" {
		return u.Message(false, "Il n'y a aucun sujet de vote avec cet titre")
	}else{

		if(json.Title != ""){
			row.Title = json.Title
		}

		if(json.Description != ""){
			row.Description = json.Description
		}

		if(json.StartDate.IsZero()){
			row.StartDate = json.StartDate
		}

		if(json.EndDate.IsZero()){
			row.EndDate = json.StartDate
		}

		row.UpdatedAt = time.Now()

		GetDB().Model(&row).Update(row)
	}

	response := u.Message(true, "Le sujet de vote a été édité")
	response["vote"] = row
	return response
}

func DeleteVote(params string, json *Vote)  (map[string]interface{}) {
	row := &Vote{}
	err := GetDB().Table("votes").Where("UUID = ?", params).First(row).Error
	
	//si le sujet de vote n'existe pas
	if row.Title == "" {
		fmt.Println(err)
		return u.Message(false, "Il n'y a aucun sujet de vote avec cet titre")
	}

	db.Delete(&row)
	

	response := u.Message(true, "Ce sujet de vote a bien été supprimé")
	response["vote"] = row
	return response

} 

func SingleVote(params string, json *Vote)  (map[string]interface{}) {

	row := &Vote{}
	err := GetDB().Table("votes").Where("UUID = ?", params).First(row).Error

	if row.Title == "" {
		fmt.Println(err)
		return u.Message(false, "Il n'y a aucun sujet de vote avec cet titre")
	}

	response := u.Message(true, "Le sujet de vote")
	response["vote"] = row
	return response

}


func SubmitVote(uuidvote pq.StringArray,json *Vote) (map[string]interface{}) {

	//Check si le sujet de vote existe via son ID
	vote := &Vote{}
	err := GetDB().Table("votes").Where("UUID = ?", json.UUID).First(vote).Error

	if err != nil {
		return u.Message(false, "Il n'y a aucun sujet de vote avec cet UUID")
	}

	vote.UUIDVote = append(vote.UUIDVote, uuidvote[0])

	GetDB().Model(&vote).Update(vote)

	/*else{
		fmt.Println("Vote existant")
		fmt.Println(row)
	}*/


		//modifier le vote pour y mettre l'uuidvote 
	//rowAccount := &Account{}
	//erraccount := GetDB().First(&rowAccount, uuidaccount)

	//if erraccount != nil{
		//return u.Message(false,"Not found")
	//}

	//récupérer le vote
	//rowVote := &Vote{}
	//err := GetDB().First(&rowVote, uuidvote)
	//fmt.Println(uuidaccount)
	//rowVote.UUIDVote = append(rowVote.UUIDVote, uuidaccount)



	//if err != nil{
		//return u.Message(false,"Not found")
	//}

	response := u.Message(true, "Le vote a été soumis")
	response["vote"] = vote
	return response
}


func FetchVote() (map[string]interface{}) {

	vt := []Vote{}
	GetDB().Table("votes").Find(&vt)
	
	response := u.Message(true, "Liste des votes existantes")
	response["vote"] = vt
	return response

}
