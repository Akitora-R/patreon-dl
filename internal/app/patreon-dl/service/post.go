package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"patreon-dl/internal/app/patreon-dl/config"
	"patreon-dl/internal/app/patreon-dl/entity"
	"patreon-dl/internal/app/patreon-dl/util"
	"sync"
	"sync/atomic"
)

var (
	postQuery = map[string][]string{
		"include":                          {"campaign,access_rules,attachments,audio,images,media,poll.choices,poll.current_user_responses.user,poll.current_user_responses.choice,poll.current_user_responses.poll,user,user_defined_tags,ti_checks"},
		"fields[campaign]":                 {"currency,show_audio_post_download_links,avatar_photo_url,earnings_visibility,is_nsfw,is_monthly,name,url"},
		"fields[post]":                     {"change_visibility_at,comment_count,content,current_user_can_comment,current_user_can_delete,current_user_can_view,current_user_has_liked,embed,image,is_paid,like_count,meta_image_url,min_cents_pledged_to_view,post_file,post_metadata,published_at,patreon_url,post_type,pledge_url,thumbnail_url,teaser_text,title,upgrade_url,url,was_posted_by_campaign_owner,has_ti_violation"},
		"fields[post_tag]":                 {"tag_type,value"},
		"fields[user]":                     {"image_url,full_name,url"},
		"fields[access_rule]":              {"access_rule_type,amount_cents"},
		"fields[media]":                    {"id,image_urls,download_url,metadata,file_name"},
		"filter[contains_exclusive_posts]": {"true"},
		"filter[is_draft]":                 {"false"},
		"sort":                             {"-published_at"},
		"json-api-version":                 {"1.0"},
	}
)

func getResp(campaign entity.Campaign, onEachPage func(entity.PatreonResp, entity.Campaign)) {
	var params url.Values = util.CopyMap(postQuery)
	params["filter[campaign_id]"] = []string{campaign.Id}
	s := "https://www.patreon.com/api/posts?" + params.Encode()
	for i := 1; ; i++ {
		var re entity.PatreonResp
		log.Println("第", i, "页")
		b, r, err := util.NewGetRequest(s).SetHeader(header).SetOptionalProxy().Exec()
		if err != nil {
			panic(err)
		}
		if r.StatusCode != 200 {
			panic(r.StatusCode)
		}
		if err = json.Unmarshal(b, &re); err != nil {
			panic(err)
		}
		onEachPage(re, campaign)
		s = re.Links.Next
		if s == "" {
			break
		}
	}
}

func dl(info entity.DlInfo) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("下载失败:", info.Url, info.Path)
			atomic.AddUint32(&config.Failed, 1)
		}
	}()
	if util.FileExists(info.Path) {
		log.Println(info.Path, "已存在，跳过")
		atomic.AddUint32(&config.Ignored, 1)
		return
	}
	b, response, err := util.NewGetRequest(info.Url).SetOptionalProxy().SetTimeout(60 * 1000).Exec()
	if err != nil {
		panic(err)
	}
	if response.StatusCode != 200 {
		panic(response.Status)
	}
	if err = ioutil.WriteFile(info.Path, b, perm); err != nil {
		panic(err)
	}
	atomic.AddUint32(&config.Succeed, 1)
	log.Println(info.Path)
}

func dlPage(resp entity.PatreonResp, campaign entity.Campaign) {
	wg := sync.WaitGroup{}
	for _, datum := range resp.Data {
		if !datum.Attributes.CurrentUserCanView {
			log.Println(datum.Attributes.Url, "当前用户无权限查看，跳过")
			atomic.AddUint32(&config.Ignored, 1)
			continue
		}
		dirName := fmt.Sprintf("out/%s/%s/%s", campaign.Name, datum.Attributes.PublishedAt.Format("2006-01"), datum.Id)
		if err := os.MkdirAll(dirName, perm); err != nil {
			panic(err)
		}
		for _, md := range datum.Relationships.Media.Data {
			for _, inc := range resp.Included {
				if md.Id == inc.Id {
					wg.Add(1)
					go func() {
						info := entity.DlInfo{
							Url:  inc.Attributes.ImageUrls.Original,
							Path: dirName + "/" + inc.Attributes.FileName,
						}
						dl(info)
						defer wg.Done()
					}()
					break
				}
			}
		}
		content := datum.Attributes.Title + "\n" + datum.Attributes.Content + "\n" + datum.Attributes.PublishedAt.Format(format)
		if err := ioutil.WriteFile(dirName+"/content.txt", []byte(content), perm); err != nil {
			panic(err)
		}
	}
	wg.Wait()
}

func DlAllPost(campaign entity.Campaign) {
	getResp(campaign, dlPage)
}
