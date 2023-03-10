package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"walpapperCollectRestAPI/lib/tools"
)

func WallpaperCollection(c *gin.Context) {

	id, err := tools.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	dir1 := http.Dir("././assets/" + id + "/wallpaper_collection")
	dir2 := http.Dir("././assets/" + id + "/profile")
	UserData, err := tools.GetUserDataWithId(id)
	c.FileFromFS("", dir2)
	c.JSON(http.StatusOK, gin.H{
		"user_name":    UserData.UserName,
		"description":  UserData.Description,
		"phone_number": UserData.PhoneNumber,
		"email":        UserData.Email,
	})

	c.FileFromFS("", dir1)
}
