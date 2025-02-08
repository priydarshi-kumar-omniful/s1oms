package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/omniful/go_commons/http"
	interservice_client "github.com/omniful/go_commons/interservice-client"
)

var Client *interservice_client.Client

type Hub struct {
	ID        string    `json:"_id,omitempty" gorm:"primaryKey"`
	TenantID  string    `json:"tenant_id" gorm:"not null"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func InitInterSrvClient() {
	config := interservice_client.Config{
		ServiceName: "user-service",
		BaseURL:     "http://localhost:8081/wms",
		Timeout:     5 * time.Second,
	}

	client, err := interservice_client.NewClientWithConfig(config)
	if err != nil {
		panic(err)
	}
	Client = client
	
}

func GetReq(ctx context.Context, userData interface{}, Url string) (interface{}, *interservice_client.Error) {
	request := &http.Request{Url: Url}

	// Ensure userData is passed as a pointer for unmarshaling
	respData := userData // Create a new variable to hold the response

	_, err := Client.Get(request, &respData) // Pass as pointer
	if err != nil {
		return nil, err
	}

	fmt.Println("Response:", respData) // Debugging output
	return respData, nil
}

func PostReq(ctx context.Context, userData interface{}, Url string, body interface{}) (interface{}, *interservice_client.Error) {
	request := &http.Request{
		Url:  Url,
		Body: body,
	}

	_, err := Client.Post(request, &userData)
	if err != nil {
		return nil, err
	}

	jsonData, _ := json.Marshal(userData)
	fmt.Println(string(jsonData))

	return &userData, nil
}