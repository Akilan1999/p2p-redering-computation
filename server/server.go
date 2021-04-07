package server

import (
	"encoding/json"
	"fmt"
	"git.sr.ht/~akilan1999/p2p-rendering-computation/p2p"
	"git.sr.ht/~akilan1999/p2p-rendering-computation/server/docker"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	//"fmt"
)

func Server() error{
	r := gin.Default()

	// Gets default information of the server
	r.GET("/server_info", func(c *gin.Context) {
		c.JSON(http.StatusOK, ServerInfo())
	})

	// Speed test with 50 mbps
	r.GET("/50", func(c *gin.Context){
		c.File("/etc/p2p-rendering/50.bin")
	})

	// Route build to do a speed test
	r.GET("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")

		// Upload the file to specific dst.
		// c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	//Gets Ip Table from server node
	r.POST("/IpTable", func(c *gin.Context) {

		// Variable to store IP table information
		var IPTable p2p.IpAddresses

		// Receive file from POST request
		body, err := c.FormFile("json")
		if err != nil {
			c.String(http.StatusOK, fmt.Sprint(err))
		}

		// Open file
		open, err := body.Open()
		if err != nil {
			c.String(http.StatusOK, fmt.Sprint(err))
		}

		// Open received file
		file, err := ioutil.ReadAll(open)
		if err != nil {
			c.String(http.StatusOK, fmt.Sprint(err))
		}

		json.Unmarshal(file,&IPTable)

		// Runs speed test to return only servers in the IP table pingable
		err = IPTable.SpeedTestUpdatedIPTable()
		if err != nil {
			c.String(http.StatusOK, fmt.Sprint(err))
		}

		// Reads IP addresses from ip table
		IpAddresses,err := p2p.ReadIpTable()
		if err != nil {
			c.String(http.StatusOK, fmt.Sprint(err))
		}

		c.JSON(http.StatusOK, IpAddresses)
	})

    // Starts docker container in server
	r.GET("/startcontainer", func(c *gin.Context) {

		resp, err := docker.BuildRunContainer()

		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}

		c.JSON(http.StatusOK, resp)
	})

	// Future feature
	/*r.GET("/create_vm/:virtualization", func(c *gin.Context) {
		virtualization := c.Param("virtualization")
		// Runs based on Preallocated VM size
		if virtualization == "docker" {
		  sshinfo,err := docker.RunVM()
		 if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		 }
		 if sshinfo != nil {
            c.JSON(http.StatusOK, sshinfo)
		 }

		} else {
			c.String(200,"virtualization tool not selected")
		}
	})*/

	// Port running on
	err := r.Run(":8088")
	if err != nil {
		return err
	}

	return nil
}
