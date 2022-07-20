package main

import (
	"log"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func main() {
	ui.Main(setupUI)
	ui.OnShouldQuit(func() bool {
		log.Println("ui.quit")
		return true
	})
}

func setupUI() {
	var win = ui.NewWindow("title", 500, 300, false)
	win.SetMargined(true)
	win.OnClosing(func(w *ui.Window) bool {
		log.Println("win.closing")
		ui.Quit()
		return true
	})

	var vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	win.SetChild(vbox)
	{
		var label = ui.NewLabel("title")
		vbox.Append(label, false)
	}

	{
		var btn = ui.NewButton("ok")
		vbox.Append(btn, false)
	}

	{
		var btn = ui.NewColorButton()
		btn.OnChanged(func(cb *ui.ColorButton) {
			log.Println(cb.Color())
		})
		vbox.Append(btn, false)
	}

	{
		var btn = ui.NewFontButton()
		btn.OnChanged(func(fb *ui.FontButton) {
			log.Println(fb.Font())
		})
		vbox.Append(btn, false)
	}

	{
		var box = ui.NewCombobox()
		box.Append("r")
		box.Append("g")
		box.Append("b")
		box.SetSelected(0)
		box.OnSelected(func(c *ui.Combobox) {
			log.Println(c.Selected())
		})
		vbox.Append(box, false)
	}

	{
		var btn = ui.NewButton("alert")
		btn.OnClicked(func(b *ui.Button) {
			ui.MsgBox(win, "tips", "ruok")
		})
		vbox.Append(btn, false)
	}

	{
		var form = ui.NewForm()
		form.SetPadded(true)
		vbox.Append(form, false)

		var name = ui.NewEntry()
		form.Append("name", name, false)

		var pwd = ui.NewPasswordEntry()
		form.Append("pwd", pwd, false)

		var btn = ui.NewButton("confirm")
		btn.OnClicked(func(b *ui.Button) {
			log.Println(name.Text() + "," + pwd.Text())
		})
		form.Append("", btn, false)
	}

	win.Show()
}
