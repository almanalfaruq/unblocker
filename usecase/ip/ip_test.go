package ip

import (
	"fmt"
	"testing"

	"github.com/almanalfaruq/unblocker/model"
	"github.com/golang/mock/gomock"
)

func TestIP_WriteToHosts(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	repoIP := NewMockRepoIPIface(ctl)
	repoFile := NewMockRepoFileIface(ctl)
	repoCommand := NewMockRepoCommandIface(ctl)

	ip := New(repoIP, repoFile, repoCommand)
	type args struct {
		url    string
		system string
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr bool
	}{
		{
			name: "Normal-LinuxUnix",
			mock: func() {
				repoIP.EXPECT().GetListIP("test.com").Return(model.Response{
					Answer: []model.Answer{
						{
							Data: "192.168.8.1",
						},
						{
							Data: "192.168.8.2",
						},
					},
				}, nil)
				repoFile.EXPECT().WriteToFile(model.NonWindowsPath, "\n# Whitelist for www.test.com\n").Return(nil)
				repoFile.EXPECT().WriteToFile(model.NonWindowsPath, "192.168.8.1 test.com www.test.com\n").Return(nil)
				repoFile.EXPECT().WriteToFile(model.NonWindowsPath, "192.168.8.2 test.com www.test.com\n").Return(nil)
			},
			args: args{
				url:    "https://www.test.com",
				system: "linux",
			},
		},
		{
			name: "Normal-Windows",
			mock: func() {
				repoIP.EXPECT().GetListIP("test.com").Return(model.Response{
					Answer: []model.Answer{
						{
							Data: "192.168.8.1",
						},
						{
							Data: "192.168.8.2",
						},
					},
				}, nil)
				repoFile.EXPECT().WriteToFile(model.WindowsPath, "\n# Whitelist for www.test.com\n").Return(nil)
				repoFile.EXPECT().WriteToFile(model.WindowsPath, "192.168.8.1 test.com www.test.com\n").Return(nil)
				repoFile.EXPECT().WriteToFile(model.WindowsPath, "192.168.8.2 test.com www.test.com\n").Return(nil)
				repoCommand.EXPECT().RunFlushDNS().Return(nil)
			},
			args: args{
				url:    "https://www.test.com",
				system: model.Windows,
			},
		},
		{
			name: "Error-GetListIP",
			mock: func() {
				repoIP.EXPECT().GetListIP("test.com").Return(model.Response{}, fmt.Errorf("error"))
			},
			args: args{
				url:    "https://www.test.com",
				system: "linux",
			},
			wantErr: true,
		},
		{
			name: "Error-WriteComment",
			mock: func() {
				repoIP.EXPECT().GetListIP("test.com").Return(model.Response{
					Answer: []model.Answer{
						{
							Data: "192.168.8.1",
						},
						{
							Data: "192.168.8.2",
						},
					},
				}, nil)
				repoFile.EXPECT().WriteToFile(model.NonWindowsPath, "\n# Whitelist for www.test.com\n").Return(fmt.Errorf("error"))
			},
			args: args{
				url:    "https://www.test.com",
				system: "linux",
			},
			wantErr: true,
		},
		{
			name: "Error-WriteIP",
			mock: func() {
				repoIP.EXPECT().GetListIP("test.com").Return(model.Response{
					Answer: []model.Answer{
						{
							Data: "192.168.8.1",
						},
					},
				}, nil)
				repoFile.EXPECT().WriteToFile(model.NonWindowsPath, "\n# Whitelist for www.test.com\n").Return(nil)
				repoFile.EXPECT().WriteToFile(model.NonWindowsPath, "192.168.8.1 test.com www.test.com\n").Return(fmt.Errorf("error"))
			},
			args: args{
				url:    "https://www.test.com",
				system: "linux",
			},
			wantErr: true,
		},
		{
			name: "Error-FlushDNS",
			mock: func() {
				repoIP.EXPECT().GetListIP("test.com").Return(model.Response{
					Answer: []model.Answer{
						{
							Data: "192.168.8.1",
						},
					},
				}, nil)
				repoFile.EXPECT().WriteToFile(model.WindowsPath, "\n# Whitelist for www.test.com\n").Return(nil)
				repoFile.EXPECT().WriteToFile(model.WindowsPath, "192.168.8.1 test.com www.test.com\n").Return(nil)
				repoCommand.EXPECT().RunFlushDNS().Return(fmt.Errorf("error"))
			},
			args: args{
				url:    "https://www.test.com",
				system: model.Windows,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if err := ip.WriteToHosts(tt.args.url, tt.args.system); (err != nil) != tt.wantErr {
				t.Errorf("IP.WriteToHosts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
