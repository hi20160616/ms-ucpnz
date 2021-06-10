package fetcher

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/hi20160616/exhtml"
	"github.com/hi20160616/ms-ucpnz/configs"
	"github.com/pkg/errors"
)

// pass test
func TestFetchArticle(t *testing.T) {
	tests := []struct {
		url string
		err error
	}{
		{"https://ucpnz.co.nz/2021/02/05/%e9%9f%a9%e5%9b%bd%e5%80%99%e9%80%89%e4%ba%ba%e9%80%80%e5%87%ba%e4%b8%96%e8%b4%b8%e6%80%bb%e5%b9%b2%e4%ba%8b%e9%80%89%e4%b8%be/", ErrTimeOverDays},
		{"https://ucpnz.co.nz/2021/06/10/%E9%99%86%E5%86%9B%E5%A4%9A%E8%AF%BE%E7%9B%AE%E4%BC%9E%E9%99%8D%E8%AE%AD%E7%BB%83-%E9%94%A4%E7%82%BC%E7%AB%8B%E4%BD%93%E6%B8%97%E9%80%8F%E8%83%BD%E5%8A%9B/", nil},
	}
	for _, tc := range tests {
		a := NewArticle()
		a, err := a.fetchArticle(tc.url)
		if err != nil {
			if !errors.Is(err, ErrTimeOverDays) {
				t.Error(err)
			} else {
				fmt.Println("ignore old news pass test: ", tc.url)
			}
		} else {
			fmt.Println("pass test: ", a.Content)
		}
	}
}

func TestFetchTitle(t *testing.T) {
	tests := []struct {
		url   string
		title string
	}{
		{"https://www.ucpnz.com/realtime/world/story20210602-1151196", "马国男子腰缠巨蟒骑摩托车送往放生引热议"},
		{"https://www.ucpnz.com/realtime/world/story20210607-1153241", "以色列将于14日前投票批准新政府"},
	}
	for _, tc := range tests {
		a := NewArticle()
		u, err := url.Parse(tc.url)
		if err != nil {
			t.Error(err)
		}
		a.U = u
		// Dail
		a.raw, a.doc, err = exhtml.GetRawAndDoc(a.U, timeout)
		if err != nil {
			t.Error(err)
		}
		got, err := a.fetchTitle()
		if err != nil {
			if !errors.Is(err, ErrTimeOverDays) {
				t.Error(err)
			} else {
				fmt.Println("ignore pass test: ", tc.url)
			}
		} else {
			if tc.title != got {
				t.Errorf("\nwant: %s\n got: %s", tc.title, got)
			}
		}
	}

}

func TestFetchUpdateTime(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{
			"https://www.ucpnz.com/realtime/world/story20210602-1151196",
			"2021-06-02 15:44:33 +0800 UTC",
		},
		{
			"https://www.ucpnz.com/realtime/world/story20210607-1153241",
			"2021-06-07 21:38:53 +0800 UTC",
		},
	}
	var err error
	if err := configs.Reset("../../"); err != nil {
		t.Error(err)
	}

	for _, tc := range tests {
		a := NewArticle()
		a.U, err = url.Parse(tc.url)
		if err != nil {
			t.Error(err)
		}
		// Dail
		a.raw, a.doc, err = exhtml.GetRawAndDoc(a.U, timeout)
		if err != nil {
			t.Error(err)
		}
		tt, err := a.fetchUpdateTime()
		if err != nil {
			t.Error(err)
		} else {
			ttt := tt.AsTime()
			got := shanghai(ttt)
			if got.String() != tc.want {
				t.Errorf("\nwant: %s\n got: %s", tc.want, got.String())
			}
		}
	}
}

func TestFetchContent(t *testing.T) {
	tests := []struct {
		url  string
		want string
	}{
		{
			"https://www.ucpnz.com/realtime/world/story20210602-1151196",
			"2021-06-02 15:44:33 +0800 UTC",
		},
		{
			"https://www.ucpnz.com/realtime/world/story20210607-1153241",
			"2021-06-07 21:38:53 +0800 UTC",
		},
	}
	var err error
	if err := configs.Reset("../../"); err != nil {
		t.Error(err)
	}

	for _, tc := range tests {
		a := NewArticle()
		a.U, err = url.Parse(tc.url)
		if err != nil {
			t.Error(err)
		}
		// Dail
		a.raw, a.doc, err = exhtml.GetRawAndDoc(a.U, timeout)
		if err != nil {
			t.Error(err)
		}
		c, err := a.fetchContent()
		if err != nil {
			t.Error(err)
		}
		fmt.Println(c)
	}
}
