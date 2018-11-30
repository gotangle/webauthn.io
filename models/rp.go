package models

import (
	"fmt"
)

// RelyingParty is the group the User is authenticating with
type RelyingParty struct {
	ID          string `json:"id" gorm:"not null;unique" gorm:"primary_key;size:255"`
	DisplayName string `json:"display_name" sql:"not null"`
	Icon        string `json:"icon,omitempty"`
	Users       []User `json:"users,omitempty" gorm:"many2many:user_relying_parties"`
}

// GetRelyingPartyByHost gets the RP by hostname which in this case is the ID
func GetRelyingPartyByHost(hostname string) (RelyingParty, error) {
	rp := RelyingParty{}
	err := db.Where("id = ?", hostname).First(&rp).Error
	if err != nil {
		return rp, err
	}
	return rp, nil
}

// PutRelyingParty creates or updates a Relying Party
func PutRelyingParty(rp *RelyingParty) error {
	if db.NewRecord(&rp) {
		fmt.Println("New Relying Party Added")
	}
	err := db.Save(&rp).Error
	return err
}
