package lib

import "time"

type Timed struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Feedback struct {
	// embedded struct
	Timed
	Comment string
	Rating  int
}

// constructor
func NewFeedback(rating int) *Feedback {
	// for literal struts, the embedded struct should be explicitly declared
	return &Feedback{Rating: rating, Timed: Timed{CreatedAt: time.Now()}}
}

// testing if structs are passed by reference or by value
// checked - passed by value
func SetRating(feedback Feedback, rating int) {
	feedback.Rating = rating
}

func SetRatingPtr(feedback *Feedback, rating int) {
	// the * can be omitted for structs
	feedback.Rating = rating
}

// method with pointer receiver
func (f *Feedback) Timezone() string {
	// embedded struct fields can be accessed directly
	return f.CreatedAt.Location().String()
}

// method with value receiver
func (f Feedback) RatingDescription() string {
	if f.Rating < 3 {
		return "below average"
	} else if f.Rating == 3 {
		return "average"
	} else {
		return "above average"
	}
}
