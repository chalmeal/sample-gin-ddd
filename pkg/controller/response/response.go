package response

import (
	"net/http"
	e "sample-gin-ddd/pkg/errors"
	"sample-gin-ddd/pkg/usecase/dto"

	"golang.org/x/exp/slices"

	"github.com/gin-gonic/gin"
)

func Res(c *gin.Context, result *dto.Dto) {
	status := http.StatusOK

	if slices.Contains(e.BadRequest, result.Error) {
		status = http.StatusBadRequest
	} else if slices.Contains(e.NotFound, result.Error) {
		status = http.StatusNotFound
	} else if slices.Contains(e.UnAuthorized, result.Error) {
		status = http.StatusUnauthorized
	} else if e.INTERNAL_SERVER_ERROR == result.Error {
		status = http.StatusInternalServerError
	}

	if status != http.StatusOK {
		c.JSON(status, gin.H{
			"Error": result.Error.Error(),
		})
	} else {
		c.JSON(status, gin.H{
			"Response": result,
		})
	}

}
