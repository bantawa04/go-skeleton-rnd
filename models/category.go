package models

// Category DB model
type Category struct {
	Base
	Title          string     `json:"title"`
}

// TableName returns table name of model
func (c Category) TableName() string {
	return "categories"
}

// ToMap  maps category
func (c Category) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":              c.ID,
		"title":           c.Title,
	}
}
