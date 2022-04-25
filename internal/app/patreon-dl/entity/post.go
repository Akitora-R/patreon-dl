package entity

import "time"

type PatreonResp struct {
	Data     []PostData `json:"data"`
	Included []Include  `json:"included"`
	Links    struct {
		Next string `json:"next"`
	} `json:"links"`
	Meta struct {
		Pagination struct {
			Cursors struct {
				Next string `json:"next"`
			} `json:"cursors"`
			Total int `json:"total"`
		} `json:"pagination"`
	} `json:"meta"`
}

type Include struct {
	Attributes    IncludeAttribute `json:"attributes"`
	Id            string           `json:"id"`
	Relationships struct {
		Campaign struct {
			Data struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"campaign,omitempty"`
		Creator struct {
			Data struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"creator,omitempty"`
		Goals struct {
			Data []struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"goals,omitempty"`
		Rewards struct {
			Data []struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"rewards,omitempty"`
		Post struct {
			Data struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"post,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

type IncludeAttribute struct {
	FullName                   string  `json:"full_name,omitempty"`
	ImageUrl                   string  `json:"image_url,omitempty"`
	Url                        *string `json:"url,omitempty"`
	AvatarPhotoUrl             string  `json:"avatar_photo_url,omitempty"`
	Currency                   string  `json:"currency,omitempty"`
	EarningsVisibility         string  `json:"earnings_visibility,omitempty"`
	IsMonthly                  bool    `json:"is_monthly,omitempty"`
	IsNsfw                     bool    `json:"is_nsfw,omitempty"`
	Name                       string  `json:"name,omitempty"`
	ShowAudioPostDownloadLinks bool    `json:"show_audio_post_download_links,omitempty"`
	DownloadUrl                string  `json:"download_url,omitempty"`
	FileName                   string  `json:"file_name,omitempty"`
	ImageUrls                  struct {
		Default   string `json:"default"`
		Original  string `json:"original"`
		Thumbnail string `json:"thumbnail"`
	} `json:"image_urls,omitempty"`
	Metadata struct {
		Dimensions struct {
			H int `json:"h"`
			W int `json:"w"`
		} `json:"dimensions"`
	} `json:"metadata,omitempty"`
	AccessRuleType       string     `json:"access_rule_type,omitempty"`
	AmountCents          *int       `json:"amount_cents,omitempty"`
	PostCount            int        `json:"post_count,omitempty"`
	TagType              string     `json:"tag_type,omitempty"`
	Value                string     `json:"value,omitempty"`
	Amount               int        `json:"amount,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	Description          string     `json:"description,omitempty"`
	PatronCurrency       string     `json:"patron_currency,omitempty"`
	Remaining            *int       `json:"remaining,omitempty"`
	RequiresShipping     bool       `json:"requires_shipping,omitempty"`
	UserLimit            any        `json:"user_limit"`
	DiscordRoleIds       any        `json:"discord_role_ids"`
	EditedAt             time.Time  `json:"edited_at,omitempty"`
	PatronAmountCents    int        `json:"patron_amount_cents,omitempty"`
	PatronCount          int        `json:"patron_count,omitempty"`
	Published            bool       `json:"published,omitempty"`
	PublishedAt          time.Time  `json:"published_at,omitempty"`
	Title                string     `json:"title,omitempty"`
	UnpublishedAt        any        `json:"unpublished_at"`
	WelcomeMessage       any        `json:"welcome_message"`
	WelcomeMessageUnsafe any        `json:"welcome_message_unsafe"`
	WelcomeVideoEmbed    any        `json:"welcome_video_embed"`
	WelcomeVideoUrl      any        `json:"welcome_video_url"`
	CompletedPercentage  int        `json:"completed_percentage,omitempty"`
	ReachedAt            *time.Time `json:"reached_at,omitempty"`
}

type Relationship struct {
	AccessRules struct {
		Data []struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"data"`
	} `json:"access_rules"`
	Attachments struct {
		Data []struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"data"`
	} `json:"attachments"`
	Audio struct {
		Data any `json:"data"`
	} `json:"audio"`
	Campaign struct {
		Data struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"data"`
		Links struct {
			Related string `json:"related"`
		} `json:"links"`
	} `json:"campaign"`
	Images struct {
		Data []struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"data"`
	} `json:"images"`
	Media struct {
		Data []struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"data"`
	} `json:"media"`
	Poll struct {
		Data any `json:"data"`
	} `json:"poll"`
	TiChecks struct {
		Data []any `json:"data"`
	} `json:"ti_checks"`
	User struct {
		Data struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"data"`
		Links struct {
			Related string `json:"related"`
		} `json:"links"`
	} `json:"user"`
	UserDefinedTags struct {
		Data []struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"data"`
	} `json:"user_defined_tags"`
}

type PostData struct {
	Attributes struct {
		ChangeVisibilityAt    any    `json:"change_visibility_at"`
		CommentCount          int    `json:"comment_count"`
		Content               string `json:"content"`
		CurrentUserCanComment bool   `json:"current_user_can_comment"`
		CurrentUserCanDelete  bool   `json:"current_user_can_delete"`
		CurrentUserCanView    bool   `json:"current_user_can_view"`
		CurrentUserHasLiked   bool   `json:"current_user_has_liked"`
		Embed                 any    `json:"embed"`
		HasTiViolation        bool   `json:"has_ti_violation"`
		Image                 *struct {
			Height   int    `json:"height"`
			LargeUrl string `json:"large_url"`
			ThumbUrl string `json:"thumb_url"`
			Url      string `json:"url"`
			Width    int    `json:"width"`
		} `json:"image"`
		IsPaid                bool   `json:"is_paid"`
		LikeCount             int    `json:"like_count"`
		MetaImageUrl          string `json:"meta_image_url"`
		MinCentsPledgedToView *int   `json:"min_cents_pledged_to_view"`
		PatreonUrl            string `json:"patreon_url"`
		PledgeUrl             string `json:"pledge_url"`
		PostFile              *struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"post_file"`
		PostMetadata *struct {
			ImageOrder []string `json:"image_order"`
		} `json:"post_metadata"`
		PostType                 string    `json:"post_type"`
		PublishedAt              time.Time `json:"published_at"`
		TeaserText               string    `json:"teaser_text"`
		Title                    string    `json:"title"`
		UpgradeUrl               string    `json:"upgrade_url"`
		Url                      string    `json:"url"`
		WasPostedByCampaignOwner bool      `json:"was_posted_by_campaign_owner"`
	} `json:"attributes"`
	Id            string       `json:"id"`
	Relationships Relationship `json:"relationships"`
	Type          string       `json:"type"`
}
