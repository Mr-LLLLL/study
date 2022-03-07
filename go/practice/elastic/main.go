package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/olivere/elastic"
)

func main() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.0.124:9200"),
		elastic.SetMaxRetries(10),
	)
	if err != nil {
		log.Fatal(err)
	}

	info, code, err := client.Ping("http://192.168.0.124:9200").Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if code == 200 {
		log.Printf("connected to es cluster: %s ,version: %s", info.ClusterName, info.Version.Number)
	}

	bq := elastic.NewBoolQuery()
	bq.Must(elastic.NewTermQuery("goodsId", "1467740956156497920"))
	res, err := client.Search().Index("mk_mall_goods_1_current").Type("_doc").Query(bq).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[string]interface{})
	for _, v := range res.Hits.Hits {
		tmp := make(map[string]interface{})

		json.Unmarshal(*v.Source, &tmp)
		fmt.Println("name", tmp["name"])
		fmt.Println("remark", tmp["remark"])
		fmt.Println("id", tmp["id"])
		fmt.Println("resources", tmp["resources"])
		fmt.Println("companyId", tmp["companyId"])
		fmt.Println("goodsExt", tmp["goodsExt"])
	}

	m["name"] = "5"
	m["remark"] = "6"
	m["resources"] = GoodsResource{
		VideoCover:       []string{"hello"},
		Images:           make([]*GoodsImage, 0),
		Videos:           make([]*GoodsVideo, 0),
		DetailVideoCover: make([]string, 0),
		DetailImages:     make([]*GoodsImage, 0),
		DetailVideos:     make([]*GoodsVideo, 0),
	}
	m["goodsExt.standardPrice"] = 200

	sql := ""
	for k := range m {
		k2 := strings.ReplaceAll(k, ".", "_")
		m[k2] = m[k]
		sql += fmt.Sprintf("ctx._source.%s = params.%s;", k, k2)
	}
	script := elastic.NewScript(sql).Params(m)

	_, err = client.UpdateByQuery().Index("mk_mall_goods_1_current").Type("_doc").Query(bq).Script(script).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

type GoodsResource struct {
	VideoCover       []string      `bson:"videoCovers" json:"MarketPrice"`             // 头图视频封面
	Images           []*GoodsImage `bson:"images" json:"images"`                       // 头图图片
	Videos           []*GoodsVideo `bson:"videos" json:"videos"`                       // 头图视频
	DetailVideoCover []string      `bson:"detailVideoCovers" json:"detailVideoCovers"` // 详情视频封面
	DetailImages     []*GoodsImage `bson:"detailImages" json:"detailImages"`           //  详情图
	DetailVideos     []*GoodsVideo `bson:"detailVideos" json:"detailVideos"`           // 详情视频
}

// GoodsImage 商品图片
type GoodsImage struct {
	Name string `bson:"name" json:"name"` // 图片名称
	URL  string `bson:"url" json:"url"`   // 图片地址
}

// GoodsVideo 商品视频
type GoodsVideo struct {
	URL    string `bson:"url" json:"url"`       // 视频地址
	Name   string `bson:"name" json:"name"`     // 视频名称
	Size   int64  `bson:"size" json:"size"`     // 视频大小，单位：字节
	Second int64  `bson:"second" json:"second"` // 视频长度，单位：秒
}
