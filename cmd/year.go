package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type File struct {
	name string
	path string
}

func yearTable(url string) {
	for{
		action := func() {
			time.Sleep(2 * time.Second)
		}
		if err := spinner.New().Title("Fetching years").Action(action).Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		files, err := yearReq(url)
		if err != nil {
			fmt.Print(errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
			return
		}

		var selectedOption string
		var options []huh.Option[string]
		var fileList []File

		// Convert files to huh options.
		for _, file := range files {
			fileItem := File(file)
			fileList = append(fileList, fileItem)
			options = append(options, huh.NewOption(fileItem.name, fileItem.path))
		}
		// Add back option.
		options = append(options, huh.NewOption("Back to Main Menu", "back"))
		options = append(options, huh.NewOption("Quit", "quit"))

		// Create the form.
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Select Year").
					Options(options...).
					Value(&selectedOption),
			),
		)

		// Run the form.
		err = form.Run()
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}

		// Handle selection.
		switch selectedOption {
		case "back":
			huhMenuStart() // Go back to main menu.
		case "quit":
			fmt.Println(fetchStatusStyle.Render("Exiting..."))
			os.Exit(0)
		default:
			// Find selected file and process it
			for _, fileItem := range fileList {
				if fileItem.path == selectedOption {
					url := BASE_URL + fileItem.path
					openBrowser(url) // Function to open the browser with the selected URL.
					break
				}
			}
		}
	}
}

func year(url string) {
	yearTable(url) // Call the yearTable function to display the menu.
}