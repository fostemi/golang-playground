
package main

import (

  "net/http"

  "github.com/gin-gonic/gin"

)

type movie struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Year int `json:"year"`
  Rating string `json:"rating"`
}

var movies = []movie {
  {ID: "1", Title: "Caddy Shack", Year: 1980, Rating: "Rotten Tomatoes: 72%"},
}

func getMovies(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, movies)
}

func postMovies(c *gin.Context) {
  var newMovie movie

  if err := c.BindJSON(&newMovie); err != nil {
    return
  }

  movies = append(movies, newMovie)
  c.IndentedJSON(http.StatusCreated, newMovie)
}

func getMovieByID(c *gin.Context) {
  id := c.Param("id")

  for _, m := range movies {
    if m.ID == id {
      c.IndentedJSON(http.StatusOK, m)
      return
    }
  }
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
  router := gin.Default()
  router.GET("/movies", getMovies)
  router.GET("/movies/:id", getMovieByID)
  router.POST("/movies", postMovies)

  router.Run("localhost:8080")
}

