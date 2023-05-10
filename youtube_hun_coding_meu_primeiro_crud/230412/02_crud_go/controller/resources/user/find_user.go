package user

import "github.com/gin-gonic/gin"

// The gin.Context have everything about the request.
func (u userController) FindUserById(c *gin.Context) {
}

func (u userController) FindUserByEmail(c *gin.Context) {}
