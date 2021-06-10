package fetcher

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hi20160616/exhtml"
	"github.com/hi20160616/gears"
	"github.com/hi20160616/ms-ucpnz/configs"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Article struct {
	Id            string
	Title         string
	Content       string
	WebsiteId     string
	WebsiteDomain string
	WebsiteTitle  string
	UpdateTime    *timestamppb.Timestamp
	U             *url.URL
	raw           []byte
	doc           *html.Node
}

func NewArticle() *Article {
	return &Article{
		WebsiteDomain: configs.Data.MS.Domain,
		WebsiteTitle:  configs.Data.MS.Title,
		WebsiteId:     fmt.Sprintf("%x", md5.Sum([]byte(configs.Data.MS.Domain))),
	}
}

// List get all articles from database
func (a *Article) List() ([]*Article, error) {
	return load()
}

// Get read database and return the data by rawurl.
func (a *Article) Get(id string) (*Article, error) {
	as, err := load()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		if a.Id == id {
			return a, nil
		}
	}
	return nil, fmt.Errorf("[%s] no article with id: %s, url: %s",
		configs.Data.MS.Title, id, a.U.String())
}

func (a *Article) Search(keyword ...string) ([]*Article, error) {
	as, err := load()
	if err != nil {
		return nil, err
	}

	as2 := []*Article{}
	for _, a := range as {
		for _, v := range keyword {
			v = strings.ToLower(strings.TrimSpace(v))
			switch {
			case a.Id == v:
				as2 = append(as2, a)
			case a.WebsiteId == v:
				as2 = append(as2, a)
			case strings.Contains(strings.ToLower(a.Title), v):
				as2 = append(as2, a)
			case strings.Contains(strings.ToLower(a.Content), v):
				as2 = append(as2, a)
			case strings.Contains(strings.ToLower(a.WebsiteDomain), v):
				as2 = append(as2, a)
			case strings.Contains(strings.ToLower(a.WebsiteTitle), v):
				as2 = append(as2, a)
			}
		}
	}
	return as2, nil
}

type ByUpdateTime []*Article

func (u ByUpdateTime) Len() int      { return len(u) }
func (u ByUpdateTime) Swap(i, j int) { u[i], u[j] = u[j], u[i] }
func (u ByUpdateTime) Less(i, j int) bool {
	return u[i].UpdateTime.AsTime().Before(u[j].UpdateTime.AsTime())
}

var timeout = func() time.Duration {
	t, err := time.ParseDuration(configs.Data.MS.Timeout)
	if err != nil {
		log.Printf("[%s] timeout init error: %v", configs.Data.MS.Title, err)
		return time.Duration(1 * time.Minute)
	}
	return t
}()

// fetchArticle fetch article by rawurl
func (a *Article) fetchArticle(rawurl string) (*Article, error) {
	var err error
	a.U, err = url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	// Dail
	a.raw, a.doc, err = exhtml.GetRawAndDoc(a.U, timeout)
	if err != nil {
		return nil, err
	}

	a.Id = fmt.Sprintf("%x", md5.Sum([]byte(rawurl)))

	a.Title, err = a.fetchTitle()
	if err != nil {
		return nil, err
	}

	a.UpdateTime, err = a.fetchUpdateTime()
	if err != nil {
		return nil, err
	}

	// filter work
	if a, err = a.filter(3); errors.Is(err, ErrTimeOverDays) {
		return nil, err
	}

	// content should be the last step to fetch
	a.Content, err = a.fetchContent()
	if err != nil {
		return nil, err
	}

	a.Content, err = a.fmtContent(a.Content)
	if err != nil {
		return nil, err
	}
	return a, nil

}

func (a *Article) fetchTitle() (string, error) {
	n := exhtml.ElementsByTag(a.doc, "title")
	if n == nil {
		return "", fmt.Errorf("[%s] getTitle error, there is no element <title>", configs.Data.MS.Title)
	}
	title := n[0].FirstChild.Data
	title = strings.TrimSpace(title)
	gears.ReplaceIllegalChar(&title)
	return title, nil
}

