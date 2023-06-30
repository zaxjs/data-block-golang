package data_block

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"

	"dario.cat/mergo"
	"github.com/bytedance/sonic"
)

// golang 泛型，纯粹就是废物
// 2023-6-26
// @see https://github.com/golang/go/issues/48522

const APP_OPEN_KEY = `x-data-block-openkey`

type Options struct {
	Key           string     `json:"$id,omitempty"`           // API KEY
	Api           string     `json:"type,omitempty"`          // Request url
	ShowSysField  bool       `json:"showSysField,omitempty"`  // 展示系统字段
	ShowGroupInfo bool       `json:"showGroupInfo,omitempty"` // 展示组信息
	Ttl           string     `json:"ttl,omitempty"`           // 缓存时间，默认5s //  `${number}${'d' | 'h' | 'm' | 's'}`
	KeyType       BLOCK_TYPE `json:"keyType,omitempty"`
}

func init() {
	// Do some init
	log.SetPrefix("[data-block]: ")
}

func distData[T Block | Kv](resAll *map[string]T, opt Options) (*map[string]T, error) {
	for key, item := range *resAll {
		if !opt.ShowSysField {
			// Remove system fields
			v := reflect.ValueOf(&item).Elem()
			v.FieldByName("Slugs").Set(reflect.ValueOf(""))
			v.FieldByName("Stage").Set(reflect.ValueOf(nil))
			v.FieldByName("Description").Set(reflect.ValueOf(nil))
			v.FieldByName("IsMultipleGroup").Set(reflect.ValueOf(nil))
			v.FieldByName("AtUsers").Set(reflect.ValueOf(nil))
			v.FieldByName("BlockStatus").Set(reflect.ValueOf(nil))
			v.FieldByName("SysId").Set(reflect.ValueOf(0))
			v.FieldByName("CreatedBy").Set(reflect.ValueOf(""))
			v.FieldByName("CreatedAt").Set(reflect.ValueOf(nil))
			v.FieldByName("UpdatedBy").Set(reflect.ValueOf(nil))
			v.FieldByName("UpdatedAt").Set(reflect.ValueOf(nil))
			v.FieldByName("PublishedBy").Set(reflect.ValueOf(nil))
			v.FieldByName("PublishedAt").Set(reflect.ValueOf(nil))
			v.FieldByName("SpaceName").Set(reflect.ValueOf(""))
			v.FieldByName("SpaceId").Set(reflect.ValueOf(""))
			v.FieldByName("ModelCode").Set(reflect.ValueOf(""))
			v.FieldByName("SyncAt").Set(reflect.ValueOf(nil))

			(*resAll)[key] = item // Update the modified item in the map
		}
		if !opt.ShowGroupInfo {
			// Remove group field
			if len(item.BlockData) > 0 {
				tmpSub := []BlockData{}
				for _, sub := range item.BlockData {
					subMd := BlockData{}
					mySubByte, err := sonic.Marshal(sub)
					if err != nil {
						return nil, err
					}
					err = sonic.Unmarshal(mySubByte, &subMd)
					if err != nil {
						return nil, err
					}
					tmpSub = append(tmpSub, subMd.Data...)
				}
				item.BlockData = tmpSub
				(*resAll)[key] = item // Update the modified item in the map
			}
		}
	}

	return resAll, nil
}

func fixBodyData[T Block | Kv](body []byte, opt Options) (*map[string]T, error) {
	newMap := make(map[string]T)

	md := &BaseResponseModel[map[string]T]{}
	err := sonic.Unmarshal(body, &md)
	if err != nil {
		log.Println("Unmarshal err：", err)
		return nil, err
	}
	dt, err := distData(&md.Data, opt)
	if err != nil {
		return nil, err
	}
	for key, value := range *dt {
		newMap[key] = value
	}
	return &newMap, nil
}

type DataBlockService[T Block | Kv] struct {
	Options *Options
}

func New[T Block | Kv](opt Options) (*DataBlockService[T], error) {
	// inject default params

	if len(opt.Api) <= 0 {
		return nil, errors.New("api can not be empty")
	}
	if len(opt.Key) <= 0 {
		return nil, errors.New("key can not be empty")
	}
	myOpt := &Options{ShowSysField: false, ShowGroupInfo: false}
	mergo.Merge(myOpt, opt, mergo.WithOverride)

	svc := &DataBlockService[T]{
		Options: myOpt,
	}
	return svc, nil
}

func (svc *DataBlockService[T]) Get(codes []string, newOpt Options) (*map[string]T, error) {
	if len(codes) <= 0 {
		return nil, errors.New("code can not be empty")
	}

	if len(newOpt.KeyType) <= 0 {
		return nil, errors.New("KeyType param lost")
	}

	opt := *svc.Options // Get real value, prevent being polluted

	mergo.Merge(&opt, newOpt, mergo.WithOverride)

	key := opt.Key
	url := opt.Api + "/" + string(opt.KeyType) + "/" + strings.Join(codes, ",")
	method := "GET"
	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("Error make request:", err)
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	req.Header.Add(APP_OPEN_KEY, key)

	res, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return nil, err
	}
	defer res.Body.Close()

	// Get body from api response
	body, _ := ioutil.ReadAll(res.Body)

	if opt.KeyType == BT_BLOCK || opt.KeyType == BT_KV {
		return fixBodyData[T](body, opt)
	}

	// 展示原始信息
	md := &BaseResponseModel[map[string]T]{}
	err = sonic.Unmarshal(body, &md)
	if err != nil {
		log.Println("Unmarshal err：", err)
		return nil, err
	}
	return &md.Data, nil
}

// Deprecated: GetBlock is deprecated.
func (svc *DataBlockService[T]) GetBlock(codes []string, newOpt *Options) (*map[string]T, error) {
	return svc.Block(codes, newOpt)
}

// Block
func (svc *DataBlockService[T]) Block(codes []string, newOpt *Options) (*map[string]T, error) {
	opt := *svc.Options // Get real value, prevent being polluted
	if newOpt == nil || len(newOpt.KeyType) <= 0 {
		opt.KeyType = BT_BLOCK
	}
	if newOpt != nil {
		mergo.Merge(&opt, newOpt, mergo.WithOverride)
	}
	return svc.Get(codes, opt)
}

// Deprecated: GetKv is deprecated.
func (svc *DataBlockService[T]) GetKv(codes []string, newOpt *Options) (*map[string]T, error) {
	return svc.Kv(codes, newOpt)
}

// Kv
func (svc *DataBlockService[T]) Kv(codes []string, newOpt *Options) (*map[string]T, error) {
	opt := *svc.Options // Get real value, prevent being polluted
	if newOpt == nil || len(newOpt.KeyType) <= 0 {
		opt.KeyType = BT_KV
	}
	if newOpt != nil {
		mergo.Merge(&opt, newOpt, mergo.WithOverride)
	}
	return svc.Get(codes, opt)
}
