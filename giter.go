package hagit

import "os"

type Giter struct {
	Path string
}

func NewGiter() *Giter {
	return &Giter{}
}

func (g *Giter) SetPath(p string) *Giter {
	g.Path = p
	return g
}

func (g *Giter) Push(cmt string) *Giter {
	if g.Path != "" {
		os.Chdir(g.Path)
	}
	Push(cmt)
	return g
}

func (g *Giter) PushTag(tag string) *Giter {
	if g.Path != "" {
		os.Chdir(g.Path)
	}
	PushTag(tag)
	return g
}

func (g *Giter) Delete(tag string) *Giter {
	if g.Path != "" {
		os.Chdir(g.Path)
	}
	DeleteTag(tag)
	return g
}

func (g *Giter) SetRemoteUrl(url string) *Giter {
	if g.Path != "" {
		os.Chdir(g.Path)
	}
	SetRemoteUrl(url)
	return g
}
