package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)
func query(loc string)[] string{
	res, err:=http.Get("https://api.openweathermap.org/data/2.5/weather?q="+loc+"&APPID=2aa4e2f1f576b58d0eb58b6d1cc7ec14");

	if err!=nil{
		fmt.Println(err)
	}
	defer res.Body.Close()
	
	body, err := ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println(err)
	}
	weather, err:=UnmarshalWelcome(body)
	if err!=nil{
		fmt.Println(err)
	}
	var list[] string

	list=append(list,weather.Sys.Country)
	list=append(list,fmt.Sprintf("%.2f",weather.Wind.Speed))
	list=append(list,fmt.Sprintf("%.2f",weather.Main.Temp))
	list=append(list,fmt.Sprintf("%d",weather.Main.Humidity))
	return list
}
func showWeatherApp(w fyne.Window){
	
	input := widget.NewEntry()
	input.SetPlaceHolder(" Enter the Location ")
	
	img:= canvas.NewImageFromFile("Images//bg.png")
	img.FillMode = canvas.ImageFillOriginal

	label1:=canvas.NewText("Weather Details",color.White)
	label1.TextStyle = fyne.TextStyle{Bold:true}

	label2:= widget.NewLabel("Country ")
	
	label3:=widget.NewLabel("Wind Speed ")

	label4:=widget.NewLabel("Temperature ")

	label5:=widget.NewLabel("Humidity ")

	content:=container.NewVBox(
			input,
			widget.NewButton("Search",func() {
				list:=query(input.Text)
				label2.SetText("Country Code:  "+list[0])
				label3.SetText("Wind Speed:   "+list[1])
				label4.SetText("Temperature: "+list[2])
				label5.SetText("Humidity:       "+list[3])
			}),
			label1,
			img,
			label2,
			label3,
			label4,
			label5,
		)
	w.SetContent(
		container.NewBorder(nil,nil,panelContent,nil,content),
	)
	w.Show()
}





func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`     
	Weather    []Weather `json:"weather"`   
	Base       string    `json:"base"`      
	Main       Main      `json:"main"`      
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`      
	Clouds     Clouds    `json:"clouds"`    
	Dt         int64     `json:"dt"`        
	Sys        Sys       `json:"sys"`       
	Timezone   int64     `json:"timezone"`  
	ID         int64     `json:"id"`        
	Name       string    `json:"name"`      
	Cod        int64     `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
	SeaLevel  int64   `json:"sea_level"` 
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type Weather struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
	Gust  float64 `json:"gust"` 
}
