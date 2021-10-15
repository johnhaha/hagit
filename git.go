package hagit

import "github.com/johnhaha/hakit/hacmd"

func Push(commit string) {
	hacmd.Execute("git", "add", ".")
	hacmd.Execute("git", "commit", "-m", commit)
	hacmd.Execute("git", "push")
}

func PushTag(tag string) {
	hacmd.Execute("git", "tag", tag)
	hacmd.Execute("git", "push", "origin", tag)
}

func SetRemoteUrl(url string) {
	hacmd.Execute("git", "remote", "set-url", "origin", url)
}
