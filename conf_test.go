package golib

import "testing"

const (
	HOST_NAME = "golib"
	IP        = "127.0.0.1"
	PORT      = 123
)

func TestInitConf(t *testing.T) {
	_, err := InitConf("./conf.json")

	if err != nil {
		t.Error("init conf failed:", err)
	}
}

func TestReadSimpleItem(t *testing.T) {
	conf, err := InitConf("./conf.json")

	if err != nil {
		t.Error("init conf failed:", err)
	}

	host_name, ok := conf["host_name"].(string)
	if !ok {
		t.Error("get host_name failed")
	} else if host_name != HOST_NAME {
		t.Error("wrong host_name:", host_name)
	}
}

func TestReadComplicatedItem(t *testing.T) {
	conf, err := InitConf("./conf.json")

	if err != nil {
		t.Error("init conf failed:", err)
	}

	address, ok := conf["address"].([]interface{})
	if !ok {
		t.Error("get address failed")
	}

	address_item, ok := address[0].(map[string]interface{})
	if !ok {
		t.Error("get address item failed")
	}

	ip, ok := address_item["ip"].(string)
	if !ok {
		t.Error("get ip failed")
	} else if ip != IP {
		t.Errorf("wrong ip, %v != %v", ip, IP)
	}

	port, ok := address_item["port"].(float64)
	if !ok {
		t.Error("get port failed")
	} else if port != PORT {
		t.Errorf("wrong port, %v != %v", port, PORT)
	}
}
