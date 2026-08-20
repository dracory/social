package social

import (
	"testing"
)

func TestAllGetShareUrls(t *testing.T) {
	params := ShareLinksParams{
		URL:          "https://example.com/page",
		Title:        "Sample Title",
		Description:  "Sample Description",
		ImageURL:     "https://example.com/image.jpg",
		Via:          "samplevia",
		HashTags:     "tag1,tag2",
		EmailAddress: "test@example.com",
		PhoneNumber:  "+1234567890",
		UserID:       "user123",
		CCEmail:      "cc@example.com",
		BCCEmail:     "bcc@example.com",
	}

	s := New(params)

	tests := []struct {
		name     string
		got      string
		expected string
	}{
		{"Blogger", s.GetBloggerShareUrl(), "https://www.blogger.com/blog-this.g?"},
		{"Bluesky", s.GetBlueskyShareUrl(), "https://bsky.app/intent/compose?"},
		{"Diaspora", s.GetDiasporaShareUrl(), "https://share.diasporafoundation.org/?"},
		{"Discord", s.GetDiscordShareUrl(), "https://discord.com"},
		{"Douban", s.GetDoubanShareUrl(), "http://www.douban.com/recommend/?"},
		{"Email", s.GetEmailShareUrl(), "mailto:test@example.com?"},
		{"Evernote", s.GetEvernoteShareUrl(), "https://www.evernote.com/clip.action?"},
		{"Facebook", s.GetFacebookShareUrl(), "https://www.facebook.com/sharer.php?"},
		{"Flipboard", s.GetFlipboardShareUrl(), "https://share.flipboard.com/bookmarklet/popout?"},
		{"GitHub", s.GetGitHubShareUrl(), "https://github.com"},
		{"Gmail", s.GetGmailShareUrl(), "https://mail.google.com/mail/?"},
		{"GoogleBookmarks", s.GetGoogleBookmarksShareUrl(), "https://www.google.com/bookmarks/mark?"},
		{"Instagram", s.GetInstagramShareUrl(), "https://www.instagram.com"},
		{"Instapaper", s.GetInstapaperShareUrl(), "http://www.instapaper.com/edit?"},
		{"LineMe", s.GetLineMeShareUrl(), "https://lineit.line.me/share/ui?"},
		{"LinkedIn", s.GetLinkedInShareUrl(), "https://www.linkedin.com/sharing/share-offsite/?"},
		{"LiveJournal", s.GetLiveJournalShareUrl(), "http://www.livejournal.com/update.bml?"},
		{"HackerNews", s.GetHackerNewsShareUrl(), "https://news.ycombinator.com/submitlink?"},
		{"Mastodon", s.GetMastodonShareUrl(), "https://mastodon.social/share?"},
		{"Medium", s.GetMediumShareUrl(), "https://medium.com"},
		{"OkRu", s.GetOkRuShareUrl(), "https://connect.ok.ru/dk?"},
		{"Pinterest", s.GetPinterestShareUrl(), "https://pinterest.com/pin/create/button/?"},
		{"Pocket", s.GetPocketShareUrl(), "https://getpocket.com/save?"},
		{"QZone", s.GetQZoneShareUrl(), "http://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?"},
		{"Reddit", s.GetRedditShareUrl(), "https://reddit.com/submit?"},
		{"Renren", s.GetRenrenShareUrl(), "http://widget.renren.com/dialog/share?"},
		{"Skype", s.GetSkypeShareUrl(), "https://web.skype.com/share?"},
		{"SkypeCall", s.GetSkypeCallShareUrl(), "skype:user123?call"},
		{"SkypeChat", s.GetSkypeChatShareUrl(), "skype:user123?chat"},
		{"SMS", s.GetSMSShareUrl(), "sms:+1234567890?"},
		{"Snapchat", s.GetSnapchatShareUrl(), "https://www.snapchat.com"},
		{"Telegram", s.GetTelegramShareUrl(), "https://t.me/share/url?"},
		{"Telephone", s.GetTelephoneShareUrl(), "tel:+1234567890"},
		{"Threema", s.GetThreemaShareUrl(), "threema://compose?"},
		{"Threads", s.GetThreadsShareUrl(), "https://threads.net/intent/post?"},
		{"TikTok", s.GetTikTokShareUrl(), "https://www.tiktok.com"},
		{"Tumblr", s.GetTumblrShareUrl(), "https://www.tumblr.com/widgets/share/tool?"},
		{"Twitter", s.GetTwitterShareUrl(), "https://x.com/intent/tweet?"},
		{"Viber", s.GetViberShareUrl(), "viber://forward?"},
		{"VK", s.GetVKShareUrl(), "http://vk.com/share.php?"},
		{"Weibo", s.GetWeiboShareUrl(), "http://service.weibo.com/share/share.php?"},
		{"WhatsApp", s.GetWhatsAppShareUrl(), "https://api.whatsapp.com/send?"},
		{"Xing", s.GetXingShareUrl(), "https://www.xing.com/spi/shares/new?"},
		{"Yahoo", s.GetYahooShareUrl(), "http://compose.mail.yahoo.com/?"},
		{"YouTube", s.GetYouTubeShareUrl(), "https://www.youtube.com"},
		{"Print", s.GetPrintShareUrl(), "javascript:window.print();"},
		{"CopyLink", s.GetCopyLinkShareUrl(), "javascript:navigator.clipboard.writeText('https://example.com/page');"},
		{"NativeShare", s.GetNativeShareShareUrl(), "javascript:if(navigator.share){navigator.share({title:'Sample Title',url:'https://example.com/page'});}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.got) < len(tt.expected) || tt.got[:len(tt.expected)] != tt.expected {
				t.Errorf("Get%sShareUrl() = %q, expected prefix %q", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestEmailDefaultBody(t *testing.T) {
	// Without description, body should default to URL
	s := New(ShareLinksParams{
		URL:          "https://example.com/page",
		Title:        "Sample Title",
		EmailAddress: "test@example.com",
	})

	url := s.GetEmailShareUrl()
	if url != "mailto:test@example.com?body=https%3A%2F%2Fexample.com%2Fpage&subject=Sample+Title" &&
		url != "mailto:test@example.com?subject=Sample+Title&body=https%3A%2F%2Fexample.com%2Fpage" {
		t.Errorf("Unexpected email URL format: %s", url)
	}
}
