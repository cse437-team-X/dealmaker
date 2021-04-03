package email

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)


type WorkerInstance struct {
	client *sendgrid.Client
}

func (w WorkerInstance) Init() *WorkerInstance {
	key := os.Getenv("SENDGRID_API_KEY")
	fmt.Println(key)
	w.client = sendgrid.NewSendClient(key)
	return &w
}
