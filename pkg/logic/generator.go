package logic

import (
	"math/rand"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func GeneratePass(quantSymbols int, upper bool, lower bool, num bool, symbols bool, labelError *canvas.Text, labelMessage *canvas.Text, btnCopy *widget.Button) string {
	min := 0
	max := 4
	var pass string
	upperLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerLetters := "abcdefghijklmnopqrstuvmxyz"
	numbers := "0123456789"
	symbolsL := "@@@#$%^&*()!!@**!&@&"
	if upper || lower || num || symbols {
		labelError.Hidden = true
		labelMessage.Hidden = true
		btnCopy.Hidden = false
		for i := 0; i < quantSymbols; i++ {
			randNum := rand.Intn(max-min) + min
			switch {
			case randNum == 0:
				if upper {
					pass += string(upperLetters[rand.Intn(len(upperLetters))])
				} else {
					i--
				}
			case randNum == 1:
				if lower {
					pass += string(lowerLetters[rand.Intn(len(lowerLetters))])
				} else {
					i--
				}
			case randNum == 2:
				if num {
					pass += string(numbers[rand.Intn(len(numbers))])
				} else {
					i--
				}
			case randNum == 3:
				if symbols {
					pass += string(symbolsL[rand.Intn(len(symbolsL))])
				} else {
					i--
				}
			}
		}
	} else {
		labelError.Hidden = false
		btnCopy.Hidden = true
	}

	return pass
}
