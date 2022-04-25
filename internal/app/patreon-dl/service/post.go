package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"patreon-dl/internal/app/patreon-dl/entity"
	"patreon-dl/internal/app/patreon-dl/util"
	"sync"
)

const (
	perm   os.FileMode = 644
	format string      = "2006-01-02T15:04:05.999999999-0700"
)

var (
	query = map[string]string{
		"include":                          "campaign,access_rules,attachments,audio,images,media,poll.choices,poll.current_user_responses.user,poll.current_user_responses.choice,poll.current_user_responses.poll,user,user_defined_tags,ti_checks",
		"fields[campaign]":                 "currency,show_audio_post_download_links,avatar_photo_url,earnings_visibility,is_nsfw,is_monthly,name,url",
		"fields[post]":                     "change_visibility_at,comment_count,content,current_user_can_comment,current_user_can_delete,current_user_can_view,current_user_has_liked,embed,image,is_paid,like_count,meta_image_url,min_cents_pledged_to_view,post_file,post_metadata,published_at,patreon_url,post_type,pledge_url,thumbnail_url,teaser_text,title,upgrade_url,url,was_posted_by_campaign_owner,has_ti_violation",
		"fields[post_tag]":                 "tag_type,value",
		"fields[user]":                     "image_url,full_name,url",
		"fields[access_rule]":              "access_rule_type,amount_cents",
		"fields[media]":                    "id,image_urls,download_url,metadata,file_name",
		"filter[campaign_id]":              "273633",
		"filter[contains_exclusive_posts]": "true",
		"filter[is_draft]":                 "false",
		"sort":                             "-published_at",
		"json-api-version":                 "1.0",
	}
	header = map[string][]string{
		"accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"},
		"accept-encoding":           {"gzip, deflate, br"},
		"accept-language":           {"zh-CN,zh;q=0.9,en-GB;q=0.8,en;q=0.7,zh-TW;q=0.6"},
		"cache-control":             {"no-cache"},
		"dnt":                       {"1"},
		"pragma":                    {"no-cache"},
		"sec-ch-ua":                 {"\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"100\", \"Google Chrome\";v=\"100\""},
		"sec-ch-ua-mobile":          {"?0"},
		"sec-ch-ua-platform":        {"Windows"},
		"sec-fetch-dest":            {"document"},
		"sec-fetch-mode":            {"navigate"},
		"sec-fetch-site":            {"none"},
		"sec-fetch-user":            {"?1"},
		"upgrade-insecure-requests": {"1"},
		"user-agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"},
	}
)

func init() {
	b, err := ioutil.ReadFile("cookies.txt")
	if err != nil {
		panic(err)
	}
	header["cookie"] = []string{string(b)}
}

func getResp() (entity.PatreonResp, error) {
	var re entity.PatreonResp
	params := url.Values{}
	for k, v := range query {
		params[k] = []string{v}
	}
	s := "https://www.patreon.com/api/posts?" + params.Encode()
	b, r, err := util.NewGetRequest(s).SetHeader(header).SetDefaultProxy().Exec()
	if err != nil {
		return re, err
	}
	if r.StatusCode != 200 {
		return re, errors.New(r.Status)
	}
	if err = json.Unmarshal(b, &re); err != nil {
		return re, err
	}
	return re, nil
}

func dl(info entity.DlInfo, wg *sync.WaitGroup) {
	b, response, err := util.NewGetRequest(info.Url).SetDefaultProxy().Exec()
	if err != nil {
		panic(err)
	}
	if response.StatusCode != 200 {
		panic(response.Status)
	}
	if err = ioutil.WriteFile(info.Path, b, perm); err != nil {
		panic(err)
	}
	log.Println(info.Path)
	defer wg.Done()
}

func DlPost() {
	resp, err := getResp()
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	for _, datum := range resp.Data {
		dirName := fmt.Sprintf("out/%s/%s", datum.Attributes.PublishedAt.Format("2006-01"), datum.Id)
		if err = os.MkdirAll(dirName, perm); err != nil {
			panic(err)
		}
		for _, md := range datum.Relationships.Media.Data {
			for _, inc := range resp.Included {
				if md.Id == inc.Id {
					wg.Add(1)
					go dl(entity.DlInfo{
						Url:  inc.Attributes.ImageUrls.Original,
						Path: dirName + "/" + inc.Attributes.FileName,
					}, &wg)
				}
			}
		}
		content := datum.Attributes.Title + "\n" + datum.Attributes.Content + "\n" + datum.Attributes.PublishedAt.Format(format)
		if err = ioutil.WriteFile(dirName+"/content.txt", []byte(content), perm); err != nil {
			panic(err)
		}
	}
	wg.Wait()
}
