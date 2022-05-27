package main

import (
	"fmt"
	"os/exec"

	"github.com/gen2brain/dlgs"
	"github.com/getlantern/systray"
	"github.com/imroc/req/v3"
	"github.com/martinlindhe/notify"
	"github.com/wlh320/portguard-systray2/icon"
)

func main() {
	// parameters
	iconPath := "/home/wlh/portguard/icon/icon.png"
	onReady := func() {
		go setMenu(iconPath)
	}
	onExit := func() {
	}
	systray.Run(onReady, onExit)
}

func setMenu(iconPath string) {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Portgaurd Systray")
	systray.SetTooltip("Portguard Systray App")

	// 1. clash_mode
	mClash := systray.AddMenuItem("Clash mode", "Change clash mode")
	// init
	mClash.SetTitle("clash mode: " + ToggleClashMode())
	mClash.SetTitle("clash mode: " + ToggleClashMode())

	// 2. pg_mode
	mPg := systray.AddMenuItem("Portguard mode", "Change portguard mode")
	// init
	mPg.SetTitle("pg mode: " + TogglePGMode())
	mPg.SetTitle("pg mode: " + TogglePGMode())

	// 3. About & Quit button
	systray.AddSeparator()
	mAbout := systray.AddMenuItem("About", "About this app")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			select {
			case <-mClash.ClickedCh:
				output := ToggleClashMode()
				mClash.SetTitle("clash mode: " + output)
				notify.Notify("clash", "clash mode", "Change clash mode to "+output, iconPath)
			case <-mPg.ClickedCh:
				output := TogglePGMode()
				mPg.SetTitle("pg mode: " + output)
				notify.Notify("portguard", "pg mode", "Change portguard mode to "+output, iconPath)
			case <-mAbout.ClickedCh:
				dlgs.MessageBox("About portguard-systray2", "Version 0.0.1")
			case <-mQuit.ClickedCh:
				fmt.Println("Requesting quit")
				systray.Quit()
				fmt.Println("Finished quitting")
				return
			}
		}
	}()
}

func ToggleClashMode() string {
	url := "http://127.0.0.1:9090/configs"
	res, err := req.Get(url)
	if err != nil {
		return "none"
	}
	ans := make(map[string]string)
	res.UnmarshalJson(&ans)
	currMode := ans["mode"]
	if currMode == "rule" { // rule -> direct
		if _, err := req.SetBodyJsonString(`{"mode": "direct"}`).Patch(url); err != nil {
			dlgs.Warning("Fail", err.Error())
			return "rule"
		}
		return "direct"
	} else { // direct -> rule
		if _, err := req.SetBodyJsonString(`{"mode": "rule"}`).Patch(url); err != nil {
			dlgs.Warning("Fail", err.Error())
			return "direct"
		}
		return "rule"
	}
}

func TogglePGMode() string {
	stdout, _ := exec.Command("pgrep", "-x", "dorm2free").Output()
	isRunningV2ray := (len(stdout) != 0)
	if isRunningV2ray { // v2ray -> socks5
		if err := exec.Command("systemctl", "--user", "stop", "dorm2free").Run(); err != nil {
			dlgs.Warning("Fail", err.Error())
			return "v2ray"
		}
		if err := exec.Command("systemctl", "--user", "start", "dorm2free_direct").Run(); err != nil {
			dlgs.Warning("Fail", err.Error())
			return "v2ray"
		}
		return "direct"
	} else { // socks5 -> v2ray
		if err := exec.Command("systemctl", "--user", "stop", "dorm2free_direct").Run(); err != nil {
			dlgs.Warning("Fail", err.Error())
			return "direct"
		}
		if err := exec.Command("systemctl", "--user", "start", "dorm2free").Run(); err != nil {
			dlgs.Warning("Fail", err.Error())
			return "direct"
		}
		return "v2ray"
	}
}
