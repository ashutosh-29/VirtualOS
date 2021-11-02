package main

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"io/ioutil"
	"log"
	"strings"
	"fyne.io/fyne/v2/canvas"
)

func showGalleryApp(w fyne.Window) {
	
	root_src:="E:\\Wallpaper";

	files, err := ioutil.ReadDir(root_src)
    if err != nil {
        log.Fatal(err)
    }
	var imgArr[] string
    for _, file := range files {
		if !file.IsDir(){
			ext:=strings.Split(file.Name(),".")[1]
			if ext=="png" || ext=="jpg" || ext=="jpeg"{
				imgArr=append(imgArr, root_src+"\\"+file.Name())
			}
		}
    }
	tabs := container.NewAppTabs(
		container.NewTabItem(strings.Split(imgArr[0],"\\")[2], canvas.NewImageFromFile(imgArr[0])),
	)
	for i:=1;i<len(imgArr);i++ {
		tabs.Append(container.NewTabItem(strings.Split(imgArr[i],"\\")[2], canvas.NewImageFromFile(imgArr[i])))
	}
	w.SetContent(
		container.NewBorder(nil,nil,panelContent,nil,tabs),
	)
	w.Show() 
}