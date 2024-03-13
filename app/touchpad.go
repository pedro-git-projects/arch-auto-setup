package app

import (
	"fmt"
	"os"
	"os/exec"
)

func (app *App) enableTouchpad() {
	p := "/etc/X11/xorg.conf.d/30-touchpad.conf"

	content := `Section "InputClass"
        Identifier "MyTouchpad"
        MatchIsTouchpad "on"
        Driver "libinput"
        Option "Tapping" "on"
EndSection`

	cmd := exec.Command("sudo", "sh", "-c", fmt.Sprintf(`echo '%s' | sudo tee %s > /dev/null`, content, p))
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Printf("failed to enable touchpad: %s\n", err)
		return
	}

	fmt.Printf("file %s created successfully\n", p)
}
