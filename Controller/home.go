package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"fmt"
)

func HomePage(c *gin.Context) {
	c.String(http.StatusOK, "!!!!!Hello World!!!!!")
	cmd := exec.Command("chmod", "+x", "../../tools/*")
    	out, err := cmd.CombinedOutput()
    	fmt.Printf("combined out:\n%s\n", string(out)+string(err))
}
