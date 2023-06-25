package data_block

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"dario.cat/mergo"
	"github.com/bytedance/sonic"
)

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
	log.SetFlags(0)
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
			item.SysId = ""
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

type DataBlockService struct {
	Options *Options
}

func New(opt Options) (*DataBlockService, error) {
	myOpt := &Options{ShowSysField: false, ShowGroupInfo: false}
	mergo.Merge(&myOpt, opt, mergo.WithOverride)

	svc := &DataBlockService{
		Options: myOpt,
	}
	return svc, nil
}

func (svc *DataBlockService) Get(codes []string, newOpt *Options) (*map[string]Block, error) {
	if len(codes) <= 0 {
		return nil, errors.New("Code不能为空")
	}

	opt := svc.Options

	mergo.Merge(&opt, newOpt, mergo.WithOverride)

	key := opt.Key
	url := opt.Api + "/" + string(opt.KeyType) + "/" + strings.Join(codes, ",")
	method := "GET"
	client := http.Client{}
	req, _ := http.NewRequest(method, url, nil)
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"APP_OPEN_KEY": {key},
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer res.Body.Close()

	md := &map[string]Block{}

	// 读取响应内容
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("读取响应内容时出错：", err)
		return nil, err
	}

	err = sonic.Unmarshal(body, &md)
	if err != nil {
		fmt.Println("unmarshal时出错：", err)
		return nil, err
	}

	if opt.KeyType == BT_BLOCK {
		return distDataBlock(md, opt)
	}
	return md, nil
}

func (svc *DataBlockService) GetBlock(codes []string, opt *Options) (*map[string]Block, error) {
	opt.KeyType = BT_BLOCK
	return svc.Get(codes, opt)
}

func (svc *DataBlockService) GetKv(codes []string, opt *Options) (*map[string]Block, error) {
	opt.KeyType = BT_KV
	return svc.Get(codes, opt)
}
