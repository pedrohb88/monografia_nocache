package payments

import (
	"bytes"
	"encoding/json"
	"fmt"
	"monografia/model"
	"net/http"
	"os"
)

type Payments interface {
	Create(amount float64) (int, error)
	GetByID(id int) (*model.Payment, error)
}

type payments struct {
	url string
}

func New() Payments {
	return &payments{
		url: os.Getenv("PAYMENTS_URL"),
	}
}

func (p *payments) Create(amount float64) (int, error) {

	data := model.Payment{
		Amount: amount,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post(p.url+"/payments", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, nil
	}
	defer resp.Body.Close()

	var payment model.Payment
	err = json.NewDecoder(resp.Body).Decode(&payment)
	return payment.ID, err
}

func (p *payments) GetByID(id int) (*model.Payment, error) {

	url := fmt.Sprintf("%s/payments/%d", p.url, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var payment model.Payment
	err = json.NewDecoder(resp.Body).Decode(&payment)
	return &payment, err
}
