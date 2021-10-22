package database

import (
	"github.com/shivamtiwari/models"
)

var MovieList = []models.Movie{
	{
		Id:         "1",
		Title:      "The GodFather",
		Category:   "Crime/Drama",
		Year:       "1972",
		ImdbRating: "9.0/10",
	},
	{
		Id:         "2",
		Title:      "The Dark Night",
		Category:   "Crime/Drama",
		Year:       "2008",
		ImdbRating: "9.0/10",
	},
	{
		Id:         "3",
		Title:      "Fight Club",
		Category:   "Drama",
		Year:       "1999",
		ImdbRating: "8.8/10",
	},
	{
		Id:         "4",
		Title:      "The English Patient",
		Category:   "Romance/Drama",
		Year:       "1996",
		ImdbRating: "7.4/10",
	},
}
