package main

import (
	"fmt"
	"os"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var contrastValue float64

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(1200, 900)

	layout, _, _ := image()

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	splitter := widgets.NewQSplitter2(core.Qt__Horizontal, nil)
	splitter.AddWidget(widget)

	window.SetCentralWidget(splitter)

	window.Show()

	app.Exec()
}

func image() (*widgets.QHBoxLayout, *widgets.QLabel, *gui.QPixmap) {
	layout := widgets.NewQHBoxLayout()


	pixmap1 := gui.NewQPixmap()
	pixmap1.Load("photo.jpg", "", core.Qt__AutoColor)
	label1 := widgets.NewQLabel(nil, 0)
	label1.SetPixmap(pixmap1)
	scaledPixmap1 := pixmap1.Scaled2(pixmap1.Size().Width()/2, pixmap1.Size().Height()/2, core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation)
	label1.SetPixmap(scaledPixmap1)

	
	pixmap2 := gui.NewQPixmap()
	pixmap2.Load("output.png", "", core.Qt__AutoColor)
	label2 := widgets.NewQLabel(nil, 0)
	label2.SetPixmap(pixmap2)
	scaledPixmap2 := pixmap2.Scaled2(pixmap2.Size().Width()/2, pixmap2.Size().Height()/2, core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation)
	label2.SetPixmap(scaledPixmap2)

	label1.SetFixedWidth(400)
	label2.SetFixedWidth(400)

	layout.AddWidget(Contrast(), 0, core.Qt__AlignTop)
	layout.AddWidget(Gamma(), 0, core.Qt__AlignTop)
	layout.AddWidget(Brightness(), 0, core.Qt__AlignTop)
	layout.AddWidget(Hue(), 0, core.Qt__AlignTop)
	layout.AddWidget(label1, 0, core.Qt__AlignCenter)

	return layout, label2, pixmap2
}

func Contrast() *widgets.QPushButton {
	img, err := imgio.Open("photo.jpg")
	if err != nil {
		fmt.Println(err)
	}

	moreContrastButton := widgets.NewQPushButton2("Изменить контрастность", nil)
	moreContrastButton.ConnectClicked(func(checked bool) {

		moreContrastButton.SetStyleSheet(`QPushButton:pressed { background-color: #4CAF50; }`)

		dialog := widgets.NewQInputDialog(nil, 0)
		dialog.SetInputMode(widgets.QInputDialog__DoubleInput)
		dialog.SetDoubleRange(-100, 100)
		dialog.SetDoubleValue(1.0)
		ok := dialog.Exec()
		if ok == int(widgets.QDialog__Accepted) {
			contrastValue = dialog.DoubleValue()
			rotated := adjust.Contrast(img, contrastValue)

			if err := imgio.Save("output.png", rotated, imgio.PNGEncoder()); err != nil {
				fmt.Println(err)
			}
			_, imageLabel, pixmap := image()
			newPixmap := pixmap.Scaled2(800, 600, core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation)
			imageLabel.SetPixmap(newPixmap)
			imageLabel.Show()
		}
	})

	moreContrastButton.SetFixedWidth(200)
	moreContrastButton.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)

	moreContrastButton.SetStyleSheet(`QPushButton { text-align: center; border: 1px solid black; }`)

	return moreContrastButton
}

func Gamma() *widgets.QPushButton {
	img, err := imgio.Open("photo.jpg")
	if err != nil {
		fmt.Println(err)
	}

	moreGammaButton := widgets.NewQPushButton2("Изменить гамму", nil)
	moreGammaButton.ConnectClicked(func(checked bool) {
		moreGammaButton.SetStyleSheet(`QPushButton:pressed { background-color: #4CAF50; }`)

		dialog := widgets.NewQInputDialog(nil, 0)
		dialog.SetInputMode(widgets.QInputDialog__DoubleInput)
		dialog.SetDoubleRange(-100, 100)
		dialog.SetDoubleValue(1.0)

		ok := dialog.Exec()
		if ok == int(widgets.QDialog__Accepted) {
			contrastValue = dialog.DoubleValue()
			rotated := adjust.Gamma(img, contrastValue)

			if err := imgio.Save("output.png", rotated, imgio.PNGEncoder()); err != nil {
				fmt.Println(err)
			}
			_, imageLabel, pixmap := image()
			newPixmap := pixmap.Scaled2(800, 600, core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation)
			imageLabel.SetPixmap(newPixmap)
			imageLabel.Show()
		}
	})

	moreGammaButton.SetFixedWidth(200)
	moreGammaButton.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)

	moreGammaButton.SetStyleSheet(`QPushButton { text-align: center; border: 1px solid black; }`)

	return moreGammaButton
}


func Brightness() *widgets.QPushButton {
	img, err := imgio.Open("photo.jpg")
	if err != nil {
		fmt.Println(err)
	}

	moreGammaButton := widgets.NewQPushButton2("Изменить яркость", nil)
	moreGammaButton.ConnectClicked(func(checked bool) {
		moreGammaButton.SetStyleSheet(`QPushButton:pressed { background-color: #4CAF50; }`)

		dialog := widgets.NewQInputDialog(nil, 0)
		dialog.SetInputMode(widgets.QInputDialog__DoubleInput)
		dialog.SetDoubleRange(-1, 1)
		dialog.SetDoubleValue(1.0)

		ok := dialog.Exec()
		if ok == int(widgets.QDialog__Accepted) {
			contrastValue = dialog.DoubleValue()
			rotated := adjust.Brightness(img, contrastValue)

			if err := imgio.Save("output.png", rotated, imgio.PNGEncoder()); err != nil {
				fmt.Println(err)
			}
			_, imageLabel, pixmap := image()
			newPixmap := pixmap.Scaled2(800, 600, core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation)
			imageLabel.SetPixmap(newPixmap)
			imageLabel.Show()
		}
	})

	moreGammaButton.SetFixedWidth(200)
	moreGammaButton.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)

	moreGammaButton.SetStyleSheet(`QPushButton { text-align: center; border: 1px solid black; }`)

	return moreGammaButton
}

func Hue() *widgets.QPushButton {
	img, err := imgio.Open("photo.jpg")
	if err != nil {
		fmt.Println(err)
	}

	moreGammaButton := widgets.NewQPushButton2("Изменить оттенок", nil)
	moreGammaButton.ConnectClicked(func(checked bool) {
		moreGammaButton.SetStyleSheet(`QPushButton:pressed { background-color: #4CAF50; }`)

		dialog := widgets.NewQInputDialog(nil, 0)
		dialog.SetInputMode(widgets.QInputDialog__DoubleInput)
		dialog.SetDoubleRange(-360, 360)
		dialog.SetDoubleValue(1.0)

		ok := dialog.Exec()
		if ok == int(widgets.QDialog__Accepted) {
			contrastValue = dialog.DoubleValue()
			rotated := adjust.Hue(img, int(contrastValue))

			if err := imgio.Save("output.png", rotated, imgio.PNGEncoder()); err != nil {
				fmt.Println(err)
			}
			_, imageLabel, pixmap := image()
			newPixmap := pixmap.Scaled2(800, 600, core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation)
			imageLabel.SetPixmap(newPixmap)
			imageLabel.Show()
		}
	})

	moreGammaButton.SetFixedWidth(200)
	moreGammaButton.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)

	moreGammaButton.SetStyleSheet(`QPushButton { text-align: center; border: 1px solid black; }`)

	return moreGammaButton
}
