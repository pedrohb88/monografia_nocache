package payments

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"monografia/model"
	"net/http"
	"os"
)

type Payments interface {
	Create(ctx context.Context, amount float64) (int, error)
	GetByID(ctx context.Context, id int) (*model.Payment, error)
}

type payments struct {
	url string
}

func New() Payments {
	return &payments{
		url: os.Getenv("PAYMENTS_URL"),
	}
}

func (p *payments) Create(ctx context.Context, amount float64) (int, error) {

	testID, reqID, err := getHeaders(ctx)
	if err != nil {
		return 0, err
	}

	data := model.Payment{
		Amount: amount,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.url+"/payments", bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-test", testID)
	req.Header.Add("x-req", reqID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var payment model.Payment
	err = json.NewDecoder(resp.Body).Decode(&payment)
	return payment.ID, err
}

func (p *payments) GetByID(ctx context.Context, id int) (*model.Payment, error) {
	testID, reqID, err := getHeaders(ctx)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/payments/%d", p.url, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-test", testID)
	req.Header.Add("x-req", reqID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payment model.Payment
	err = json.NewDecoder(resp.Body).Decode(&payment)
	return &payment, err
}

func getHeaders(ctx context.Context) (string, string, error) {

	if os.Getenv("ENV") != "production" {
		return "", "", nil
	}

	testID := ctx.Value("x-test").(string)
	if testID == "" {
		return "", "", fmt.Errorf("missing x-test header")
	}

	reqID := ctx.Value("x-req").(string)
	if reqID == "" {
		return "", "", fmt.Errorf("missing x-req header")
	}
	return testID, reqID, nil
}
