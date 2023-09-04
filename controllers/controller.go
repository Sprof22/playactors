package controllers

import (
	"encoding/json"
	"log"
	"playactors/initializer"
	"playactors/models"
	"time"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/mitchellh/mapstructure"
)

var RedisClient *redis.Client
var AlgoliaClient *algoliasearch.Client
var AlgoliaIndex algoliasearch.Index

func init() {
	// Initialize Redis client using the environment variable
	redisURL := "redis-cli --tls -u rediss://default:c84511a1414541deb264399bd6e53cad@usw1-ready-killdeer-33564.upstash.io:33564"
	if redisURL == "" {
		panic("REDIS_URL environment variable is not set")
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // Redis password if required
		DB:       0,  // Default DB.
	})

	// Initialize Algolia client and index
	algoliaAppID := "SHITSH8101"
	algoliaAPIKey := "5e71825b63c0e3a3e6f2f1649f1fd870"
	if algoliaAppID == "" || algoliaAPIKey == "" {
		log.Fatal("Algolia credentials not set in environment variables")
	}

	AlgoliaClient := algoliasearch.NewClient(algoliaAppID, algoliaAPIKey)
	AlgoliaIndex = AlgoliaClient.InitIndex("playactors")

}

func CreateActor(c *gin.Context) {

	var body struct {
		ActorName       string
		ActorRating     int
		ImagePath       string
		AlternativeName string
		ActorID         int
	}

	c.Bind(&body)
	actor := models.Actors{ActorName: body.ActorName, ActorRating: body.ActorRating, ImagePath: body.ImagePath, AlternativeName: body.AlternativeName, ActorID: body.ActorID}

	result := initializer.DB.Create(&actor)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// Clear the cache since data has been modified
	_ = RedisClient.Del("actors")

	// Add the actor to Algolia index
	_, err := AlgoliaIndex.AddObject(map[string]interface{}{
		"ActorName":       actor.ActorName,
		"ActorRating":     actor.ActorRating,
		"ImagePath":       actor.ImagePath,
		"AlternativeName": actor.AlternativeName,
		"ActorID":         actor.ActorID,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to add actor to Algolia index",
		})
		return
	}

	c.JSON(200, gin.H{
		"actor": actor,
	})
}

func GetActors(c *gin.Context) {
	// Check if a search query parameter is provided
	query := c.DefaultQuery("query", "")

	if query != "" {
		// Check if the search results are already cached in Redis
		cachedResults, err := RedisClient.Get(query).Result()
		if err == nil {
			var actors []models.Actors
			if err := json.Unmarshal([]byte(cachedResults), &actors); err == nil {
				c.JSON(200, gin.H{
					"message": "Search results fetched from Redis cache.",
					"actors":  actors,
				})
				return
			}
		}

		// Perform Algolia search using AlgoliaIndex
		searchResponse, err := AlgoliaIndex.Search(query, algoliasearch.Map{})
		if err != nil {
			c.JSON(500, gin.H{"error": "Algolia search error"})
			return
		}

		// Process search results
		var actors []models.Actors
		for _, hit := range searchResponse.Hits {
			// Convert Algolia search results to your model
			var actor models.Actors
			if err := mapstructure.Decode(hit, &actor); err != nil {
				continue
			}
			actors = append(actors, actor)
		}

		// Cache the search results in Redis for future use
		if serializedData, err := json.Marshal(actors); err == nil {
			_ = RedisClient.Set(query, serializedData, 24*time.Hour)
		}

		c.JSON(200, gin.H{
			"message": "Search results fetched from Algolia and cached in Redis.",
			"actors":  actors,
		})
	} else {
		// Try to get actors from cache
		cachedData, err := RedisClient.Get("actors").Result()
		if err == nil {
			var actors []models.Actors
			if err := json.Unmarshal([]byte(cachedData), &actors); err == nil {
				c.JSON(200, gin.H{
					"message": "Actors list fetched from Redis cache.",
					"actors":  actors,
				})
				return
			}
		}

		var actors []models.Actors
		initializer.DB.Find(&actors)

		// Serialize and store data in cache
		if serializedData, err := json.Marshal(actors); err == nil {
			_ = RedisClient.Set("actors", serializedData, 24*time.Hour)
		}

		c.JSON(200, gin.H{
			"message": "Actors list fetched from database and cached in Redis.",
			"actors":  actors,
		})
	}
}

func GetSingleActor(c *gin.Context) {

	id := c.Param("id")
	var actor models.Actors
	initializer.DB.First(&actor, id)

	c.JSON(200, gin.H{
		"actors": actor,
	})
}

func DeleteActor(c *gin.Context) {
	id := c.Param("id")
	var actor models.Actors
	initializer.DB.Delete(&models.Actors{}, id)

	_ = RedisClient.Del("actors")

	c.JSON(200, gin.H{
		"actor": actor,
	})
}

func UpdateActor(c *gin.Context) {

	var body struct {
		ActorName       string
		ActorRating     int
		ImagePath       string
		AlternativeName string
		ActorID         int
	}

	id := c.Param("id")
	var actor models.Actors
	initializer.DB.First(&actor, id)
	c.Bind(&body)

	initializer.DB.Model(&actor).Updates(models.Actors{ActorName: body.ActorName, ActorRating: 18, ImagePath: body.ImagePath, AlternativeName: body.AlternativeName, ActorID: body.ActorID})

	_ = RedisClient.Del("actors")
	c.JSON(200, gin.H{

		"actor": actor,
	})
}
