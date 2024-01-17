//go:build !js && !wasm && !test_web_driver
// +build !js,!wasm,!test_web_driver

package glfw

import (
	"fmt"

	"fyne.io/fyne/v2"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func (d *gLDriver) initGLFW() {
	initOnce.Do(func() {
		err := glfw.Init()
		if err != nil {
			fyne.LogError("failed to initialise GLFW", err)
			return
		}

		initCursors()
		d.startDrawThread()
	})
}

func (d *gLDriver) waitEvents() <-chan struct{} {
	defer func() {
		if r := recover(); r != nil {
			fyne.LogError(fmt.Sprint("GLFW poll event error: ", r), nil)
		}
	}()

	ch := make(chan struct{})

	go func() {
		defer func() { close(ch) }()
		for {
			glfw.WaitEvents()
			ch <- struct{}{}
		}
	}()

	return ch
}

func (d *gLDriver) Terminate() {
	glfw.Terminate()
}
