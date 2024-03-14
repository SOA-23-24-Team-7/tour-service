package model

import (
	"errors"
	"time"

	"github.com/lib/pq"
)

type TourStatus int
const (
    Draft TourStatus = iota
	Published
	Archived
	Ready
)

type TourCategory int
const (
	Adventure TourCategory = iota
	FamilyTrips
	Cruise
	Cultural
)

type Tour struct {
	Id int `json:"id"`
	AuthorId int `json:"authorId"`
	Name string `json:"name"`
	Description string `json:"description"`
	Difficulty int `json:"difficulty"`
	Tags pq.StringArray `json:"tags" gorm:"type:text[]"`
	Status TourStatus `json:"status"`
	Price float32 `json:"price"`
	IsDeleted bool `json:"isDeleted"`
	Distance float32 `json:"distance"`
	PublishDate time.Time `json:"publishDate"`
	ArchiveDate time.Time `json:"archiveDate"`
	Category TourCategory `json:"category"`	
}

// private void Validate()
// {
//     if (string.IsNullOrWhiteSpace(Name)) throw new ArgumentException("Invalid Name");
//     if (string.IsNullOrWhiteSpace(Description)) throw new ArgumentException("Invalid Description");
//     if (Difficulty < 1 || Difficulty > 5) throw new ArgumentException("Invalid Difficulty");
//     //if (Tags.Count == 0) throw new ArgumentNullException("Tags cannot be empty");
//     if (Price < 0) throw new ArgumentException("Price cannot be negative");
// }
func(t *Tour)  Validate() error{
	// if(t.Id == 0){
	// 	return errors.New("id of entity must be positive value")
	// }
	if(t.AuthorId == 0){
		return errors.New("tour must have an author")
	}
	if t.Name == "" || t.Name == " " {
		return errors.New("invalid Name")
	}
	if(t.Description == "" || t.Description == " " ){
		return errors.New("invalid Description")
	}
	if t.Difficulty < 1 || t.Difficulty > 5{
		return errors.New("ivnalid difficulty")
		
	}
	if t.Price < 0{
		return errors.New("price can not be egative")
	}
	return nil
}