func (a *Article) fetchUpdateTime() (*timestamppb.Timestamp, error) {
	if a.doc == nil {
		return nil, errors.Errorf("[%s] fetchUpdateTime: doc is nil: %s", configs.Data.MS.Title, a.U.String())
	}
	doc := exhtml.ElementsByTagAndClass(a.doc, "span", "td-post-date")
	d := []string{}
	if doc == nil {
		return nil, fmt.Errorf("[%s] fetchUpdateTime extract nothing: %s",
			configs.Data.MS.Title, a.U.String())
	}
	//focus on node like "<span class="td-post-date"><time class="entry-date updated td-module-date" datetime="2020-11-05T13:30:02+00:00" >2020-11-05</time></span>"
	if doc[0].LastChild.Attr[1].Val != "" {
		d = append(d, doc[0].LastChild.Attr[1].Val)
	}
	if len(d) <= 0 {
		return nil, fmt.Errorf("[%s] fetchUpdateTime got nothing: %s",
			configs.Data.MS.Title, a.U.String())
	}
	t, err := time.Parse(time.RFC3339, string(d[0]))
	if err != nil {
		return nil, err
	}
	return timestamppb.New(t), nil
}

func shanghai(t time.Time) time.Time {
	loc := time.FixedZone("UTC", 8*60*60)
	return t.In(loc)
}

var ErrTimeOverDays error = errors.New("article update time out of range")
var ErrSameArticleExist error = errors.New("article title exist")

// filter work for ignore articles by conditions
// TODO: filter redundancy articles by title
func (a *Article) filter(days int) (*Article, error) {
	// if article time out of days, return nil and `ErrTimeOverDays`
	// param days means fetch news during days from befor now.
	during := func(days int, ts *timestamppb.Timestamp) bool {
		t := shanghai(ts.AsTime())
		if time.Now().Day()-t.Day() <= days {
			return true
		}
		return false
	}
	// if during return false rt nil, and error as ErrTimeOverDays
	if !during(days, a.UpdateTime) {
		return nil, ErrTimeOverDays
	}

	return a, nil
}

func (a *Article) fetchContent() (string, error) {
	if a.raw == nil {
		return "", errors.Errorf("[%s] fetchContent: raw is nil: %s", configs.Data.MS.Title, a.U.String())
	}

	raw := a.raw
	//td-post-content tagdiv-type
	r := exhtml.DivWithAttr2(raw, "class", "td-post-content tagdiv-type")
	ps := [][]byte{}
	b := bytes.Buffer{}
	re := regexp.MustCompile(`<p.*?>(.*?)</p>`)
	for _, v := range re.FindAllSubmatch(r, -1) {
		ps = append(ps, v[1])
	}
	if len(ps) == 0 {
		return "", fmt.Errorf("no <p> matched")
	}
	for _, p := range ps {
		b.Write(p)
		b.Write([]byte("  \n"))
	}
	body := b.String()
	re = regexp.MustCompile(`「`)
	body = re.ReplaceAllString(body, "“")
	re = regexp.MustCompile(`」`)
	body = re.ReplaceAllString(body, "”")
	// re = regexp.MustCompile(`<a.*?>`)
	// body = re.ReplaceAllString(body, "")
	// re = regexp.MustCompile(`</a>`)
	// body = re.ReplaceAllString(body, "")
	// re = regexp.MustCompile(`<i.*?>`)
	// body = re.ReplaceAllString(body, "")
	// re = regexp.MustCompile(`<!.*?>`)
	// body = re.ReplaceAllString(body, "")
	// re = regexp.MustCompile(`</.*?>`)
	// body = re.ReplaceAllString(body, "")
	re = regexp.MustCompile(`<.*?>`)
	body = re.ReplaceAllString(body, "")
	re = regexp.MustCompile(`Log in to leave a comment   `)
	body = re.ReplaceAllString(body, "")
	re = regexp.MustCompile(`<script.*?</script>`)
	body = re.ReplaceAllString(body, "")
	re = regexp.MustCompile(`<blockquote.*?</blockquote>`)
	body = re.ReplaceAllString(body, "")
	re = regexp.MustCompile(`<iframe.*?</iframe>`)
	body = re.ReplaceAllString(body, "")
	re = regexp.MustCompile(`<strong.*?</strong>`)
	body = re.ReplaceAllString(body, "")

	return body, nil
}

func (a *Article) fmtContent(body string) (string, error) {
	var err error
	title := "# " + a.Title + "\n\n"
	lastupdate := a.UpdateTime.AsTime().Format(time.RFC3339)
	webTitle := fmt.Sprintf(" @ [%s](/list/?v=%[1]s): [%[2]s](http://%[2]s)", a.WebsiteTitle, a.WebsiteDomain)
	u, err := url.QueryUnescape(a.U.String())
	if err != nil {
		u = a.U.String() + "\n\nunescape url error:\n" + err.Error()
	}

	body = title +
		"LastUpdate: " + lastupdate +
		webTitle + "\n\n" +
		"---\n" +
		body + "\n\n" +
		"原地址：" + fmt.Sprintf("[%s](%[1]s)", u)
	return body, nil
}
