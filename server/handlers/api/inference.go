package api

import (
	"bytes"
	"context"
	"fmt"
	"image/png"

	"github.com/anthonynsimon/bild/transform"
	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/pkg/jsonutil"
	"github.com/medtune/beta-platform/pkg/session"
	mnist "github.com/medtune/capsules/capsules/mnist/v1"
	"github.com/vincent-petithory/dataurl"
)

func MnistRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip:")))
		return
	}

	infData := jsonutil.RunImageInference{}
	err := c.ShouldBindJSON(&infData)
	if err != nil {
		c.JSON(200, jsonutil.Fail(err))
		return
	}

	if infData.File == "" {
		dataURL, err := dataurl.DecodeString(infData.Image)
		if err != nil {
			c.JSON(200, jsonutil.Fail(err))
			return
		}

		img, err := png.Decode(bytes.NewReader(dataURL.Data))
		if err != nil {
			c.JSON(200, jsonutil.Fail(err))
			return
		}

		img = transform.Resize(img, 28, 28, transform.Linear)
		buf := new(bytes.Buffer)
		err = png.Encode(buf, img)
		if err != nil {
			c.JSON(200, jsonutil.Fail(err))
			return
		}

		send_s3 := buf.Bytes()
		req, err := mnist.PredictRequestFromBytes(send_s3)
		if err != nil {
			c.JSON(200, jsonutil.Fail(err))
			return
		}
		resp, err := mnist.RunInference(context.Background(), req)
		if err != nil {
			c.JSON(200, jsonutil.Fail(err))
			return
		}

		c.JSON(200, jsonutil.SuccessData(resp.Outputs))

	} else {

		resp, err := mnist.RunInferenceOnImagePath("./static/demos/mnist/images/" + infData.File)
		if err != nil {
			c.JSON(200, jsonutil.Fail(err))
			return
		}

		c.JSON(200, jsonutil.SuccessData(resp.Outputs))
	}
}

func InceptionImagenetRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip:")))
		return
	}
	c.JSON(200, jsonutil.Success())
}

func MuraRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip:")))
		return
	}
	c.JSON(200, jsonutil.Success())
}

func ChexrayRunInference(c *gin.Context) {
	if logged := session.GetLoginStatus(c); !logged {
		c.JSON(200, jsonutil.Fail(fmt.Errorf("access denied :rip:")))
		return
	}
	c.JSON(200, jsonutil.Success())
}