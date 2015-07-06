package golib

import (
	"encoding/json"
	"errors"
	"os"
)

func InitConf(conf_file string) (conf map[string]interface{}, err error) {
	fd, err := os.Open(conf_file)
	if err != nil {
		return nil, errors.New("open conf file: " + conf_file + " failed, error info: " + err.Error())
	}
	defer fd.Close()

	var buf [1024 * 256]byte
	read_len, err := fd.Read(buf[:])
	if err != nil {
		return nil, errors.New("read conf file: " + conf_file + " failed, error info: " + err.Error())
	}

	json_data := buf[:read_len]
	var r interface{}
	if err := json.Unmarshal(json_data, &r); err != nil {
		return nil, errors.New("unmarshal json: " + string(json_data) + " failed, error info: " + err.Error())
	}

	conf = r.(map[string]interface{})
	return conf, nil
}
