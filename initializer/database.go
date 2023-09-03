package initializer

import (
	"fmt"
	"log"
	"os"
	"playactors/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	redisURL := os.Getenv("REDIS_URL")
	fmt.Printf("Here is ther redis url %v", redisURL)
	if redisURL == "" {
		log.Fatal("REDIS_URL environment variable is not set")
	}

	//TODO: connect to db here.
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection Failed")
	}
}

func IsDatabaseEmpty() bool {
	var count int64
	result := DB.Model(&models.Actors{}).Count(&count)
	if result.Error != nil {
		log.Printf("Error counting actors: %v\n", result.Error)
		return true // Consider it empty if there's an error.
	}
	return count == 0
}

func CreateDummyActorsIfEmpty() {
	if IsDatabaseEmpty() {
		createDummyActors()
	}
}

func createDummyActors() {
	actors := []models.Actors{
		{
			ActorName:       "Rebecca Ferguson",
			ActorRating:     737,
			ImagePath:       "/a8thKB6ZnHxxqiC7crwCyaPX63d.jpg",
			AlternativeName: "",
			ActorID:         551487410,
		},
		{
			ActorName:       "Lee Majors",
			ActorRating:     734,
			ImagePath:       "/1CjhU32qX30hg8TOXju5KY2THkS.jpg",
			AlternativeName: "",
			ActorID:         551487440,
		},
		{
			ActorName:       "Philip Seymour Hoffman",
			ActorRating:     734,
			ImagePath:       "/de37JbzZ80KP1LOhzIkVe5XfSwe.jpg",
			AlternativeName: "",
			ActorID:         551487430,
		},
		{
			ActorName:       "Julia Stiles",
			ActorRating:     734,
			ImagePath:       "/wqFYU1IS1xhn4yBjlkXt9LwFYr0.jpg",
			AlternativeName: "",
			ActorID:         551487420,
		},
		{
			ActorName:       "Paul Giamatti",
			ActorRating:     730,
			ImagePath:       "/rX4LRmkYshMRfQ6lVbeZVAfqVKI.jpg",
			AlternativeName: "",
			ActorID:         551487450,
		},
		{
			ActorName:       "Salma Hayek",
			ActorRating:     729,
			ImagePath:       "/zMmEPWSqpACzsP5TnLdETV8j7eW.jpg",
			AlternativeName: "",
			ActorID:         551487460,
		},
		{
			ActorName:       "Anna Faris",
			ActorRating:     727,
			ImagePath:       "/eHh3ZVEdBlXSBNjpHaGkGKvx1QI.jpg",
			AlternativeName: "Anna Kay Faris",
			ActorID:         551487480,
		},
		{
			ActorName:       "Jon Hamm",
			ActorRating:     727,
			ImagePath:       "/7sjEnFaFNOzPeu5GhCeNJWhnOLt.jpg",
			AlternativeName: "",
			ActorID:         551487470,
		},
		{
			ActorName:       "Sandra Bullock",
			ActorRating:     725,
			ImagePath:       "/pFudVrL9n8L0AHwMpbcfvsrjUQy.jpg",
			AlternativeName: "Sandra Annette Bullock",
			ActorID:         551487490,
		},
		{
			ActorName:       "Cate Blanchett",
			ActorRating:     723,
			ImagePath:       "/X3CMrI6lkzLdS5ZQqQWeRJkAGU.jpg",
			AlternativeName: "Catherine Elise Blanchett",
			ActorID:         551487500,
		},
	}

	for _, actor := range actors {
		result := DB.Create(&actor)
		if result.Error != nil {
			log.Printf("Error creating actor: %v\n", result.Error)
		} else {
			log.Printf("Actor created: %s\n", actor.ActorName)
		}
	}
}
