package checker

import (
	"broken-link-checker/internal/delivery/http/api/v1/response"
	"broken-link-checker/internal/models"
	"broken-link-checker/internal/service/linkChecker"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type responseData struct {
	Duration   string
	BreakLinks []string
	Info       linkChecker.Info
}

func SearchBrokenLinks(c *gin.Context) {
	var data models.CheckerRequestData

	if err := c.BindJSON(&data); err != nil {
		log.Println("checker -> SearchBrokenLinks: error parse data. reason: ", err.Error())
		response.Error(c, "Error in request data: "+err.Error())
		return
	}

	fmt.Println("Run found: ", data.Link)

	checker := linkChecker.New(data.Link, data.Depth)
	if err := checker.Run(); err != nil {
		log.Println("checker -> SearchBrokenLinks: linkChecker error. reason: ", err.Error())
		response.Error(c, "Server error: "+err.Error())
		return
	}

	log.Println("Broken links found: ", len(checker.GetBreakLinks()))
	fmt.Printf("Time spent: %s\n", checker.GetDuration())

	response.Success(c, responseData{
		BreakLinks: checker.GetBreakLinks(),
		Info:       checker.GetInfo(),
	})
}
