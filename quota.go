package chatnio

import (
	"fmt"
	"github.com/Deeptrain-Community/chatnio-api-go/utils"
)

type Quota struct {
	Status bool    `json:"status"`
	Quota  float32 `json:"quota"`
}

type QuotaBuy struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

type Package struct {
	Status bool `json:"status"`
	Data   struct {
		Cert     bool `json:"cert"`
		Teenager bool `json:"teenager"`
	} `json:"data"`
}

type Subscription struct {
	Status       bool  `json:"status"`
	IsSubscribed bool  `json:"is_subscribed"`
	Expired      int64 `json:"expired"`
}

type Subscribe struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

func (i *Instance) GetQuota() (float32, error) {
	quota, err := utils.GetForm[Quota](i.Mix("/quota"), i.GetHeaders())

	if err != nil {
		return 0., err
	} else if !quota.Status {
		return 0., fmt.Errorf("quota status is false")
	}

	return quota.Quota, nil
}

func (i *Instance) BuyQuota(quota int) error {
	data, err := utils.PostForm[QuotaBuy](i.Mix("/buy"), i.GetHeaders(), map[string]interface{}{
		"quota": quota,
	})

	if err != nil {
		return err
	} else if !data.Status {
		return fmt.Errorf(data.Error)
	}

	return nil
}

func (i *Instance) GetPackage() (*Package, error) {
	data, err := utils.GetForm[Package](i.Mix("/package"), i.GetHeaders())

	if err != nil {
		return nil, err
	} else if !data.Status {
		return nil, fmt.Errorf("package status is false")
	}

	return data, nil
}

func (i *Instance) GetSubscription() (*Subscription, error) {
	data, err := utils.GetForm[Subscription](i.Mix("/subscription"), i.GetHeaders())

	if err != nil {
		return nil, err
	} else if !data.Status {
		return nil, fmt.Errorf("subscription status is false")
	}

	return data, nil
}

func (i *Instance) Subscribe(level int, month int) error {
	data, err := utils.PostForm[Subscribe](i.Mix("/subscribe"), i.GetHeaders(), map[string]interface{}{
		"level": level,
		"month": month,
	})

	if err != nil {
		return err
	} else if !data.Status {
		return fmt.Errorf(data.Error)
	}

	return nil
}
