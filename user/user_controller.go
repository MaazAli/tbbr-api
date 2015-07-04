package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Index is used when the user's index is routed to
// this handler will run. Generally, it will
// come with some query parameters like limit and offset
// @returns an array of users
func Index(c *gin.Context) {
	user := User{
		Name:      "Maaz Ali",
		Username:  "maaz_ali",
		CreatedAt: time.Now(),
		Email:     "maazali40@gmail.com",
	}

	c.JSON(http.StatusOK, user)
}

// Show is used to show one specific user
// *returns a user struct
func Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"userShow": c.Param("id")})
}

// Create is used to create one specific user, it'll come with some form data
// @returns the newly created user struct
func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"userCreate": "someContent"})
}

// Update is used to update a specific group, it'll also come with some form data
// @returns a user struct
func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"userUpdate": "someContent"})
}

// Delete is used to delete one specific user with a `id`
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"userDelete": "someContent"})
}
