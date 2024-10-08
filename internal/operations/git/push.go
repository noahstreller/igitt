package git

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/noahstreller/igitt/internal/utilities"
	"github.com/noahstreller/igitt/internal/utilities/logger"
)

func PushRemote() {
	fmt.Println("Pushing to remote repository")
	progressIndicator := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	progressIndicator.Start()
	byteOut, errOut := exec.Command("git", "push").CombinedOutput()
	progressIndicator.Stop()

	if errOut != nil {
		logger.ErrorLogger.Println("Error pushing to remote repository:", errOut, string(byteOut))
		utilities.PrintError(string(byteOut))
		return
	}
	logger.InfoLogger.Println("Pushing to remote repository:", errOut, string(byteOut))
}
