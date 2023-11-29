package chatnio

import (
	"fmt"
	"github.com/Deeptrain-Community/chatnio-api-go/utils"
	"testing"
)

func TestInstance_GetQuota(t *testing.T) {
	if !instance.IsAuthenticated() {
		return
	}

	quota, err := instance.GetQuota()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(fmt.Sprintf("[quota] get quota: %f", quota))
}

func TestInstance_BuyQuota(t *testing.T) {
	if true {
		return
	} // buy quota is disabled because it's not free

	if !instance.IsAuthenticated() {
		return
	}

	err := instance.BuyQuota(1)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(fmt.Sprintf("[quota] buy quota: status %t", err == nil))
}

func TestInstance_GetPackage(t *testing.T) {
	if !instance.IsAuthenticated() {
		return
	}

	pkg, err := instance.GetPackage()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(fmt.Sprintf("[quota] get package: %s", utils.MarshalForm(*pkg)))
}

func TestInstance_GetSubscription(t *testing.T) {
	if !instance.IsAuthenticated() {
		return
	}

	sub, err := instance.GetSubscription()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(fmt.Sprintf("[quota] get subscription: %s", utils.MarshalForm(*sub)))
}

func TestInstance_Subscribe(t *testing.T) {
	if true {
		return
	} // subscribe is disabled because it's not free
	if !instance.IsAuthenticated() {
		return
	}

	err := instance.Subscribe(1, 1)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(fmt.Sprintf("[quota] subscribe: status %t", err == nil))
}
