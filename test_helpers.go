package selenium

import (
	"testing"
)

// A single-return-value interface to WebDriverT that is useful when using WebDrivers in test code.
// Obtain a WebDriverT by calling webDriver.T(t), where t *testing.T is the test handle for the
// current test. The methods of WebDriverT call wt.t.Errorf upon encountering errors instead of using
// multiple returns to indicate errors.
type WebDriverT interface {
	NewSession() string

	SetAsyncScriptTimeout(ms uint)
	SetImplicitWaitTimeout(ms uint)
	Quit()

	CurrentWindowHandle() string
	WindowHandles() []string
	CurrentURL() string
	Title() string
	PageSource() string
	Close()
	SwitchFrame(frame string)
	SwitchWindow(name string)
	CloseWindow(name string)

	Get(url string)
	Forward()
	Back()
	Refresh()

	FindElement(by, value string) WebElementT
	FindElements(by, value string) []WebElementT
	ActiveElement() WebElement

	GetCookies() []Cookie
	AddCookie(cookie *Cookie)
	DeleteAllCookies()
	DeleteCookie(name string)

	Click(button int)
	DoubleClick()
	ButtonDown()
	ButtonUp()

	SendModifier(modifier string, isDown bool)
	Screenshot() []byte

	DismissAlert()
	AcceptAlert()
	AlertText() string
	SetAlertText(text string)

	ExecuteScript(script string, args []interface{}) interface{}
	ExecuteScriptAsync(script string, args []interface{}) interface{}
}

type webDriverT struct {
	d WebDriver
	t *testing.T
}

func (wt *webDriverT) NewSession() (id string) {
	var err error
	if id, err = wt.d.NewSession(); err != nil {
		wt.t.Errorf("NewSession: %s", err)
	}
	return
}

func (wt *webDriverT) SetAsyncScriptTimeout(ms uint) {
	if err := wt.d.SetAsyncScriptTimeout(ms); err != nil {
		wt.t.Errorf("SetAsyncScriptTimeout(%d msec): %s", ms, err)
	}
}

func (wt *webDriverT) SetImplicitWaitTimeout(ms uint) {
	if err := wt.d.SetImplicitWaitTimeout(ms); err != nil {
		wt.t.Errorf("SetImplicitWaitTimeout(%d msec): %s", ms, err)
	}
}

func (wt *webDriverT) Quit() {
	if err := wt.d.Quit(); err != nil {
		wt.t.Errorf("Quit: %s", err)
	}
}

func (wt *webDriverT) CurrentWindowHandle() (v string) {
	var err error
	if v, err = wt.d.CurrentWindowHandle(); err != nil {
		wt.t.Errorf("CurrentWindowHandle: %s", err)
	}
	return
}

func (wt *webDriverT) WindowHandles() (hs []string) {
	var err error
	if hs, err = wt.d.WindowHandles(); err != nil {
		wt.t.Errorf("WindowHandles: %s", err)
	}
	return
}

func (wt *webDriverT) CurrentURL() (v string) {
	var err error
	if v, err = wt.d.CurrentURL(); err != nil {
		wt.t.Errorf("CurrentURL: %s", err)
	}
	return
}

func (wt *webDriverT) Title() (v string) {
	var err error
	if v, err = wt.d.Title(); err != nil {
		wt.t.Errorf("Title: %s", err)
	}
	return
}

func (wt *webDriverT) PageSource() (v string) {
	var err error
	if v, err = wt.d.PageSource(); err != nil {
		wt.t.Errorf("PageSource: %s", err)
	}
	return
}

func (wt *webDriverT) Close() {
	if err := wt.d.Close(); err != nil {
		wt.t.Errorf("Close: %s", err)
	}
}

func (wt *webDriverT) SwitchFrame(frame string) {
	if err := wt.d.SwitchFrame(frame); err != nil {
		wt.t.Errorf("SwitchFrame(%q): %s", frame, err)
	}
}

func (wt *webDriverT) SwitchWindow(name string) {
	if err := wt.d.SwitchWindow(name); err != nil {
		wt.t.Errorf("SwitchWindow(%q): %s", name, err)
	}
}

func (wt *webDriverT) CloseWindow(name string) {
	if err := wt.d.CloseWindow(name); err != nil {
		wt.t.Errorf("CloseWindow(%q): %s", name, err)
	}
}

func (wt *webDriverT) Get(name string) {
	if err := wt.d.Get(name); err != nil {
		wt.t.Errorf("Get(%q): %s", name, err)
	}
}

func (wt *webDriverT) Forward() {
	if err := wt.d.Forward(); err != nil {
		wt.t.Errorf("Forward: %s", err)
	}
}

