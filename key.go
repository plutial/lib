package main

import (
	// GLFW
	"github.com/go-gl/glfw/v3.3/glfw"
)

// Key types
const (
	KeyUnknown      int = int(glfw.KeyUnknown)
	KeySpace        int = int(glfw.KeySpace)
	KeyApostrophe   int = int(glfw.KeyApostrophe)
	KeyComma        int = int(glfw.KeyComma)
	KeyMinus        int = int(glfw.KeyMinus)
	KeyPeriod       int = int(glfw.KeyPeriod)
	KeySlash        int = int(glfw.KeySlash)
	Key0            int = int(glfw.Key0)
	Key1            int = int(glfw.Key1)
	Key2            int = int(glfw.Key2)
	Key3            int = int(glfw.Key3)
	Key4            int = int(glfw.Key4)
	Key5            int = int(glfw.Key5)
	Key6            int = int(glfw.Key6)
	Key7            int = int(glfw.Key7)
	Key8            int = int(glfw.Key8)
	Key9            int = int(glfw.Key9)
	KeySemicolon    int = int(glfw.KeySemicolon)
	KeyEqual        int = int(glfw.KeyEqual)
	KeyA            int = int(glfw.KeyA)
	KeyB            int = int(glfw.KeyB)
	KeyC            int = int(glfw.KeyC)
	KeyD            int = int(glfw.KeyD)
	KeyE            int = int(glfw.KeyE)
	KeyF            int = int(glfw.KeyF)
	KeyG            int = int(glfw.KeyG)
	KeyH            int = int(glfw.KeyH)
	KeyI            int = int(glfw.KeyI)
	KeyJ            int = int(glfw.KeyJ)
	KeyK            int = int(glfw.KeyK)
	KeyL            int = int(glfw.KeyL)
	KeyM            int = int(glfw.KeyM)
	KeyN            int = int(glfw.KeyN)
	KeyO            int = int(glfw.KeyO)
	KeyP            int = int(glfw.KeyP)
	KeyQ            int = int(glfw.KeyQ)
	KeyR            int = int(glfw.KeyR)
	KeyS            int = int(glfw.KeyS)
	KeyT            int = int(glfw.KeyT)
	KeyU            int = int(glfw.KeyU)
	KeyV            int = int(glfw.KeyV)
	KeyW            int = int(glfw.KeyW)
	KeyX            int = int(glfw.KeyX)
	KeyY            int = int(glfw.KeyY)
	KeyZ            int = int(glfw.KeyZ)
	KeyLeftBracket  int = int(glfw.KeyLeftBracket)
	KeyBackslash    int = int(glfw.KeyBackslash)
	KeyRightBracket int = int(glfw.KeyRightBracket)
	KeyGraveAccent  int = int(glfw.KeyGraveAccent)
	KeyWorld1       int = int(glfw.KeyWorld1)
	KeyWorld2       int = int(glfw.KeyWorld2)
	KeyEscape       int = int(glfw.KeyEscape)
	KeyEnter        int = int(glfw.KeyEnter)
	KeyTab          int = int(glfw.KeyTab)
	KeyBackspace    int = int(glfw.KeyBackspace)
	KeyInsert       int = int(glfw.KeyInsert)
	KeyDelete       int = int(glfw.KeyDelete)
	KeyRight        int = int(glfw.KeyRight)
	KeyLeft         int = int(glfw.KeyLeft)
	KeyDown         int = int(glfw.KeyDown)
	KeyUp           int = int(glfw.KeyUp)
	KeyPageUp       int = int(glfw.KeyPageUp)
	KeyPageDown     int = int(glfw.KeyPageDown)
	KeyHome         int = int(glfw.KeyHome)
	KeyEnd          int = int(glfw.KeyEnd)
	KeyCapsLock     int = int(glfw.KeyCapsLock)
	KeyScrollLock   int = int(glfw.KeyScrollLock)
	KeyNumLock      int = int(glfw.KeyNumLock)
	KeyPrintScreen  int = int(glfw.KeyPrintScreen)
	KeyPause        int = int(glfw.KeyPause)
	KeyF1           int = int(glfw.KeyF1)
	KeyF2           int = int(glfw.KeyF2)
	KeyF3           int = int(glfw.KeyF3)
	KeyF4           int = int(glfw.KeyF4)
	KeyF5           int = int(glfw.KeyF5)
	KeyF6           int = int(glfw.KeyF6)
	KeyF7           int = int(glfw.KeyF7)
	KeyF8           int = int(glfw.KeyF8)
	KeyF9           int = int(glfw.KeyF9)
	KeyF10          int = int(glfw.KeyF10)
	KeyF11          int = int(glfw.KeyF11)
	KeyF12          int = int(glfw.KeyF12)
	KeyF13          int = int(glfw.KeyF13)
	KeyF14          int = int(glfw.KeyF14)
	KeyF15          int = int(glfw.KeyF15)
	KeyF16          int = int(glfw.KeyF16)
	KeyF17          int = int(glfw.KeyF17)
	KeyF18          int = int(glfw.KeyF18)
	KeyF19          int = int(glfw.KeyF19)
	KeyF20          int = int(glfw.KeyF20)
	KeyF21          int = int(glfw.KeyF21)
	KeyF22          int = int(glfw.KeyF22)
	KeyF23          int = int(glfw.KeyF23)
	KeyF24          int = int(glfw.KeyF24)
	KeyF25          int = int(glfw.KeyF25)
	KeyKP0          int = int(glfw.KeyKP0)
	KeyKP1          int = int(glfw.KeyKP1)
	KeyKP2          int = int(glfw.KeyKP2)
	KeyKP3          int = int(glfw.KeyKP3)
	KeyKP4          int = int(glfw.KeyKP4)
	KeyKP5          int = int(glfw.KeyKP5)
	KeyKP6          int = int(glfw.KeyKP6)
	KeyKP7          int = int(glfw.KeyKP7)
	KeyKP8          int = int(glfw.KeyKP8)
	KeyKP9          int = int(glfw.KeyKP9)
	KeyKPDecimal    int = int(glfw.KeyKPDecimal)
	KeyKPDivide     int = int(glfw.KeyKPDivide)
	KeyKPMultiply   int = int(glfw.KeyKPMultiply)
	KeyKPSubtract   int = int(glfw.KeyKPSubtract)
	KeyKPAdd        int = int(glfw.KeyKPAdd)
	KeyKPEnter      int = int(glfw.KeyKPEnter)
	KeyKPEqual      int = int(glfw.KeyKPEqual)
	KeyLeftShift    int = int(glfw.KeyLeftShift)
	KeyLeftControl  int = int(glfw.KeyLeftControl)
	KeyLeftAlt      int = int(glfw.KeyLeftAlt)
	KeyLeftSuper    int = int(glfw.KeyLeftSuper)
	KeyRightShift   int = int(glfw.KeyRightShift)
	KeyRightControl int = int(glfw.KeyRightControl)
	KeyRightAlt     int = int(glfw.KeyRightAlt)
	KeyRightSuper   int = int(glfw.KeyRightSuper)
	KeyMenu         int = int(glfw.KeyMenu)
	KeyLast         int = int(glfw.KeyLast)
)

