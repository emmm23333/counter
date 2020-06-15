package cmd

import (
	"bytes"
	"counter/common"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting server ...")
		upolad()
	},
}

var uploadFile string

func init() {
	rootCmd.AddCommand(clientCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	clientCmd.PersistentFlags().StringVar(&uploadFile, "file", "", "specify file to upload")
	// runCmd.PersistentFlags().StringVar(&uploadFile, "file", "", "specify file to upload")
}

func upolad() {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	bodyWriter.WriteField("rect", `{"x":0,"y":0,"width":480,"height":960}`)
	name := filepath.Base(uploadFile)
	// fmt.Printf("file path:%s name:%s\n", uploadFile, name)
	fileWriter, _ := bodyWriter.CreateFormFile("file", name)

	file, _ := os.Open(uploadFile)
	defer file.Close()

	io.Copy(fileWriter, file)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post("http://localhost:8011/upload", contentType, bodyBuffer)
	if err != nil {
		common.Log.Errorf("post error:%v", err)
	}
	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	common.Log.Debugf("resp body: %s", string(b))
	common.Log.Debugf("resp status: %v", resp.Status)
}
