package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
  HistorySearch map[string]string
)

func iKnowHandler(c *gin.Context) {
  search := c.Query("q")

  result := HistorySearch[search]
  if result == "" {
    // Convert map to slice
    likeValues := []string{}
    values := []string{}
    for _, value := range HistorySearch {
      if strings.Contains(value, search) {
        likeValues = append(likeValues, value)
      }
      values = append(values, value)
    }

    if len(likeValues) > 0 {
      c.JSON(http.StatusOK, likeValues)
      HistorySearch[search] = search
      return
    }

    c.JSON(http.StatusOK, values)
    HistorySearch[search] = search
    return
  }

  c.JSON(http.StatusOK, []string{HistorySearch[search]})
}


func main() {
  fmt.Println("Start server")
  // Initialize data store
	HistorySearch = make(map[string]string)

  r := gin.Default()
  api := r.Group("/api")

  {
    api.GET("/iknow", iKnowHandler)
  }

	r.Run(":8000") // listen and serve on 0.0.0.0:8000
}
