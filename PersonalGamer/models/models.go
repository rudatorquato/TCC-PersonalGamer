package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Telephone string             `bson:"telephone" json:"telephone"`
	Info      string             `bson:"info" json:"info"`
	TypeUser  string             `bson:"typeuser" json:"typeuser"` // pode ser um bool porem não aparece no bd

	Traning  *Traning  `bson:"traning" json:"traning"`
	Measures *Measures `bson:"measures" json:"measures"`
}

type Measures struct {
	Weight        float32 `bson:"weight" json:"weight"`
	Stature       float32 `bson:"stature" json:"stature"`
	Shoulder      float32 `bson:"shoulder" json:"shoulder"`
	InspiredChest float32 `bson:"inspired_chest" json:"inspired_chest"`
	RelaxedArm    float32 `bson:"relaxed_arm" json:"relaxed_arm"`
	Thigh         float32 `bson:"thigh" json:"thigh"`
	Forearm       float32 `bson:"forearm" json:"forearm"`
	ContractedArm float32 `bson:"contracted_arm" json:"contracted_arm"`
	Waist         float32 `bson:"waist" json:"waist"`
	Abdomen       float32 `bson:"abdomen" json:"abdomen"`
	Hip           float32 `bson:"hip" json:"hip"`
	Leg           float32 `bson:"leg" json:"leg"`
}

type Traning struct {
	Sequence   int    `bson:"sequence" json:"sequence"`
	Place      string `bson:"place" json:"place"`
	Exercise   string `bson:"exercise" json:"exercise"`
	Series     int    `bson:"series" json:"series"`
	Repetition int    `bson:"repetition" json:"repetition"`
	Charge     int    `bson:"charge" json:"charge"`

	QrCode *QrCode `bson:"qrcode" json:"qrcode"`
}

type QrCode struct {
	Images       string `bson:"Images" json:"Images"` //jsonb
	Links        string `bson:"links" json:"links"`
	ImagesQrcode string `bson:"ImagesQrcode" json:"ImagesQrcode"` //jsonb
}
