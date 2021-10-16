package hagit

import "os"

type Giter struct {
	Path       string
	OriginPath string
}

func NewGiter() *Giter {
	g := new(Giter)
	of, _ := os.Getwd()
	g.OriginPath = of
	return g
}

func (g *Giter) SetPath(p string) *Giter {
	g.Path = p
	return g
}

//back to origin dir
func (g *Giter) Back() *Giter {
	if g.OriginPath != "" {
		os.Chdir(g.OriginPath)
	}
	return g
}

//push file
func (g *Giter) Push(cmt string) *Giter {
	if g.Path != "" {
		os.Chdir(g.Path)
	}
	Push(cmt)
	g.Back()
	return g
}

func (g *Giter) PushTag(tag string) *Giter {
	if g.Path != "" {
		os.Chdir(g.Path)
	}
	PushTag(tag)
	g.Back()
	return g
}

func (g *Giter) Delete(tag string) *Giter {
	if g.Path != "" {
		os.Chdir(g.Path)
	}
	DeleteTag(tag)
	g.Back()
	return g
}

func (g *Giter) SetRemoteUrl(url string) *Giter {
	if g.Path != "" {
		os.Chdir(g.Path)
	}
	SetRemoteUrl(url)
	g.Back()
	return g
}
