package item_get

import (
	"gitee.com/fat_marmota/infra/log"
	"gitee.com/fat_marmota/streamline"
	"github.com/dealmaker/dal"
	"github.com/dealmaker/procedure/auth_db"
	"github.com/dealmaker/procedure/item_upload"
	"testing"
	"time"
)

func TestQueryItem(t *testing.T) {
	log.InitZapSugared(true, false, 1)
	dal.InitDatabaseClient("root:12345678@tcp(127.0.0.1:3306)/dealmaker?parseTime=true", nil, "mysql")
	auth_db.InitUserCredModel()
	item_upload.InitItemModel()
	item_upload.InitTagsModel()


	dataDomain := ItemFilter{
		Uploader:  0,
		Tags:      []string{"B"},
		BeginTime: time.Unix(0,0),
		EndTime:   time.Unix(0,0),
	}

	c := streamline.ConveyorBelt{
		DataDomain: &dataDomain,
		S:          &streamline.Streamline{Name: "test"},
		Ctx:        nil,
		Logger:     log.GlobalLogger,
		LogInfoGen: func(belt *streamline.ConveyorBelt) string {
			return belt.S.Name
		},
	}

	QueryItem(&c)
}
