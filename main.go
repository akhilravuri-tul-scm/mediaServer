package main

import (
	"mediaserver/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/v0/audio/upload", api.UploadAudio)
	router.DELETE("/v0/audio/:audioid", api.DeleteAudio)
	router.GET("/v0/audio/:audioid", api.GetAudio)
	router.GET("/v0/audio/:audioid/info", api.GetAudioInfo)

	router.Run("localhost:8080")
}
