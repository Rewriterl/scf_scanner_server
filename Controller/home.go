package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"log"
	"fmt"
)

func HomePage(c *gin.Context) {
	c.String(http.StatusOK, "!!!!!Hello World!!!!!")
	cmd := exec.Command("chmod", "+x", "../../tools/*")
    	out, err := cmd.CombinedOutput()
    	if err != nil {
        	log.Fatalf("cmd.Run() failed with %s\n", err)
   	}
    	fmt.Printf("combined out:\n%s\n", string(out))
}
