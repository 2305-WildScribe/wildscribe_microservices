package model

// RecordID defines a record id. Together with RecordType identifies a unique record across all types.
type RecordID string

// RecordType defines a record type. Together with RecordID identifies a unique record across all types.
type RecordType string

// Existing Record types.
const (
	RecordTypeMovie = RecordType("movie")
)

// UserID defines a user id.
type UserID string

// RatingValue defines a value of a rating record.
type RatingValue int 

// Rating defines an individual rating created by a user.
type Rating struct {
	RecordID 		string 				`json:"recordId"`
	RecordType 	string 				`json:"recordType"`
	UserID 			UserID 				`json:"userId"`
	Value 			RatingValue 	`json:"value"`
}