package model

type Review struct {
	ID   string `bson:"_id"`
	Book int64  `bson:"book"`
	Text string `bson:"text"`
}
