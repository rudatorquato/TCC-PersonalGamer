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
	Weight             float32 `bson:"weight" json:"weight"`
	Stature            float32 `bson:"stature" json:"stature"`
	Shoulder           float32 `bson:"shoulder" json:"shoulder"`
	InspiredChest      float32 `bson:"inspired_chest" json:"inspired_chest"`
	LeftRelaxedArm     float32 `bson:"left_relaxed_arm" json:"left_relaxed_arm"`
	RightRelaxedArm    float32 `bson:"right_relaxed_arm" json:"right_relaxed_arm"`
	LeftThigh          float32 `bson:"left_thigh" json:"left_thigh"`
	RightThigh         float32 `bson:"right_thigh" json:"right_thigh"`
	LeftForearm        float32 `bson:"left_forearm" json:"left_forearm"`
	RightForearm       float32 `bson:"right_forearm" json:"right_forearm"`
	LeftContractedArm  float32 `bson:"left_contracted_arm" json:"left_contracted_arm"`
	RightContractedArm float32 `bson:"right_contracted_arm" json:"right_contracted_arm"`
	Waist              float32 `bson:"waist" json:"waist"`
	Abdomen            float32 `bson:"abdomen" json:"abdomen"`
	Hip                float32 `bson:"hip" json:"hip"`
	LeftLeg            float32 `bson:"left_leg" json:"left_leg"`
	RightLeg           float32 `bson:"right_leg" json:"right_leg"`
}

type Traning struct {
	NameTraning string `bson:"name_traning" json:"name_traning"`
	Sequence    int    `bson:"sequence" json:"sequence"`
	Place       string `bson:"place" json:"place"`
	Exercise    string `bson:"exercise" json:"exercise"`
	Series      int    `bson:"series" json:"series"`
	Repetition  int    `bson:"repetition" json:"repetition"`
	Charge      int    `bson:"charge" json:"charge"`

	//QrCode *QrCode `bson:"qrcode,omitempty" json:"qrcode,omitempty"`
}

type QrCode struct {
	Images       string `bson:"Images" json:"Images"` //jsonb
	Links        string `bson:"links" json:"links"`
	ImagesQrcode string `bson:"ImagesQrcode" json:"ImagesQrcode"` //jsonb
}
