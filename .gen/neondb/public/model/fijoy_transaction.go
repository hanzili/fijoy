//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type FijoyTransaction struct {
	ID              string    `sql:"primary_key" json:"ID"`
	TransactionType string    `json:"TransactionType"`
	Amount          float64   `json:"Amount"`
	Currency        string    `json:"Currency"`
	FromAccountID   *string   `json:"FromAccountID"`
	ToAccountID     *string   `json:"ToAccountID"`
	UserID          string    `json:"UserID"`
	WorkspaceID     string    `json:"WorkspaceID"`
	Datetime        time.Time `json:"Datetime"`
	Note            *string   `json:"Note"`
	CategoryID      *string   `json:"CategoryID"`
	PayeeName       *string   `json:"PayeeName"`
	PayerName       *string   `json:"PayerName"`
	TagName         *string   `json:"TagName"`
}
