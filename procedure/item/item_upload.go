package item

import (
	"bytes"
	"encoding/base64"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/dealmaker/procedure/item/model"
	model2 "github.com/dealmaker/shared/auth/model"
	"github.com/itzmeerkat/streamline"
	"image"
	"math"
	"net/http"
	"strings"
	"time"
)

type InsertItemInterface interface {
	GetItem() *model.Item
	GetJwtAuth() *model2.JwtAuth
}

const DEFAULT_IMG = ``

func (w *WorkerInstance) InsertItem(c *streamline.ConveyorBelt) int {
	data := c.DataDomain.(InsertItemInterface).GetItem()
	jwtData := c.DataDomain.(InsertItemInterface).GetJwtAuth()
	c.Debugw(
		"desc",data.Description,
		"title", data.Title,
		"tags", data.Tags)

	data.Uploader = jwtData.Uid
	data.UpdateTime = time.Now().UnixNano() / 1000

	if len(data.Images) <= 0 {
		data.Images = append(data.Images, DEFAULT_IMG)
	}

	for _, v := range data.Images {
		res, err := createThumbnail(v)
		if err != nil {
			c.Errorw("create thumbnail failed", err)
			return http.StatusInternalServerError
		}
		data.Thumbnails = append(data.Thumbnails, res)
	}


	objid, err := w.FuncInsertItem(c.Ctx, data)
	if err != nil {
		c.Errorw("Insert item", err)
		return http.StatusForbidden
	}
	data.ObjId = objid
	return http.StatusOK
}

const (
	THUMBNAIL_WIDTH = 286
	THUMBNAIL_HEIGHT = 180
)

func createThumbnail(raw string) (string, error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(raw))
	m, _, err := image.Decode(reader)
	if err != nil {
		return "", err
	}
	width := float64(m.Bounds().Max.X)
	height := float64(m.Bounds().Max.Y)
	wr := width / THUMBNAIL_WIDTH
	hr := height / THUMBNAIL_HEIGHT

	r := math.Min(wr,hr)

	nw := int(math.Round(width * r))
	nh := int(math.Round(height * r))
	resized := transform.Resize(m,nw,nh,transform.Linear)

	buf := new(bytes.Buffer)
	encoder := imgio.JPEGEncoder(50)
	err = encoder(buf, resized)
	if err != nil {
		return "", err
	}
	b64res := new(bytes.Buffer)
	b64encoder := base64.NewEncoder(base64.StdEncoding,b64res)
	b64encoder.Write(buf.Bytes())
	b64encoder.Close()

	return b64res.String(), nil
}