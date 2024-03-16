package input

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/libretro/ludo/libretro"
)

var keyBinds = map[glfw.Key]uint32{
	glfw.KeyA:      libretro.DeviceIDJoypadA,
	glfw.KeyS:      libretro.DeviceIDJoypadB,
	glfw.KeyQ:      libretro.DeviceIDJoypadY,
	glfw.KeyW:      libretro.DeviceIDJoypadX,
	glfw.KeyL:      libretro.DeviceIDJoypadL,
	glfw.KeyR:      libretro.DeviceIDJoypadR,
	glfw.KeyUp:     libretro.DeviceIDJoypadUp,
	glfw.KeyDown:   libretro.DeviceIDJoypadDown,
	glfw.KeyLeft:   libretro.DeviceIDJoypadLeft,
	glfw.KeyRight:  libretro.DeviceIDJoypadRight,
	glfw.KeyX:      libretro.DeviceIDJoypadStart,
	glfw.KeyZ:      libretro.DeviceIDJoypadSelect,
	glfw.KeySpace:  ActionFastForwardToggle,
	glfw.KeyM:      ActionMenuToggle,
	glfw.KeyF:      ActionFullscreenToggle,
	glfw.KeyEscape: ActionShouldClose,
}
