package gravago

const (
	gravatarURL    string = "https://www.gravatar.com/avatar/"
	defaultIconSet string = "retro"
	defaultRating  string = "g"
	defaultSize    uint   = 64
)

// Gravatar - gravatar object
type gravatar struct {
	IconSet string
	Rating  string
	Size    uint
}

var grav gravatar
