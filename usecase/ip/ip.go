package ip

import (
	"fmt"
	"strings"

	"github.com/almanalfaruq/unblocker/model"
)

//go:generate mockgen -destination=ip_mock_test.go -package=ip github.com/almanalfaruq/unblocker/usecase/ip RepoIPIface,RepoFileIface,RepoCommandIface
type IP struct {
	repoIP      RepoIPIface
	repoFile    RepoFileIface
	repoCommand RepoCommandIface
}

type RepoIPIface interface {
	GetListIP(url string) (model.Response, error)
}

type RepoFileIface interface {
	WriteToFile(path, text string) error
}

type RepoCommandIface interface {
	RunFlushDNS() error
}

func New(repoIP RepoIPIface, repoFile RepoFileIface, repoCommand RepoCommandIface) *IP {
	return &IP{
		repoIP:      repoIP,
		repoFile:    repoFile,
		repoCommand: repoCommand,
	}
}

func (ip *IP) WriteToHosts(url, system string) error {
	var listIP model.ListIP

	url = ip.formatURL(url)

	resp, err := ip.repoIP.GetListIP(url)
	if err != nil {
		return err
	}

	for _, ans := range resp.Answer {
		ip := ans.Data
		listIP = append(listIP, ip)
	}

	path := model.NonWindowsPath
	if system == model.Windows {
		path = model.WindowsPath
	}

	err = ip.writeComment(path, url)
	if err != nil {
		return err
	}

	err = ip.writeIPs(path, url, listIP)
	if err != nil {
		return err
	}

	if system == model.Windows {
		err = ip.repoCommand.RunFlushDNS()
		if err != nil {
			return err
		}
	}

	return nil
}

func (ip *IP) formatURL(url string) string {
	if strings.Contains(url, "http") || strings.Contains(url, "https") {
		splitURL := strings.Split(url, "/")
		url = splitURL[2]
	}

	if strings.Contains(url, "www") {
		splitURL := strings.Split(url, "www.")
		url = splitURL[1]
	}

	return url
}

func (ip *IP) writeComment(path, url string) error {
	comment := fmt.Sprintf("\n# Whitelist for www.%s\n", url)
	return ip.repoFile.WriteToFile(path, comment)
}

func (ip *IP) writeIPs(path, url string, listIP model.ListIP) error {
	for _, oneIP := range listIP {
		text := fmt.Sprintf("%s %s www.%s\n", oneIP, url, url)
		err := ip.repoFile.WriteToFile(path, text)
		if err != nil {
			return err
		}
	}

	return nil
}
