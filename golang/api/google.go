package api

import (
	"context"

	"google.golang.org/api/cloudkms/v1"
)

type PublishApi interface {
	Connect()
	Publish(content string)
}

type GoogleApi struct {
	Template   string
	Dest       string
	KMSService *cloudkms.Service
	Flat       bool
}

func (x *GoogleApi) Connect() error {
	ctx := context.Background()
	kmsService, err := cloudkms.NewService(ctx)
	if err != nil {
		return err
	}
	x.KMSService = kmsService
	return nil
}

func (x GoogleApi) Publish(content string) {

}
