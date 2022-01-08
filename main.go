package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var upload = &cobra.Command{
	Use:   "upload",
	Short: "Run the ftgo client",
	Long:  "Upload files to the ftgo host",
	Run: func(cmd *cobra.Command, args []string) {

		ip, err := cmd.Flags().GetString("ip")
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}

		cmdUrl, err := cmd.Flags().GetString("url")
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}

		fileData, err := os.Open(file)
		if err != nil {

			log.Fatal(err)

		}

		var url string

		if cmdUrl == "" {
			url = "http://" + ip + ":" + port + "/"
		} else {
			url = cmdUrl
		}

		body := new(bytes.Buffer)

		mw := multipart.NewWriter(body)

		fileName := strings.Split(file, "\\")[len(strings.Split(file, "\\")) - 1]

		w , err := mw.CreateFormFile("file", fileName)

		if err != nil {
			log.Fatal(err)
		}

		if _, err = io.Copy(w, fileData); err != nil {
			log.Fatal(err)
		}

		mw.Close()

		req, err := http.NewRequest("POST", url, body)

		if err != nil {
			log.Fatal(err)
		}

		req.Header.Add("Content-Type", mw.FormDataContentType())

		res , err := http.DefaultClient.Do(req)

		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

	},
}

var host = &cobra.Command{
	Use:              "host",
	Short:            "Run the ftgo host",
	Long:             "Serve the directory to upload files to it",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()
		r.SetTrustedProxies(nil)

		port, err := cmd.Flags().GetString("port")

		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}

		src, err := cmd.Flags().GetString("src")

		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}

		r.POST("/", func(c *gin.Context) {

			file, err := c.FormFile("file")
			if err != nil {
				c.String(http.StatusBadRequest, "Bad request"+err.Error())
				return
			}

			dst := src + "\\" + file.Filename

			err = c.SaveUploadedFile(file, dst)

			if err != nil {
				c.String(http.StatusBadRequest, "File not readable"+err.Error())
				return
			}
		})

		go r.Run(":" + port)

		exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://127.0.0.1:4040").Start()

		e := exec.Command("ngrok", "http", port).Run()

		if e != nil {

			fmt.Println(e)
			os.Exit(1)

		}

	},
}

func main() {

	upload.Flags().StringP("ip", "i", "", "IP addres to connect to")
	upload.Flags().StringP("port", "p", "8080", "Port to connect to")
	upload.Flags().StringP("file", "f", "", "Path of the file to upload")
	upload.Flags().StringP("url", "u", "", "Full url to connect to")
	upload.MarkFlagRequired("file")

	initialPath, errPath := os.Getwd()

	if errPath != nil {
		fmt.Println("Error getting your current path")
		os.Exit(1)
	}

	host.Flags().StringP("src", "s", initialPath, "Source directory to upload files")
	host.Flags().StringP("port", "p", "8080", "Port to serve")

	var ftgoCmd = &cobra.Command{Use: "ftgo"}
	ftgoCmd.AddCommand(host)
	ftgoCmd.AddCommand(upload)

	ftgoCmd.Execute()

}