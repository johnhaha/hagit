package hagit

import (
	"strings"
	"time"

	"github.com/johnhaha/hakit/hacmd"
)

func Push(commit string, slow bool) {
	hacmd.Execute("git", "add", ".")
	hacmd.Execute("git", "commit", "-m", commit)
	if slow {
		time.Sleep(time.Second)
	}
	hacmd.Execute("git", "push", "-u", "origin", "main")
}

func PushTag(tag string) {
	hacmd.Execute("git", "tag", tag)
	hacmd.Execute("git", "push", "origin", tag)
}

func DeleteTag(tag string) {
	hacmd.Execute("git", "tag", "-d", tag)
	hacmd.Execute("git", "push", "--delete", "origin", tag)
}

func SetRemoteUrl(url string) {
	hacmd.Execute("git", "remote", "set-url", "origin", url)
}

func NewBranch(name string) {
	hacmd.Execute("git", "branch", name)
	hacmd.Execute("git", "checkout", name)
}

func Commit(content string) {
	hacmd.Execute("git", "add", ".")
	hacmd.Execute("git", "commit", "-m", content)
}

func CreateNewRepo(name string) {
	hacmd.Execute("gh", "repo", "create", name, "--private")
}

func SetRepoSecretFromFile(path string) {
	hacmd.Execute("gh", "secret", "set", "-f", path)
}

func CheckStatusContains(msg ...string) bool {
	res, err := hacmd.Run("git", "status")
	if err != nil {
		panic(err)
	}
	for _, m := range msg {
		if strings.Contains(res, m) {
			return true
		}
	}
	return false
}
