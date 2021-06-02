package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/fatih/color"

)

func main() {
	var Choice int
	fmt.Println(" 1 For Winrar, 2 For Chrome")
	fmt.Print(" Chose Num : ")
	fmt.Scanln(&Choice)
	if Choice == 1 {
		fileUrl := "https://www.win-rar.com/fileadmin/winrar-versions/winrar/winrar-x64-601.exe"
		err := DownloadFile("winrar-x64-601.exe", fileUrl)
		if err != nil {
			panic(err)
		}
		color.Green("\n Winrar Downloaded Successfully")
		color.Red("\n Press Enter To Exit... ")
		fmt.Scanln()
	}

	if Choice == 2{
		fileUrl := "https://dl.google.com/tag/s/appguid%3D%7B8A69D345-D564-463C-AFF1-A69D9E530F96%7D%26iid%3D%7B2574B9B5-5CA9-F135-D92E-0EE4741FF279%7D%26lang%3Den%26browser%3D4%26usagestats%3D1%26appname%3DGoogle%2520Chrome%26needsadmin%3Dprefers%26ap%3Dx64-stable-statsdef_1%26brand%3DCHBD%26installdataindex%3Dempty/update2/installers/ChromeSetup.exe"
		err := DownloadFile("ChromeSetup.exe", fileUrl)
		if err != nil {
			panic(err)
		}
		color.Green("\n Chrome Downloaded Successfully")
		color.Red("\n Press Enter To Exit... ")
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