func (wt *webDriverT) Back() {
	if err := wt.d.Back(); err != nil {
		wt.t.Errorf("Back: %s", err)
	}
}

func (wt *webDriverT) Refresh() {
	if err := wt.d.Refresh(); err != nil {
		wt.t.Errorf("Refresh: %s", err)
	}
}

func (wt *webDriverT) FindElement(by, value string) (elem WebElementT) {
	if elem_, err := wt.d.FindElement(by, value); err == nil {
		elem = elem_.T(wt.t)
	} else {
		wt.t.Errorf("FindElement(by=%s, value=%s): %s", by, value, err)
	}
	return
}

func (wt *webDriverT) FindElements(by, value string) (elems []WebElementT) {
	if elems_, err := wt.d.FindElements(by, value); err == nil {
		for _, elem := range elems_ {
			elems = append(elems, elem.T(wt.t))
		}
	} else {
		wt.t.Errorf("FindElements(by=%s, value=%s): %s", by, value, err)
	}
	return
}

func (wt *webDriverT) ActiveElement() (elem WebElement) {
	var err error
	if elem, err = wt.d.ActiveElement(); err != nil {
		wt.t.Errorf("ActiveElement: %s", err)
	}
	return
}

func (wt *webDriverT) GetCookies() (c []Cookie) {
	var err error
	if c, err = wt.d.GetCookies(); err != nil {
		wt.t.Errorf("GetCookies: %s", err)
	}
	return
}

func (wt *webDriverT) AddCookie(cookie *Cookie) {
	if err := wt.d.AddCookie(cookie); err != nil {
		wt.t.Errorf("AddCookie(%+q): %s", cookie, err)
	}
	return
}

func (wt *webDriverT) DeleteAllCookies() {
	if err := wt.d.DeleteAllCookies(); err != nil {
		wt.t.Errorf("DeleteAllCookies: %s", err)
	}
}

func (wt *webDriverT) DeleteCookie(name string) {
	if err := wt.d.DeleteCookie(name); err != nil {
		wt.t.Errorf("DeleteCookie(%q): %s", name, err)
	}
}

func (wt *webDriverT) Click(button int) {
	if err := wt.d.Click(button); err != nil {
		wt.t.Errorf("Click(%d): %s", button, err)
	}
}

func (wt *webDriverT) DoubleClick() {
	if err := wt.d.DoubleClick(); err != nil {
		wt.t.Errorf("DoubleClick: %s", err)
	}
}

func (wt *webDriverT) ButtonDown() {
	if err := wt.d.ButtonDown(); err != nil {
		wt.t.Errorf("ButtonDown: %s", err)
	}
}

func (wt *webDriverT) ButtonUp() {
	if err := wt.d.ButtonUp(); err != nil {
		wt.t.Errorf("ButtonUp: %s", err)
	}
}

func (wt *webDriverT) SendModifier(modifier string, isDown bool) {
	if err := wt.d.SendModifier(modifier, isDown); err != nil {
		wt.t.Errorf("SendModifier(modifier=%q, isDown=%s): %s", modifier, isDown, err)
	}
}

func (wt *webDriverT) Screenshot() (data []byte) {
	var err error
	if data, err = wt.d.Screenshot(); err != nil {
		wt.t.Errorf("Screenshot: %s", err)
	}
	return
}

func (wt *webDriverT) DismissAlert() {
	if err := wt.d.DismissAlert(); err != nil {
		wt.t.Errorf("DismissAlert: %s", err)
	}
}

func (wt *webDriverT) AcceptAlert() {
	if err := wt.d.AcceptAlert(); err != nil {
		wt.t.Errorf("AcceptAlert: %s", err)
	}
}

func (wt *webDriverT) AlertText() (text string) {
	var err error
	if text, err = wt.d.AlertText(); err != nil {
		wt.t.Errorf("AlertText: %s", err)
	}
	return
}

func (wt *webDriverT) SetAlertText(text string) {
	var err error
	if err = wt.d.SetAlertText(text); err != nil {
		wt.t.Errorf("SetAlertText(%q): %s", text, err)
	}
}

func (wt *webDriverT) ExecuteScript(script string, args []interface{}) (res interface{}) {
	var err error
	if res, err = wt.d.ExecuteScript(script, args); err != nil {
		wt.t.Errorf("ExecuteScript(script=%q, args=%+q): %s", script, args, err)
	}
	return
}

