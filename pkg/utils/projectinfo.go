package utils

import (
	"fmt"
)

// ANSI color codes
const (
	Reset  = "\033[0m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func ProjectInfo() {
	header := "DockerManager"
	divider := "----------------"

	fmt.Println()

	fmt.Printf("%s%s\n%s%s%s\n%s%s\n\n", Cyan, divider, Purple, header, Reset, Cyan, divider)

	fmt.Printf(
		"%s🚀 Welcome to DockerManager! Making Docker management fun and easy! 🐳%s\n",
		White, Reset,
	)

	fmt.Printf(
		"%s📍 Made by: tenderpanini%s\n\n",
		White, Reset,
	)

	fmt.Println()
}

func ExitInfo() {
	fmt.Println()
	fmt.Println("👋 Thank you for using Docker Manager. See you next time!")
	fmt.Println()
}
