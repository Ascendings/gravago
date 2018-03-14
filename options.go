package gravago

// GravatarOptions - gravatar struct settings
type GravatarOptions struct {
  ForceHTTPS bool
  IconSet    string
  IconRating string
  IconSize   int
}

const (
  defaultBaseURL    string = "gravatar.com/avatar/"
  defaultIconSet    string = "retro"
  defaultIconRating string = "g"
  defaultIconSize   int    = 64
)
