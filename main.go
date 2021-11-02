package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App=app.New()

var myWindow fyne.Window = myApp.NewWindow("Fake OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var img fyne.CanvasObject
var DeskBtn fyne.Widget
var panelContent *fyne.Container
func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
func main(){
	myApp.Settings().SetTheme(theme.DarkTheme())
	myTheme:=false
	img=canvas.NewImageFromFile("Images//background.png")

	i1, _ := os.Open("Images//weatherIcon.png")
	r1 := bufio.NewReader(i1)
	weatherIcon, _ := ioutil.ReadAll(r1)
	btn1 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", weatherIcon), func() {
		showWeatherApp(myWindow)
	})
	i1, _ = os.Open("Images//calc.png")
	r1 = bufio.NewReader(i1)
	calcIcon, _ := ioutil.ReadAll(r1)

	btn2 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", calcIcon), func() {
		showCalculator()
	})
	i1, _ = os.Open("Images//galleryicon.png")
	r1 = bufio.NewReader(i1)
	galleryIcon, _ := ioutil.ReadAll(r1)
	
	btn3 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", galleryIcon), func() {
		showGalleryApp(myWindow)
	})
	

	i1, _ = os.Open("Images//editorIcon.png")
	r1 = bufio.NewReader(i1)
	editorIcon, _ := ioutil.ReadAll(r1)
	btn4 = widget.NewButtonWithIcon("", fyne.NewStaticResource("icon", editorIcon), func() {
		showTextEditor()
	})
	
	i1, _ = os.Open("Images//themeIcon.png")
	r1 = bufio.NewReader(i1)
	themeIcon, _ := ioutil.ReadAll(r1)
	btn5:=widget.NewButtonWithIcon("",fyne.NewStaticResource("icon",themeIcon),func() {
		if !myTheme{
			myApp.Settings().SetTheme(theme.LightTheme())
		}else {
			myApp.Settings().SetTheme(theme.DarkTheme())
		}
		myTheme=!myTheme
	})
	i1, _ = os.Open("Images//chrome.png")
	r1 = bufio.NewReader(i1)
	chromeIcon, _ := ioutil.ReadAll(r1)
	btn6:=widget.NewButtonWithIcon("",fyne.NewStaticResource("icon",chromeIcon),func() {
		openbrowser("https://google.com")
	})
	
	i1, _ = os.Open("Images//deskicon.png")
	r1 = bufio.NewReader(i1)
	deskIcon, _ := ioutil.ReadAll(r1)
	DeskBtn =widget.NewButtonWithIcon("",fyne.NewStaticResource("icon",deskIcon),func() {
		myWindow.SetContent(container.NewBorder(nil,nil,panelContent,nil,img))
	})
	
	panelContent=container.NewVBox((container.NewGridWithRows(7,DeskBtn,btn1,btn2,btn3,btn4,btn5,btn6)))

	myWindow.Resize(fyne.NewSize(1280,720))
	myWindow.CenterOnScreen()
	myWindow.SetContent(
		container.NewBorder(nil,nil,panelContent,nil,img),
	)
	myWindow.ShowAndRun()
}