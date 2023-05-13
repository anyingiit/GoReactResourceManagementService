package models

import (
	"encoding/json"

	"github.com/anyingiit/GoReactResourceManagement/structs"
	"github.com/anyingiit/GoReactResourceManagement/utils"
	"gorm.io/gorm"
)

type InvateClient struct {
	gorm.Model
	InvateCode string
	ClientID   uint
	Client     Client `gorm:"foreignKey:ClientID"`
}

func (i *InvateClient) GenerateInvateCode(data *structs.InvateClientMessage) error {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	i.InvateCode = utils.BytesToBase64String(jsonByte)
	return nil
}
