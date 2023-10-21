package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	l "GUI-APP/pkg/logic"
)

func main() {
	var upperLetters, lowerLetters, numbers, symbols bool
	num := 20
	myApp := app.New()
	window := myApp.NewWindow("Text Generator")
	imageLogo := canvas.NewImageFromFile("src/logo.png")
	//label := widget.NewLabel("Generate Strong Password -- IMAGE HERE")
	label2 := widget.NewLabel("Settings:")
	labeLength := widget.NewLabel("Length: " + strconv.Itoa(num))
	checkUpperCase := widget.NewCheck("Include Alpha Upper (A-Z)", func(b bool) {
		upperLetters = b
	})
	checkLowerCase := widget.NewCheck("Include Alpha Lower (a-z)", func(b bool) {
		lowerLetters = b
	})
	checkNums := widget.NewCheck("Include Number (0-9)", func(b bool) {
		numbers = b
	})
	checkSymbol := widget.NewCheck("Include Symbol", func(b bool) {
		symbols = b
	})
	input := widget.NewEntry()
	input.SetPlaceHolder("password will apear here...")
	sliderLength := widget.NewSlider(10, 40)
	sliderLength.Value = float64(num)
	imageLogo.FillMode = canvas.ImageFillOriginal
	labelError := canvas.NewText("Choose any check box!", color.NRGBA{R: 255, G: 33, B: 33, A: 255})
	labelError.Hidden = true
	labelMessage := canvas.NewText("Copied successfully!", color.NRGBA{R: 47, G: 221, B: 48, A: 255})
	labelMessage.Hidden = true
	btnCopy := widget.NewButton("Copy Password", func() {
		window.Clipboard().SetContent(input.Text)
		labelMessage.Hidden = false
	})
	btnCopy.Hidden = true
	btn1 := widget.NewButton("Generete", func() {
		input.SetText(l.GeneratePass(int(sliderLength.Value), upperLetters, lowerLetters, numbers, symbols, labelError, labelMessage, btnCopy))
	})
	window.SetContent(
		container.New(
			layout.NewVBoxLayout(),
			container.New(layout.NewHBoxLayout(), layout.NewSpacer(), imageLogo, layout.NewSpacer()),
			container.New(layout.NewHBoxLayout(), label2),
			container.New(layout.NewHBoxLayout(), checkUpperCase, checkLowerCase, checkNums, checkSymbol),
			container.New(layout.NewHBoxLayout(), labeLength),
			container.New(layout.NewVBoxLayout(), sliderLength),
			container.New(layout.NewVBoxLayout(), input),
			container.New(layout.NewHBoxLayout(), btn1, btnCopy),
			container.New(layout.NewHBoxLayout(), labelError, labelMessage),
		),
	)

	window.Resize(fyne.NewSize(700, 500))

	sliderLength.OnChanged = func(value float64) {
		labeLength.SetText("Length: " + strconv.FormatFloat(value, 'f', -1, 64))
	}

	window.ShowAndRun()
}
