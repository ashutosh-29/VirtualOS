package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)
var count int=1
func showTextEditor() {
	w:=myApp.NewWindow("Meta Text Editor")
	content:=container.NewVBox(
		
	)
	input:=widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")

	saveBtn:=widget.NewButton("Save",func() {
		saveFileDialog:=dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData:=[]byte(input.Text)
				uc.Write(textData)
			},w)
			saveFileDialog.SetFileName("New File "+strconv.Itoa(count)+".txt")
			saveFileDialog.Show()
	})
	openBtn:=widget.NewButton("Open",func() {
		openFileDialog:=dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				readData, _ := ioutil.ReadAll(r)
				filename:=r.URI().Name()
				output:=fyne.NewStaticResource(filename,readData)
				viewData:=widget.NewMultiLineEntry()
				viewData.SetText(string(output.StaticContent))
				nw:=fyne.CurrentApp().NewWindow(
					string(output.StaticName),
				)
				savemeBtn:=widget.NewButton("Save",func() {
					saveFileDialog:=dialog.NewFileSave(
						func(uc fyne.URIWriteCloser, _ error) {
							textData:=[]byte(viewData.Text)
							uc.Write(textData)
						},w)
						saveFileDialog.SetFileName(filename)
						saveFileDialog.Show()
				})
				viewData.Resize(fyne.NewSize(800,800))
				nw.SetContent(container.NewVBox(
					savemeBtn,
					container.NewWithoutLayout(viewData),
					
					))
				nw.Resize(fyne.NewSize(400,400))
				nw.Show()
			},w)
		openFileDialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}),
		)
		openFileDialog.Show()
	})
	content.Add(
		container.NewHBox(
			widget.NewButton("Add New File",func() {
				content.Add(
					widget.NewLabel("New File "+strconv.Itoa(count)),
				)
				count++
			}),
			saveBtn,
			openBtn,
		),
	)
	//input.Move(fyne.NewPos(0, 0))
	input.Resize(fyne.NewSize(500, 500))
	w.Resize(fyne.NewSize(500,500))
	myEditorContent:=container.NewVBox(
		content,
		container.NewWithoutLayout(input,),
	)
	w.SetContent(
		container.NewBorder(
		nil,nil,nil,nil,myEditorContent),
	)
	w.Show()
	
}