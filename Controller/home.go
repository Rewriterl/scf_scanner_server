package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"fmt"
)

func HomePage(c *gin.Context) {
	c.String(http.StatusOK, "!!!!!Hello World!!!!!")
	cmd := exec.Command("chmod", "+x", "./tools/*")
    	out,err := cmd.CombinedOutput()
	if (err==nil){
    	  fmt.Printf("combined out:\n%s\n", string(out))
	}
	cmd = exec.Command("ls","tools/","-al")
	out,err = cmd.CombinedOutput()
	if (err==nil){
    	  fmt.Printf("combined out:\n%s\n", string(out))
	}
}
