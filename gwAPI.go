package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jroimartin/gocui"
	"github.com/nextmetaphor/gwAPI/connection"
	"github.com/nextmetaphor/gwAPI/controller"
	"gopkg.in/square/go-jose.v1/json"
	"net/http"
	"os"
	"strconv"
)

const (
	logo = "" +
		" \x1b[36m┌─┐┬ ┬\x1b[37m╔═╗╔═╗╦ \n" +
		" \x1b[36m│ ┬│││\x1b[37m╠═╣╠═╝║ \n" +
		" \x1b[36m└─┘└┴┘\x1b[37m╩ ╩╩  ╩ "

	apiSummaryView     = "apis"
	apiDetailView      = "apiDetails"
	sessionSummaryView = "sessions"
)

var credentials = connection.ConnectionCredentials{
	GatewayURL: "http://192.168.64.8:30002",
	AuthToken:  "ThisInNotTheSecretYouAreLookingFor"}

var connector = connection.Connection{}

func layout(g *gocui.Gui) error {
	maxX, _ := g.Size()

	// create the "logo" view
	if view, viewErr := g.SetView("logo", 0, 0, 16, 4); viewErr != nil {
		if viewErr != gocui.ErrUnknownView {
			return viewErr
		}
		view.Frame = true

		fmt.Fprintln(view, logo)
	}

	if v, err := g.SetView("domain", 19, 1, 100, 3); err != nil {
		fmt.Fprintln(v, "http://localhost:8080")

		v.FgColor = gocui.ColorCyan
		v.Frame = true
		v.Title = "Dashboard URL [Connected] "
		v.Editable = true
		g.SetCurrentView("domain")
	}

	if v, err := g.SetView("side", 0, 6, 16, 14); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Frame = true
		v.Title = "Objects"

		fmt.Fprintln(v, "APIs")
		fmt.Fprintln(v, "Policies")
		fmt.Fprintln(v, "Keys")

	}

	if v, err := g.SetView(apiSummaryView, 19, 6, 200, 26); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = true
		v.Title = "APIs"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

	}

	if v, err := g.SetView(sessionSummaryView, 19, 28, 52, 42); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = true
		v.Title = "sessions"
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

	}

	log.Debug(maxX)

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()

		log.Debug("HERE", oy, cy)
		getKeys(g, "1")
		if (oy == 0) && (cy == 1) {
			return nil
		} else if (oy == 1) && (cy == 1) {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}

		}
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func login(gui *gocui.Gui, view *gocui.View) error {
	maxX, maxY := gui.Size()
	if v, err := gui.SetView("login", maxX/2-30, maxY/2, maxX/2+30, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = true
		v.Title = "Enter Auth Key for Domain"
		v.Autoscroll = false
		v.Wrap = false
		fmt.Fprintln(v, "Snouts")
		if _, err := gui.SetCurrentView("login"); err != nil {
			return err
		}
	}
	return nil
}

func cancelAuthenticationView(gui *gocui.Gui, view *gocui.View) error {
	err := gui.DeleteView("login")
	gui.SetCurrentView("domain")

	return err

}

func saveAPI(gui *gocui.Gui, view *gocui.View) error {
	view.Title = view.Title + " Saving API..."

	bufferString := view.Buffer()
	_, httpResponse, updateErr := controller.UpdateAPI(credentials, connector, "1", &bufferString)

	if updateErr != nil {
		return updateErr
	}

	if httpResponse.StatusCode == http.StatusOK {
		view.Title = view.Title + "API Saved. Reloading definitions..."
		controller.ReloadGatewayGroup(credentials, connector)

		view.Title = view.Title + "Definitions reloaded"
	}

	return nil
}

func selectAPI(gui *gocui.Gui, view *gocui.View) error {
	_, cy := view.Cursor()
	//var line string
	var err error
	if _, err = view.Line(cy); err != nil {
		//TODO
		return err
	}
	maxX, maxY := gui.Size()
	if v, err := gui.SetView(apiDetailView, 10, 10, maxX-10, maxY-10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = true
		v.Title = "API Definition"
		v.Autoscroll = false
		v.Wrap = true

		api, _, loadAPIErr := controller.ReadAPI(credentials, connector, "1")
		if loadAPIErr != nil {
			panic(loadAPIErr)
		}
		jsonAPI, marshallErr := json.MarshalIndent(api, "", "  ")
		if marshallErr != nil {
			panic(marshallErr)
		}

		fmt.Fprintln(v, string(jsonAPI))

		if _, err := gui.SetCurrentView(apiDetailView); err != nil {
			return err
		}
	}

	return nil
}

