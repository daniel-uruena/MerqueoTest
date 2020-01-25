package Models

type Product struct {
	Id int `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Quantity int `json:"quantity,omitempty" bson:"quantity,omitempty"`
}