package model

import "gopkg.in/mgo.v2/bson"

type ClientModel struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	TitleTh   string        `json:"title_th" bson:"title_th"`
	BranchTh  string        `json:"branch_th" bson:"branch_th"`
	AddressTh string        `json:"address_th" bson:"address_th"`
	Taxid     string        `json:"taxid" bson:"taxid"`
}