// All keys
type Key struct {
	Pressed       bool
	PressedRepeat bool
	Released      bool
	Down          bool

	// Update info
	UpdatedFrame int
}

var keys [KeyLast]Key

// Key Callback to update the keys
func KeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	// New update
	keys[key].UpdatedFrame = frameCount

	// Reset all actions except down
	keys[key].Pressed = false
	keys[key].PressedRepeat = false
	keys[key].Released = false

	switch action {
	// Key pressed
	case glfw.Press:
		keys[key].Pressed = true

		// Set key down to true
		keys[key].Down = true

	// Key pressed repeat
	case glfw.Repeat:
		keys[key].PressedRepeat = true

		// Set key down to true
		keys[key].Down = true

	// Key released
	case glfw.Release:
		keys[key].Released = true

		// Set key down to false
		keys[key].Down = false

	default:
		// Reset all actions except down
		keys[key].Pressed = false
		keys[key].PressedRepeat = false
		keys[key].Released = false
	}
}

// Key events
func IsKeyDown(key int) bool {
	return keys[key].Down
}

func IsKeyPressed(key int) bool {
	if keys[key].UpdatedFrame == frameCount && keys[key].Pressed {
		return true
	} else {
		keys[key].Pressed = false
		return false
	}
}

func IsKeyPressedRepeat(key int) bool {
	if keys[key].UpdatedFrame == frameCount && keys[key].PressedRepeat {
		return true
	} else {
		keys[key].PressedRepeat = false
		return false
	}
}

func IsKeyReleased(key int) bool {
	if keys[key].UpdatedFrame == frameCount && keys[key].Released {
		return true
	} else {
		keys[key].Released = false
		return false
	}
}
