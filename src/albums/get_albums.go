package albums

import (
	"net/http"

	"github.com/Devisrisamidurai/go-rest-api/src"
	"github.com/gin-gonic/gin"
)

// get list of albums:
func GetAlbums(ctx *gin.Context) {
	statusCode := http.StatusOK // code 200
	ctx.IndentedJSON(statusCode, src.Albums)
}
