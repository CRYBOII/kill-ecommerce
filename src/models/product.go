package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" `
	Sku         string             `bson:"sku" json:"sku,omitempty"`
	Slug        string             `bson:"" json:"slug,omitempty"`
	Name        string             `bson:"" json:"name,omitempty"`
	Description string             `bson:"" json:"description,omitempty"`
	Details     struct {
		Model_number int64     `bson:"modelNumber" json:"modelNumber,omitempty"`
		Manufacturer string    `bson:"mnufacture" json:"manufaceture,omitempty"`
		Color        string    `bson:"color" json:"color,omitempty"`
		Mfg_date     time.Time `bson:"mfgDate" json:"mfgDate,omitempty"`
		Size         int64     `bson:"size" json:"size,omitempty"`
	}
	Pricing struct {
		Cost   int64 `bson:"cost" json:"cost"`
		Retail int64 `bson:"retail" json:"retail"`
	}
	Category primitive.ObjectID   `bson:"category" json:"category"`
	Tags     []string             `bson:"tags" json:"tags,omitempty"`
	Reviews  []primitive.ObjectID `bson:"review,omitempty" json:"review,omitempty"`
}
