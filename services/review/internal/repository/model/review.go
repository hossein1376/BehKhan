package model

type Review struct {
	ID   string `bson:"_id"`
	Text string `bson:"text"`
}
