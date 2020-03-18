package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Telephone string             `bson:"telephone,omitempty" json:"telephone,omitempty"`
	Info      string             `bson:"info,omitempty" json:"info,omitempty"`
	TypeUser  string             `bson:"typeuser,omitempty" json:"typeuser,omitempty"` // pode ser um bool porem não aparece no bd

	Traning  *Traning  `bson:"traning,omitempty" json:"traning,omitempty"`
	Measures *Measures `bson:"measures,omitempty" json:"measures,omitempty"`
}

type Measures struct {
	Weight        float32 `bson:"weight,omitempty" json:"weight,omitempty"`
	Stature       float32 `bson:"stature,omitempty" json:"stature,omitempty"`
	Shoulder      float32 `bson:"shoulder,omitempty" json:"shoulder,omitempty"`
	InspiredChest float32 `bson:"inspired_chest,omitempty" json:"inspired_chest,omitempty"`
	RelaxedArm    float32 `bson:"relaxed_arm,omitempty" json:"relaxed_arm,omitempty"`
	Thigh         float32 `bson:"thigh,omitempty" json:"thigh,omitempty"`
	Forearm       float32 `bson:"forearm,omitempty" json:"forearm,omitempty"`
	ContractedArm float32 `bson:"contracted_arm,omitempty" json:"contracted_arm,omitempty"`
	Waist         float32 `bson:"waist,omitempty" json:"waist,omitempty"`
	Abdomen       float32 `bson:"abdomen,omitempty" json:"abdomen,omitempty"`
	Hip           float32 `bson:"hip,omitempty" json:"hip,omitempty"`
	Leg           float32 `bson:"leg,omitempty" json:"leg,omitempty"`
}

type Traning struct {
	Sequence   int    `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Place      string `bson:"place,omitempty" json:"place,omitempty"`
	Exercise   string `bson:"exercise,omitempty" json:"exercise,omitempty"`
	Series     int    `bson:"series,omitempty" json:"series,omitempty"`
	Repetition int    `bson:"repetition,omitempty" json:"repetition,omitempty"`
	Charge     int    `bson:"charge,omitempty" json:"charge,omitempty"`

	QrCode *QrCode `bson:"qrcode,omitempty" json:"qrcode,omitempty"`
}

type QrCode struct {
	Images       string `bson:"Images,omitempty" json:"Images,omitempty"` //jsonb
	Links        string `bson:"links,omitempty" json:"links,omitempty"`
	ImagesQrcode string `bson:"ImagesQrcode,omitempty" json:"ImagesQrcode,omitempty"` //jsonb
}
