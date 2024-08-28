package main

import (
          "fmt"
          "net/http"
)

func main() {
          router := gin.Default()
          router.GET("/hello/:name", func(c *gin.Context) {
                    name := c.Param("name")
                    c.String(http.StatusOK, makeGreetin(name))
          })
          router.Run(":8080")
}

func makeGreetin(name string) string {
          return fmt.Sprintf("Hello, %d", name)
}
