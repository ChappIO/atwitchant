package open

import "os/exec"

func Browse(url string) error {
	return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
}
