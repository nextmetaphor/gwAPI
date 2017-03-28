package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jroimartin/gocui"
)

func main() {
	gui, guiError := gocui.NewGui(gocui.OutputNormal)
	if (guiError != nil) {
		log.WithFields(log.Fields{
			"error": guiError}).Error("Error creating gui.")
		return;
	}
	defer gui.Close()

	gui.Cursor = true


}
