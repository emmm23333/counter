package service

import (
	"counter/common"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

type AlgoRect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type AlgoResponse struct {
	Code  int        `json:"code"`
	Msg   string     `json:"msg"`
	Rects []AlgoRect `json:"rects"`
}

func Run() {
	err := algoInit(viper.GetString("algo.modelPath"), viper.GetString("algo.tag"))
	if err != nil {
		common.Log.Errorf("algo init error:%v", err)
		os.Exit(0)
	}
	common.Log.Debugf("algo init success")
	router := gin.Default()
	router.POST(viper.GetString("http.uri"), uploadHandler)
	router.Run(viper.GetString("http.port"))
}

func uploadHandler(c *gin.Context) {
	resp := &AlgoResponse{
		Code: 200,
		Msg:  "success",
	}

	headerFile, err := c.FormFile(viper.GetString("http.fileKey"))
	if err != nil {
		resp.Code = 201
		resp.Msg = "no file in form"
		uploadResponse(c, resp)
		return
	}
	u1 := uuid.Must(uuid.NewV4(), nil)
	dst := u1.String() + "-" + headerFile.Filename
	if err := c.SaveUploadedFile(headerFile, dst); err != nil {
		resp.Code = 202
		resp.Msg = "save file error"
		uploadResponse(c, resp)
		return
	}

	headerPara, err := c.MultipartForm()
	if err != nil {
		resp.Code = 203
		resp.Msg = fmt.Sprintf("get multipartForm error:%s", err)
		uploadResponse(c, resp)
		return
	}
	v := headerPara.Value["rect"]
	if len(v) == 0 {
		resp.Code = 204
		resp.Msg = fmt.Sprintf("rect not exist")
		uploadResponse(c, resp)
		return
	}
	rect := v[0]
	algoRect := AlgoRect{}
	if err := json.Unmarshal([]byte(rect), &algoRect); err != nil {
		resp.Code = 205
		resp.Msg = fmt.Sprintf("json unmarshal error %v", err)
		uploadResponse(c, resp)
		return
	}

	common.Log.Debugf("algoRect: %v", algoRect)
	err, rects := algoProcess(dst, algoRect)
	if err != nil {
		resp.Code = 206
		resp.Msg = fmt.Sprintf("%v", err)
		uploadResponse(c, resp)
		return
	}
	common.Log.Debugf("algo process succeed")
	// for k, v := range rects {
	// 	common.Log.Debugf("[%d]rect: %d %d %d %d", k, v.X, v.Y, v.Width, v.Height)
	// }

	resp.Rects = rects
	uploadResponse(c, resp)
}

func uploadResponse(c *gin.Context, resp *AlgoResponse) {
	if resp.Code != 200 {
		common.Log.Errorf("process failed: %v", resp)
	} else {
		common.Log.Debugf("process suceess: %v", resp)
	}
	respBuffer, err := json.Marshal(&resp)
	if err != nil {
		common.Log.Errorf("Marshal resp error: %v", resp)
		return
	}
	c.String(http.StatusOK, string(respBuffer))
}