func (wt *webDriverT) ExecuteScriptAsync(script string, args []interface{}) (res interface{}) {
	var err error
	if res, err = wt.d.ExecuteScriptAsync(script, args); err != nil {
		wt.t.Errorf("ExecuteScriptAsync(script=%q, args=%+q): %s", script, args, err)
	}
	return
}

// A single-return-value interface to WebElement that is useful when using WebElements in test code.
// Obtain a WebElementT by calling webElement.T(t), where t *testing.T is the test handle for the
// current test. The methods of WebElementT call wt.t.Errorf upon encountering errors instead of using
// multiple returns to indicate errors.
type WebElementT interface {
	Click()
	SendKeys(keys string)
	Submit()
	Clear()
	MoveTo(xOffset, yOffset int)

	FindElement(by, value string) WebElement
	FindElements(by, value string) []WebElement

	TagName() string
	Text() string
	IsSelected() bool
	IsEnabled() bool
	IsDisplayed() bool
	GetAttribute(name string) string
	Location() *Point
	LocationInView() *Point
	Size() *Size
	CSSProperty(name string) string
}

type webElementT struct {
	e WebElement
	t *testing.T
}

func (wt *webElementT) Click() {
	if err := wt.e.Click(); err != nil {
		wt.t.Errorf("Click: %s", err)
	}
}

func (wt *webElementT) SendKeys(keys string) {
	if err := wt.e.SendKeys(keys); err != nil {
		wt.t.Errorf("SendKeys(%q): %s", keys, err)
	}
}

func (wt *webElementT) Submit() {
	if err := wt.e.Submit(); err != nil {
		wt.t.Errorf("Submit: %s", err)
	}
}

func (wt *webElementT) Clear() {
	if err := wt.e.Clear(); err != nil {
		wt.t.Errorf("Clear: %s", err)
	}
}

func (wt *webElementT) MoveTo(xOffset, yOffset int) {
	if err := wt.e.MoveTo(xOffset, yOffset); err != nil {
		wt.t.Errorf("MoveTo(xOffset=%d, yOffset=%d): %s", xOffset, yOffset, err)
	}
}

func (wt *webElementT) FindElement(by, value string) (elem WebElement) {
	var err error
	if elem, err = wt.e.FindElement(by, value); err != nil {
		wt.t.Errorf("FindElement(by=%s, value=%s): %s", by, value, err)
	}
	return
}

func (wt *webElementT) FindElements(by, value string) (elems []WebElement) {
	var err error
	if elems, err = wt.e.FindElements(by, value); err != nil {
		wt.t.Errorf("FindElements(by=%s, value=%s): %s", by, value, err)
	}
	return
}

func (wt *webElementT) TagName() (v string) {
	var err error
	if v, err = wt.e.TagName(); err != nil {
		wt.t.Errorf("TagName: %s", err)
	}
	return
}

func (wt *webElementT) Text() (v string) {
	var err error
	if v, err = wt.e.Text(); err != nil {
		wt.t.Errorf("Text: %s", err)
	}
	return
}

func (wt *webElementT) IsSelected() (v bool) {
	var err error
	if v, err = wt.e.IsSelected(); err != nil {
		wt.t.Errorf("IsSelected: %s", err)
	}
	return
}

func (wt *webElementT) IsEnabled() (v bool) {
	var err error
	if v, err = wt.e.IsEnabled(); err != nil {
		wt.t.Errorf("IsEnabled: %s", err)
	}
	return
}

func (wt *webElementT) IsDisplayed() (v bool) {
	var err error
	if v, err = wt.e.IsDisplayed(); err != nil {
		wt.t.Errorf("IsDisplayed: %s", err)
	}
	return
}

func (wt *webElementT) GetAttribute(name string) (v string) {
	var err error
	if v, err = wt.e.GetAttribute(name); err != nil {
		wt.t.Errorf("GetAttribute(%q): %s", name, err)
	}
	return
}

func (wt *webElementT) Location() (v *Point) {
	var err error
	if v, err = wt.e.Location(); err != nil {
		wt.t.Errorf("Location: %s", err)
	}
	return
}

func (wt *webElementT) LocationInView() (v *Point) {
	var err error
	if v, err = wt.e.LocationInView(); err != nil {
		wt.t.Errorf("LocationInView: %s", err)
	}
	return
}

func (wt *webElementT) Size() (v *Size) {
	var err error
	if v, err = wt.e.Size(); err != nil {
		wt.t.Errorf("Size: %s", err)
	}
	return
}

func (wt *webElementT) CSSProperty(name string) (v string) {
	var err error
	if v, err = wt.e.CSSProperty(name); err != nil {
		wt.t.Errorf("CSSProperty(%q): %s", name, err)
	}
	return
}
