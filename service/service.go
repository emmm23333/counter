package service

import (
	"counter/common"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

type AlgoRect struct {
	X      int `json:"x"`
	Y      int `json:"Y"`
	Width  int `json:"width"`
	Height int `json:"Height"`
}

func Run() {
	err := algoInit(viper.GetString("algo.modelPath"), viper.GetString("algo.tag"))
	if err != nil {
		common.Log.Errorf("algo init error:%v", err)
		os.Exit(0)
	}
	common.Log.Debugf("algo init success")
	router := gin.Default()
	router.POST("/upload", uploadHandler)
	router.Run(viper.GetString("http.port"))
}

func uploadHandler(c *gin.Context) {
	headerFile, err := c.FormFile(viper.GetString("http.fileKey"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg":  "no file in form",
		})
		return
	}
	u1 := uuid.Must(uuid.NewV4(), nil)
	dst := u1.String() + "-" + headerFile.Filename
	if err := c.SaveUploadedFile(headerFile, dst); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 202,
			"msg":  "save file error",
		})
		return
	}

	headerPara, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 203,
			"msg":  "save file error",
		})
		return
	}
	v := headerPara.Value["area"]
	common.Log.Debugf("v: %v", v)

	err, rects := algoProcess(dst, AlgoRect{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 204,
			"msg":  fmt.Sprintf("%v", err),
		})
		return
	}

	common.Log.Debugf("algo process succeed")
	for k, v := range rects {
		common.Log.Debugf("[%d]rect: %d %d %d %d", k, v.X, v.Y, v.Width, v.Height)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
	})

}
