package email

import (
	"fmt"
	"github.com/dealmaker/shared/auth/model"
	"github.com/sendgrid/sendgrid-go"
	"os"
)


type WorkerInstance struct {
	client *sendgrid.Client
	FuncGetCredUser func(*model.CredUser) *model.CredUser
}

func (w WorkerInstance) Init() *WorkerInstance {
	key := os.Getenv("SENDGRID_API_KEY")
	fmt.Println(key)
	w.client = sendgrid.NewSendClient(key)
	return &w
}
