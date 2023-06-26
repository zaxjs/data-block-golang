package data_block

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"dario.cat/mergo"
	"github.com/bytedance/sonic"
)

// golang 泛型，纯粹就是废物
// 2023-6-26

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

func distDataBlock(resAll *map[string]Block, opt *Options) (*map[string]Block, error) {
	for _, item := range *resAll {
		if !opt.ShowSysField {
			// 不展示系统系统字段
			item.Slugs = ""
			item.Stage = ""
			item.IsMultipleGroup = nil
			item.AtUsers = nil
			item.BlockStatus = nil
			item.SysId = 0
			item.CreatedBy = ""
			item.CreatedAt = nil
			item.UpdatedBy = nil
			item.UpdatedAt = nil
			item.PublishedBy = nil
			item.PublishedAt = nil
			item.SpaceName = ""
			item.SpaceId = ""
			item.ModelCode = ""
			item.SyncAt = nil
		}
		if !opt.ShowGroupInfo {
			// 不展示组信息
			if len(item.BlockData) > 0 {
				tmpSub := []map[string]interface{}{}
				for _, sub := range item.BlockData {
					subMd := &BlockData{}
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
			}
		}
	}

	return resAll, nil
}

func distDataKv(resAll *map[string]Kv, opt *Options) (*map[string]Kv, error) {
	for _, item := range *resAll {
		if !opt.ShowSysField {
			// 不展示系统系统字段
			item.BlockStatus = nil
			item.SysId = 0
			item.CreatedBy = ""
			item.CreatedAt = nil
			item.UpdatedBy = nil
			item.UpdatedAt = nil
			item.PublishedBy = nil
			item.PublishedAt = nil
			item.Description = ""
			item.SyncAt = nil
		}
	}

	return resAll, nil
}

type DataBlockService struct {
	Options *Options
}

func fixBodyData[T Block | Kv](body []byte, opt *Options) (*map[string]interface{}, error) {
	newMap := make(map[string]interface{})
	if opt.KeyType == BT_BLOCK {
		md := &BaseResponseModel[map[string]Block]{}
		err := sonic.Unmarshal(body, &md)
		if err != nil {
			log.Println("Unmarshal err：", err)
			return nil, err
		}
		dt, err := distDataBlock(&md.Data, opt)
		if err != nil {
			return nil, err
		}
		for key, value := range *dt {
			newMap[key] = value
		}
	} else if opt.KeyType == BT_KV {
		md := &BaseResponseModel[map[string]Kv]{}
		err := sonic.Unmarshal(body, &md)
		if err != nil {
			log.Println("Unmarshal err：", err)
			return nil, err
		}
		dt, err := distDataKv(&md.Data, opt)
		if err != nil {
			return nil, err
		}
		for key, value := range *dt {
			newMap[key] = value
		}
	}
	return &newMap, nil
}

func New(opt Options) (*DataBlockService, error) {
	// 注入缺省

	if len(opt.Api) <= 0 {
		return nil, errors.New("api can not be empty")
	}
	if len(opt.Key) <= 0 {
		return nil, errors.New("key can not be empty")
	}
	myOpt := &Options{ShowSysField: false, ShowGroupInfo: false}
	mergo.Merge(myOpt, opt, mergo.WithOverride)

	svc := &DataBlockService{
		Options: myOpt,
	}
	return svc, nil
}

func (svc *DataBlockService) Get(codes []string, newOpt Options) (*map[string]interface{}, error) {
	if len(codes) <= 0 {
		return nil, errors.New("code can not be empty")
	}

	if len(newOpt.KeyType) <= 0 {
		return nil, errors.New("KeyType param lost")
	}

	opt := *svc.Options

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

	// 读取响应内容
	body, _ := ioutil.ReadAll(res.Body)

	if opt.KeyType == BT_BLOCK {
		return fixBodyData[Block](body, &opt)
	} else if opt.KeyType == BT_KV {
		return fixBodyData[Kv](body, &opt)
	}
	return nil, nil
}

// GetBlock
func (svc *DataBlockService) GetBlock(codes []string, newOpt *Options) (*map[string]interface{}, error) {
	opt := *svc.Options // 值类型
	if newOpt == nil || len(newOpt.KeyType) <= 0 {
		opt.KeyType = BT_BLOCK
	}
	if newOpt != nil {
		mergo.Merge(&opt, newOpt, mergo.WithOverride)
	}
	return svc.Get(codes, opt)
}

// GetKv
func (svc *DataBlockService) GetKv(codes []string, newOpt *Options) (*map[string]interface{}, error) {
	opt := *svc.Options // 值类型
	if newOpt == nil || len(newOpt.KeyType) <= 0 {
		opt.KeyType = BT_KV
	}
	if newOpt != nil {
		mergo.Merge(&opt, newOpt, mergo.WithOverride)
	}
	return svc.Get(codes, opt)
}
