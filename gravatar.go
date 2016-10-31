package gravago

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
)

const (
	baseURL       string = "//www.gravatar.com/avatar/"
	defaultIcons  string = "retro"
	defaultRating string = "g"
	defaultSize   uint   = 64
)

// gravatar object
type gravatar struct {
	GravatarURL string
	Icons       string
	Rating      string
	Size        uint
}

// New - create new gravatar object instance
func New() gravatar {
	g := &gravatar{
		GravatarURL: baseURL,
		Icons:       defaultIcons,
		Rating:      defaultRating,
		Size:        defaultSize,
	}

	return g
}

// SetSize - set size of gravatar image
func (g *gravatar) SetSize(uint size) {
	g.Size = size
}

// SetRating - set rating
func (g *gravatar) SetRating(string rating) {
	g.Rating = rating
}

// SetIcons - set the default icon to return
func (g *gravatar) SetIcons(string iconset) {
	g.Icons = iconset
}

// UseHTTPS - force using HTTPS protocol
func (g *gravatar) UseHTTPS(bool protocol) {
	if protocol {
		g.GravatarURL = "https://s.gravatar.com/avatar/"
	} else {
		g.GravatarURL = "http:" + baseURL
	}
}

// URL - return gravatar URL
func (g *gravatar) URL(string email) string {
	var GravURL *url.URL
	GravURL, err := url.Parse(g.GravatarURL)

	check(err)

	GravURL.Path += hash + ".jpg"

	parameters := url.Values{}
	parameters.Add("s", fmt.Sprintf("%v", g.Size))
	parameters.Add("r", fmt.Sprintf("%v", g.Rating))
	parameters.Add("d", fmt.Sprintf("%v", g.Icons))

	GravURL.RawQuery = parameters.Encode()

	return GravURL
}

func emailToHash(string email) string {
	hasher := md5.New()
	hasher.Write([]byte(email))
	return fmt.Sprintf("%v", hex.EncodeToString(hasher.Sum(nil)))
}
