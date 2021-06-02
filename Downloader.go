package main

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	var Choice int
	color.Cyan(" 1 For Winrar, 2 For Chrome, 3 For FireFox, 4 For Telegram, 5 For Discord, 6 For @m1c1 Registry")
	fmt.Print(" Chose Num : ")
	fmt.Scanln(&Choice)
	if Choice == 1 {
		fileUrl := "https://www.win-rar.com/fileadmin/winrar-versions/winrar/winrar-x64-601.exe"
		err := DownloadFile("winrar-x64-601.exe", fileUrl)
		if err != nil {
			panic(err)
		}
		color.Green("\n Winrar Downloaded Successfully")
		time.Sleep(time.Second)
		color.Blue("\n Programmed By: SNjX, https://instagram.com/6o9s ")
		time.Sleep(time.Second)
		color.Red("\n Press Enter To exit... ")
		fmt.Scanln()
	}

	if Choice == 2{
		fileUrl := "https://dl.google.com/tag/s/appguid%3D%7B8A69D345-D564-463C-AFF1-A69D9E530F96%7D%26iid%3D%7B2574B9B5-5CA9-F135-D92E-0EE4741FF279%7D%26lang%3Den%26browser%3D4%26usagestats%3D1%26appname%3DGoogle%2520Chrome%26needsadmin%3Dprefers%26ap%3Dx64-stable-statsdef_1%26brand%3DCHBD%26installdataindex%3Dempty/update2/installers/ChromeSetup.exe"
		err := DownloadFile("ChromeSetup.exe", fileUrl)
		if err != nil {
			panic(err)
		}
		color.Green("\n Chrome Downloaded Successfully ")
		time.Sleep(time.Second)
		color.Blue("\n Programmed By: SNjX, https://instagram.com/6o9s ")
		time.Sleep(time.Second)
		color.Red("\n Press Enter To exit... ")
		fmt.Scanln()
	}
	if Choice == 3{
		fileUrl := "https://download-installer.cdn.mozilla.net/pub/firefox/nightly/latest-mozilla-central/firefox-91.0a1.en-US.win64.installer.msi"
		err := DownloadFile("ChromeSetup.exe", fileUrl)
		if err != nil {
			panic(err)
		}
		color.Green("\n FireFox Downloaded Successfully")
		time.Sleep(time.Second)
		color.Blue("\n Programmed By: SNjX, https://instagram.com/6o9s ")
		time.Sleep(time.Second)
		color.Red("\n Press Enter To exit... ")
		fmt.Scanln()
	}
	if Choice == 4{
		fileUrl := "https://updates.tdesktop.com/tx64/tsetup-x64.2.7.4.exe"
		err := DownloadFile("tsetup-x64.2.7.4.exe", fileUrl)
		if err != nil {
			panic(err)
		}
		color.Green("\n Telegram Downloaded Successfully")
		time.Sleep(time.Second)
		color.Blue("\n Programmed By: SNjX, https://instagram.com/6o9s ")
		time.Sleep(time.Second)
		color.Red("\n Press Enter To exit... ")
		fmt.Scanln()
	}
	if Choice == 5{
		fileUrl := "https://dl.discordapp.net/distro/app/stable/win/x86/1.0.9002/DiscordSetup.exe"
		err := DownloadFile("DiscordSetup.exe", fileUrl)
		if err != nil {
			panic(err)
		}
		color.Green("\n Discord Downloaded Successfully")
		time.Sleep(time.Second)
		color.Blue("\n Programmed By: SNjX, https://instagram.com/6o9s ")
		time.Sleep(time.Second)
		color.Red("\n Press Enter To exit... ")
		fmt.Scanln()

	}
	if Choice == 6{
		fileUrl := "https://codeload.github.com/rayan1198/Falconprogram/zip/refs/heads/main"
		err := DownloadFile("main", fileUrl)
		if err != nil {
			panic(err)
		}
		color.Green("\n Reg Downloaded Successfully")
		time.Sleep(time.Second)
		color.Blue("\n Programmed By: SNjX, https://instagram.com/6o9s ")
		time.Sleep(time.Second)
		color.Red("\n Press Enter To exit... ")
		fmt.Scanln()

	}
		if Choice > 6 {
			color.Red(" Wrong Choice ")
			time.Sleep(time.Second)
			color.Blue("\n Programmed By: SNjX, https://instagram.com/6o9s ")
			time.Sleep(time.Second)
			color.Red("\n Press Enter To exit... ")
			fmt.Scanln()
		}
}
func DownloadFile(filepath string, url string) error{

		resp, err := http.Get(url)
		if err != nil{
		return err
	}
		defer resp.Body.Close()

		out, err := os.Create(filepath)
		if err != nil{
		return err
	}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		return err

}
