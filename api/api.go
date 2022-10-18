package api

import (
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Audio struct {
	AudioName           string `json:"audioname"`
	AudioSize           int64  `json:"audiosize"`
	AudioFileModifiedAt string `json:"audiodfilemodifiedat"`
}

func UploadAudio(c *gin.Context) {
	file, err := c.FormFile("data-binary")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		fileName := strconv.Itoa(int(time.Now().UnixMicro()))
		c.SaveUploadedFile(file, "./wav/"+fileName+".wav")
		cmd := exec.Command("lame", "./wav/"+fileName+".wav", "./mp3/"+fileName+".mp3")
		if _, err := cmd.CombinedOutput(); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		os.Remove("./wav/" + fileName + ".wav")
		c.IndentedJSON(http.StatusOK, gin.H{"message": "File sucessfully uploaded!", "AudioFileID": fileName + ".mp3"})
	}
}

func DeleteAudio(c *gin.Context) {
	audioID := c.Param("audioid")
	if audioID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "You are missing audioID parameter."})
		return
	} else {
		err := os.Remove("./mp3/" + audioID)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "audio not found"})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "The audio has been deleted successfully!"})
	}
}

func GetAudio(c *gin.Context) {
	audioID := c.Param("audioid")
	if audioID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "You are missing audioID parameter."})
		return
	} else {
		_, err := os.Stat("./mp3/" + audioID)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "audio not found"})
			return
		}
		c.FileAttachment("./mp3/"+audioID, audioID)
	}
}

func GetAudioInfo(c *gin.Context) {
	audioID := c.Param("audioid")
	var response = Audio{}
	if audioID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "You are missing audioID parameter."})
	} else {
		fileStat, err := os.Stat("./mp3/" + audioID)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "audio not found"})
			return
		}
		response = Audio{
			AudioName:           fileStat.Name(),
			AudioSize:           fileStat.Size(),
			AudioFileModifiedAt: fileStat.ModTime().String(),
		}
		c.IndentedJSON(http.StatusOK, response)
	}
}
