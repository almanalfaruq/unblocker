package ip

import (
	"fmt"
	"os/exec"
	"strings"
	"unblocker/model"
)

type IP struct {
	repoIP   repoIPIface
	repoFile repoFileIface
}

type repoIPIface interface {
	GetListIP(url string) (model.Response, error)
}

type repoFileIface interface {
	WriteToFile(path, text string) error
}

func New(repoIP repoIPIface, repoFile repoFileIface) *IP {
	return &IP{
		repoIP:   repoIP,
		repoFile: repoFile,
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
		cmd := exec.Command("ipconfig", "/flushdns")
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("Must run flushdns manually\n Error: %v", err)
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
