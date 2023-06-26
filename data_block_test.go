package data_block

import (
	"log"
	"testing"

	"github.com/bytedance/sonic"
)

const (
	api = "http://localhost:8089/data-block-service-api/v1/open"
	key = "Y2wwemk4aWtnMDAwMjA4bDQ4c3VrZzB5bA=="
)

func init() {
	// Do some init
	log.SetPrefix("[data-block]: ")
}

func TestNew(t *testing.T) {
	inst, err := New(Options{Api: "foo", Key: "bar"})
	if err != nil {
		log.Println(err)
	}
	if inst.Options.Api != "foo" {
		t.Errorf("Options.Api expected be 'foo'")
	}
	if inst.Options.Key != "bar" {
		t.Errorf("Options.Key expected be 'bar'")
	}
	if inst.Options.ShowGroupInfo != false {
		t.Errorf("Options.ShowGroupInfo expected be [false]")
	}
	if inst.Options.ShowSysField != false {
		t.Errorf("Options.ShowSysField expected be [false]")
	}
	if inst.Options.Ttl != "" {
		t.Errorf("Options.Ttl expected be ['']")
	}
	if inst.Options.KeyType != "" {
		t.Errorf("Options.KeyType expected be [']")
	}
}

func TestGet(t *testing.T) {
	CODES := []string{"TEST_BLOCK", "TEST_MISC"}
	inst, err := New(Options{Api: api, Key: key})
	if err != nil {
		log.Println("New error", err)
	}
	res, err := inst.Get(CODES, Options{KeyType: BT_BLOCK, ShowSysField: true})
	if err != nil {
		log.Println("Get error", err)
	}

	resB, _ := sonic.Marshal(res)
	md := map[string]Block{}
	sonic.Unmarshal(resB, &md)

	for k, v := range md {
		if *v.BlockCode != k {
			t.Errorf("BlockCode" + " not equal to the response of Key")
		}
	}
}

func TestGetBlock(t *testing.T) {
	CODES := []string{"TEST_BLOCK", "TEST_MISC"}
	inst, err := New(Options{Api: api, Key: key})
	if err != nil {
		log.Println("New error", err)
	}
	res, err := inst.GetBlock(CODES, nil)
	if err != nil {
		log.Println("GetBlock error", err)
	}
	resB, _ := sonic.Marshal(res)
	md := map[string]Block{}
	sonic.Unmarshal(resB, &md)

	for k, v := range md {
		if *v.BlockCode != k {
			t.Errorf("BlockCode fn1" + " not equal to the response of Key")
		}
	}

	res2, err := inst.GetBlock(CODES, &Options{ShowSysField: true, ShowGroupInfo: true})
	if err != nil {
		log.Println("GetBlock error", err)
	}

	resB2, _ := sonic.Marshal(res2)
	md2 := map[string]Block{}
	sonic.Unmarshal(resB2, &md2)

	for k, v := range md2 {
		if *v.BlockCode != k {
			t.Errorf("BlockCode fn2" + " not equal to the response of Key")
		}
	}
}

func TestGetKv(t *testing.T) {
	CODES := []string{"TEST_KEY", "WX_HOME_BANNER"}
	inst, err := New(Options{Api: api, Key: key})
	if err != nil {
		log.Println("New error", err)
	}
	res, err := inst.GetKv(CODES, nil)
	if err != nil {
		log.Println("GetKv error", err)
	}

	resB, _ := sonic.Marshal(res)
	md := map[string]Kv{}
	sonic.Unmarshal(resB, &md)

	for k, v := range md {
		if v.K != k {
			t.Errorf("BlockCode fn1" + " not equal to the response of Key")
		}
	}

	res2, err := inst.GetKv(CODES, &Options{ShowSysField: true, ShowGroupInfo: true})
	if err != nil {
		log.Println("GetKv error", err)
	}

	resB2, _ := sonic.Marshal(res2)
	md2 := map[string]Kv{}
	sonic.Unmarshal(resB2, &md2)

	for k, v := range md2 {
		if v.K != k {
			t.Errorf("BlockCode fn2" + " not equal to the response of Key")
		}
	}
}
