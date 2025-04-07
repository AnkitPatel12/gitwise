package main

import(
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("git", "log", "--pretty=format:%h | %an | %s", "-n", "5")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running git command:", err)
		return
	}

	commits := strings.Split(string(output), "\n")

	fmt.Println("Recent commits:")
	for _, commit := range commits {
		fmt.Println(commit)
	}
}