package gi

type Context struct {
	Title   string
	Html    string
	Head    []string
	Cookies []string
	Url     string // 完整url
}
