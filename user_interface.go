package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"fmt"
)

func main() {
	a := app.New()
	w := a.NewWindow("Login Form")
	w.Resize(fyne.NewSize(600, 400)) // Set the window size to 600x400

	usernameEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()

	loginForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Username", Widget: usernameEntry},
			{Text: "Password", Widget: passwordEntry},
		},
		OnSubmit: func() {
			if usernameEntry.Text == "admin" && passwordEntry.Text == "admin" {
				w.Close()
				showMainScreen(a)
			} else {
				dialog.ShowError(fmt.Errorf("Invalid credentials"), w)
			}
		},
                OnCancel: func() {
                	w.Close()
                },
	}

	w.SetContent(container.NewVBox(loginForm))
	w.ShowAndRun()
}

func showMainScreen(a fyne.App) {
	w := a.NewWindow("Reserve Form")
	w.Resize(fyne.NewSize(600, 400)) // Set the window size to 600x400

	meetingRoomButton := widget.NewButton("Meeting Room", func() {
		showMeetingForm(a)
	})

	cubicleButton := widget.NewButton("Cubicle", func() {
                showCubeForm(a)
	})

	cancelButton := widget.NewButton("Cancel", func() {
		w.Close()
	})
	cancelButton.Icon = theme.CancelIcon()
	cancelButton.Resize(fyne.NewSize(100, 30))

	w.SetContent(container.NewVBox(meetingRoomButton, cubicleButton, cancelButton))
	w.Show()
}

func showMeetingForm(a fyne.App) {
	w := a.NewWindow("Meeting Form")
	w.Resize(fyne.NewSize(600, 400)) // Set the window size to 600x400

	buildingOptions := []string{"BGL1", "BGL2", "BGL3", "BGL4", "BGL5"}
	buildingSelect := widget.NewSelect(buildingOptions, func(s string) {
		fmt.Println("Building selected:", s)
	})

	floorOptions := []string{}
	for i := 1; i <= 10; i++ {
		floorOptions = append(floorOptions, fmt.Sprintf("%d", i))
	}
	floorSelect := widget.NewSelect(floorOptions, func(s string) {
		fmt.Println("Floor selected:", s)
	})

	dayOptions := []string{}
	for i := 1; i <= 31; i++ {
		dayOptions = append(dayOptions, fmt.Sprintf("%02d", i))
	}
	daySelect := widget.NewSelect(dayOptions, func(s string) {
		fmt.Println("Day selected:", s)
	})

	monthOptions := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	monthSelect := widget.NewSelect(monthOptions, func(s string) {
		fmt.Println("Month selected:", s)
	})

	yearOptions := []string{}
	for i := 2024; i <= 2025; i++ {
		yearOptions = append(yearOptions, fmt.Sprintf("%d", i))
	}
	yearSelect := widget.NewSelect(yearOptions, func(s string) {
		fmt.Println("Year selected:", s)
	})

	startTimeOptions := []string{}
	for i := 0; i < 24; i++ {
		for j := 0; j < 60; j++ {
			startTimeOptions = append(startTimeOptions, fmt.Sprintf("%02d:%02d", i, j))
		}
	}
	startTimeSelect := widget.NewSelect(startTimeOptions, func(s string) {
		fmt.Println("Start time selected:", s)
	})

	endTimeOptions := startTimeOptions
	endTimeSelect := widget.NewSelect(endTimeOptions, func(s string) {
		fmt.Println("End time selected:", s)
	})

	meetingForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Building No:", Widget: buildingSelect},
			{Text: "Floor No:", Widget: floorSelect},
			{Text: "Date:", Widget: container.NewHBox(daySelect, monthSelect, yearSelect)},
			{Text: "Start Time:", Widget: startTimeSelect},
			{Text: "End Time:", Widget: endTimeSelect},
		},
		OnSubmit: func() {
			dialog.ShowInformation("Meeting Room Reserved", "Your meeting room has been reserved successfully!", w)
		},
                OnCancel: func() {
			w.Close()
                },

	}

	w.SetContent(container.NewVBox(meetingForm))
	w.Show()
}

func showCubeForm(a fyne.App) {
        w := a.NewWindow("Cube Form")
        w.Resize(fyne.NewSize(600, 400)) // Set the window size to 600x400

        buildingOptions := []string{"BGL1", "BGL2", "BGL3", "BGL4", "BGL5"}
        buildingSelect := widget.NewSelect(buildingOptions, func(s string) {
                fmt.Println("Building selected:", s)
        })

        floorOptions := []string{}
        for i := 1; i <= 10; i++ {
                floorOptions = append(floorOptions, fmt.Sprintf("%d", i))
        }
        floorSelect := widget.NewSelect(floorOptions, func(s string) {
                fmt.Println("Floor selected:", s)
        })

        dayOptions := []string{}
        for i := 1; i <= 31; i++ {
                dayOptions = append(dayOptions, fmt.Sprintf("%02d", i))
        }
        daySelect := widget.NewSelect(dayOptions, func(s string) {
                fmt.Println("Day selected:", s)
        })

        monthOptions := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
        monthSelect := widget.NewSelect(monthOptions, func(s string) {
                fmt.Println("Month selected:", s)
        })

        yearOptions := []string{}
        for i := 2024; i <= 2025; i++ {
                yearOptions = append(yearOptions, fmt.Sprintf("%d", i))
        }
        yearSelect := widget.NewSelect(yearOptions, func(s string) {
                fmt.Println("Year selected:", s)
        })
        
        startTimeOptions := []string{}
        for i := 0; i < 24; i++ {
                for j := 0; j < 60; j++ {
                        startTimeOptions = append(startTimeOptions, fmt.Sprintf("%02d:%02d", i, j))
                }
        }
        startTimeSelect := widget.NewSelect(startTimeOptions, func(s string) {
                fmt.Println("Start time selected:", s)
        })
        
        endTimeOptions := startTimeOptions
        endTimeSelect := widget.NewSelect(endTimeOptions, func(s string) {
                fmt.Println("End time selected:", s)
        })

        cubeForm := &widget.Form{
                Items: []*widget.FormItem{
                        {Text: "Building No:", Widget: buildingSelect},
                        {Text: "Floor No:", Widget: floorSelect},
                        {Text: "Date:", Widget: container.NewHBox(daySelect, monthSelect, yearSelect)},
                        {Text: "Start Time:", Widget: startTimeSelect},
                        {Text: "End Time:", Widget: endTimeSelect},
                },
                OnSubmit: func() {
                        dialog.ShowInformation("Cubicle Reserved", "Your Cubicle has been reserved successfully!", w)
                },
                OnCancel: func() {
                        w.Close()
                },

        }

        w.SetContent(container.NewVBox(cubeForm))
        w.Show()
}