//func showNodes(gui *gocui.Gui, view *gocui.View) error {
//nodeHealth := new(tykcommon.)
//}

func getKeys(gui *gocui.Gui, apiID string) {
	keys, _, err := controller.SelectKeys(credentials, connector, apiID)
	if err != nil {
		panic(err)
	}

	keyView, _ := gui.View(sessionSummaryView)

	keyView.Clear()
	const outputFormat = "%-32.32s\n"
	for _, key := range keys.APIKeys {
		fmt.Fprintf(keyView, outputFormat, key)
	}

}

func attemptLogin(gui *gocui.Gui, view *gocui.View) error {
	err := gui.DeleteView("login")

	apis, _, _ := controller.SelectAPIs(credentials, connector)
	apiView, apiViewErr := gui.View("apis")

	if apiViewErr != nil {
		log.Fatal(apiViewErr)
		return apiViewErr
	}

	const outputFormat = "%-6.6s  %-32.32s  %-24.24s  %-32.32s  %-24.24s  %-32.32s  %-100.100s\n"
	fmt.Fprintf(apiView, outputFormat, "\x1b[36m", "NAME", "ID", "API-ID", "ORG-ID", "LISTEN-PATH", "TARGET-URL")

	for index, api := range *apis {
		fmt.Fprintf(
			apiView,
			outputFormat,
			"\x1b[37m"+strconv.Itoa(index+1),
			api.Name,
			"0",
			api.APIID,
			api.OrgID,
			api.Proxy.ListenPath,
			api.Proxy.TargetURL,
		)
	}

	gui.SetCurrentView("apis")
	return err
}

func cancelEditAPI(gui *gocui.Gui, view *gocui.View) error {
	err := gui.DeleteView(apiDetailView)
	gui.SetCurrentView("apis")

	return err
}

func toggleWindows(gui *gocui.Gui, view *gocui.View) error {
	var err error
	if view.Name() == apiSummaryView {
		_, err = gui.SetCurrentView(sessionSummaryView)
	} else {
		_, err = gui.SetCurrentView(apiSummaryView)
	}

	return err
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", 'c', gocui.ModNone, quit); err != nil {
		return err
	}
	//if err := g.SetKeybinding("side", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
	//	return err
	//}
	//if err := g.SetKeybinding("main", gocui.KeyCtrlSpace, gocui.ModNone, nextView); err != nil {
	//	return err
	//}
	if err := g.SetKeybinding("side", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("domain", gocui.KeyEnter, gocui.ModNone, login); err != nil {
		return err
	}
	if err := g.SetKeybinding("login", gocui.KeyArrowDown, gocui.ModNone, nil); err != nil {
		return err
	}
	if err := g.SetKeybinding("login", gocui.KeyEnter, gocui.ModNone, attemptLogin); err != nil {
		return err
	}

	if err := g.SetKeybinding("login", gocui.KeyEsc, gocui.ModNone, cancelAuthenticationView); err != nil {
		return err
	}

	if err := g.SetKeybinding("apis", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("apis", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("apis", gocui.KeyEnter, gocui.ModNone, selectAPI); err != nil {
		return err
	}

	if err := g.SetKeybinding(apiDetailView, gocui.KeyEsc, gocui.ModNone, cancelEditAPI); err != nil {
		return err
	}

	if err := g.SetKeybinding(apiDetailView, gocui.KeyCtrlS, gocui.ModNone, saveAPI); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, toggleWindows); err != nil {
		return err
	}

	//	if err := g.SetKeybinding("side", gocui.KeyEnter, gocui.ModNone, getLine); err != nil {
	//		return err
	//	}
	//	if err := g.SetKeybinding("msg", gocui.KeyEnter, gocui.ModNone, delMsg); err != nil {
	//		return err
	//	}
	//
	//	if err := g.SetKeybinding("main", gocui.KeyCtrlS, gocui.ModNone, saveMain); err != nil {
	//		return err
	//	}
	//	if err := g.SetKeybinding("main", gocui.KeyCtrlW, gocui.ModNone, saveVisualMain); err != nil {
	//		return err
	//	}
	return nil
}

func main() {
	f, err := os.OpenFile("gwAPI.log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	log.SetLevel(log.DebugLevel)
	defer f.Close()

	gui, guiError := gocui.NewGui(gocui.Output256)
	gui.InputEsc = true
	if guiError != nil {
		log.WithFields(log.Fields{
			"error": guiError}).Debug("Error creating gui.")
		return
	}
	defer gui.Close()

	gui.Cursor = true
	gui.SetManagerFunc(layout)

	gui.SelFgColor = gocui.ColorGreen
	gui.Highlight = true

	if err := keybindings(gui); err != nil {
		log.Panicln(err)
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
