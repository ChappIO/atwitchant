package open

import "os/exec"

func Browse(url string) error {
	return exec.Command("xdg-open", url).Start()
}
