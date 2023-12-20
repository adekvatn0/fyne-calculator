package main

import (
	"strconv"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//calculator args
var firstArg float64
var action string
var secondArg float64

//output labels
var firstArgLabelOutput = widget.NewLabel("")
var mainOutput = widget.NewLabel("0")

func main() {
	window := app.New().NewWindow("Calculator")

	window.SetContent(
		container.NewVBox(
			firstArgLabelOutput,
			mainOutput,
			container.NewGridWithColumns(
				4,
				//first row
				CreateClearButton(),
				CreateSignButton(),
				CreatePercentButton(),
				CreateActionButton("/"),

				//second row
				CreateDigitalButton(7),
				CreateDigitalButton(8),
				CreateDigitalButton(9),
				CreateActionButton("x"),

				//third row
				CreateDigitalButton(4),
				CreateDigitalButton(5),
				CreateDigitalButton(6),
				CreateActionButton("-"),


				//fourth row
				CreateDigitalButton(1),
				CreateDigitalButton(2),
				CreateDigitalButton(3),
				CreateActionButton("+"),
			),
			//fifth row
			container.NewGridWithColumns(
				2,
				CreateDigitalButton(0),
				container.NewGridWithColumns(
					2,
					CreateFractionalButton(),
					CreateCalculateButton(),
				),
			),
		),
	)

	window.ShowAndRun()
}

func CreateDigitalButton(value int) *widget.Button {
	return widget.NewButton(strconv.Itoa(value), func() {
			if action == "" {
				firstArg = firstArg * 10 + float64(value)
				mainOutput.SetText(FloatToStr(firstArg))
			} else {
				secondArg = secondArg * 10 + float64(value)
				mainOutput.SetText(FloatToStr(secondArg))
			}
	})
}

func CreateActionButton(value string) *widget.Button {
	b := widget.NewButton(value, func() {
		action = value
		firstArgLabelOutput.SetText(FloatToStr(firstArg) + " " + action)
		mainOutput.SetText("0")
	})
	b.Importance = widget.SuccessImportance
	return b
}

func CreateCalculateButton() *widget.Button {
	b := widget.NewButton("=", func() {
		var result float64
		if firstArg != 0 && secondArg != 0 {
			switch action {
			case "+":
				result = firstArg + secondArg
			case "-":
				result = firstArg - secondArg
			case "/":
				result = firstArg / secondArg
			case "x":
				result = firstArg * secondArg

			}
		}
		mainOutput.SetText(FloatToStr(result))
		firstArg = result
		secondArg = 0
		action = ""
		firstArgLabelOutput.SetText("")
	})
	b.Importance = widget.SuccessImportance
	return b
}

func CreateClearButton() *widget.Button {
	b := widget.NewButton("AC", func () {
		firstArg = 0
		secondArg = 0
		action = ""
		mainOutput.SetText("0")
		firstArgLabelOutput.SetText("")
	})
	b.Importance = widget.HighImportance
	return b
}

func CreateSignButton() *widget.Button {
	b := widget.NewButton("+/-", func () {
		if action == "" {
			firstArg = firstArg * -1
			mainOutput.SetText(FloatToStr(firstArg))
		} else {
			secondArg = secondArg * -1
			mainOutput.SetText(FloatToStr(secondArg))
		}
	})
	b.Importance = widget.HighImportance
	return b
}

func CreatePercentButton() *widget.Button {
	b := widget.NewButton("%", func () {
		if action == "" {
			firstArg = firstArg / float64(100)
			mainOutput.SetText(FloatToStr(firstArg))
		} else {
			secondArg = secondArg / float64(100)
			mainOutput.SetText(FloatToStr(secondArg))
		}
	})
	b.Importance = widget.HighImportance
	return b
}

func CreateFractionalButton() *widget.Button {
	b := widget.NewButton(".", func() {

	})
	return b
}

func FloatToStr(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
