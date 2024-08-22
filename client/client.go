package main

import (
	"context"
	"flag"
	"fmt"
	pb "hybridseatreservation/reservation"
	"log"
	"strconv"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type connection struct {
	cc   *grpc.ClientConn
	c    *pb.HybridReservationServiceClient
	once sync.Once
}

const (
	defaultName = "world"
)

var (
	addr     = flag.String("addr", "localhost:50051", "the address to connect to")
	conn     *connection
	empid    uint64
	username string
	password string
)

func init() {
	conn = &connection{}
}

func InitConnection(addr string) {
	conn.once.Do(func() {
		cc, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect to server : %v", err)
		}
		conn.cc = cc
		c := pb.NewHybridReservationServiceClient(conn.cc)
		conn.c = &c
	})

}

func clientInstance() pb.HybridReservationServiceClient {
	return *conn.c
}

// Contact the server and print out its response.
func ConnectionClose() {
	defer conn.cc.Close()
}
func Init() {
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
			username = usernameEntry.Text
			password = passwordEntry.Text
			ctx := createUserCtx(username, password)
			res, err := (clientInstance()).Authenticate(ctx, &pb.AuthRequest{})
			if err == nil && res != nil {
				empid = res.Employeeid
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

	buildingOptions := []string{"1", "2", "3"}
	buildingSelect := widget.NewSelect(buildingOptions, func(s string) {
		fmt.Println("Building selected:", s)
	})

	floorOptions := []string{}
	for i := 1; i <= 3; i++ {
		floorOptions = append(floorOptions, fmt.Sprintf("%d", i))
	}
	floorSelect := widget.NewSelect(floorOptions, func(s string) {
		fmt.Println("Floor selected:", s)
	})
	/*
		fmt.Printf("selected = %s %s", firstSelected, secondSelected)
		var options []string
		// Logic to determine options based on selections
		if buildingSelect.Selected == "1" && floorSelect.Selected == "1" {
			options = []string{"bgl1_1_room1", "bgl1_1_room2", "bgl1_1_room3"}
		} else if buildingSelect.Selected == "1" && floorSelect.Selected == "2" {
			options = []string{"bgl1_2_room1", "bgl1_2_room2", "bgl1_2_room3"}
		} else if buildingSelect.Selected == "1" && floorSelect.Selected == "3" {
			options = []string{"bgl1_3_room1", "bgl1_3_room2", "bgl1_3_room3"}
		} else if buildingSelect.Selected == "2" && floorSelect.Selected == "1" {
			options = []string{"bgl2_1_room1", "bgl2_1_room2", "bgl2_1_room3"}
		} else if buildingSelect.Selected == "2" && floorSelect.Selected == "2" {
			options = []string{"bgl2_2_room1", "bgl2_2_room2", "bgl2_2_room3"}
		} else if buildingSelect.Selected == "2" && floorSelect.Selected == "3" {
			options = []string{"bgl2_3_room1", "bgl2_3_room2", "bgl2_3_room3"}
		} else if buildingSelect.Selected == "3" && floorSelect.Selected == "1" {
			options = []string{"bgl3_1_room1", "bgl3_1_room2", "bgl3_1_room3"}
		} else if buildingSelect.Selected == "3" && floorSelect.Selected == "2" {
			options = []string{"bgl3_2_room1", "bgl3_2_room2", "bgl3_2_room3"}
		} else if buildingSelect.Selected == "3" && floorSelect.Selected == "3" {
			options = []string{"bgl3_3_room1", "bgl3_3_room2", "bgl3_3_room3"}
		}
		meetingSelect := widget.NewSelect(options, func(s string) {
			fmt.Println("Floor selected:", s)
		})
	*/

	options := []string{}
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				options = append(options, fmt.Sprintf("bgl%d_%d_room%d", i, j, k))
			}
		}
	}
	meetingSelect := widget.NewSelect(options, func(s string) {
		fmt.Println("Floor selected:", s)
	})

	dayOptions := []string{}
	for i := 1; i <= 31; i++ {
		dayOptions = append(dayOptions, fmt.Sprintf("%02d", i))
	}
	daySelect := widget.NewSelect(dayOptions, func(s string) {
		fmt.Println("Day selected:", s)
	})

	monthOptions := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
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
		startTimeOptions = append(startTimeOptions, fmt.Sprintf("%02d:%02d:%02d", i, 0, 0))
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
			{Text: "Meeting Room:", Widget: meetingSelect},
			{Text: "Date:", Widget: container.NewHBox(daySelect, monthSelect, yearSelect)},
			{Text: "Start Time:", Widget: startTimeSelect},
			{Text: "End Time:", Widget: endTimeSelect},
		},
		OnSubmit: func() {

			fmt.Println("here1")
			bn, err := strconv.Atoi(buildingSelect.Selected)
			if err != nil {
				log.Fatal("Invalid buiding number")
			}
			fn, err := strconv.Atoi(floorSelect.Selected)
			if err != nil {
				log.Fatal("Invalid floor number")
			}
			day, err := strconv.Atoi(daySelect.Selected)
			if err != nil {
				log.Fatal("Invalid day")
			}
			month, err := strconv.Atoi(monthSelect.Selected)
			if err != nil {
				log.Fatal("Invalid month")
			}
			year, err := strconv.Atoi(yearSelect.Selected)
			if err != nil {
				log.Fatal("Invalid year")
			}
			fmt.Println("here1")
			date := fmt.Sprintf("%4d-%02d-%02d", year, month, day)
			ctx := createUserCtx(username, password)
			_, err = (clientInstance()).MeetingRoomReservation(ctx,
				&pb.MrRequest{
					BlgNumber:   uint32(bn),
					FloorNumber: uint32(fn),
					MeetingRoom: meetingSelect.Selected,
					Date:        date,
					StartTime:   startTimeSelect.Selected,
					EndTime:     endTimeSelect.Selected})

			if err == nil {
				fmt.Println("here2")
				dialog.ShowInformation("Meeting Room Reserved", "Your meeting room has been reserved successfully!", w)
			} else {
				fmt.Println("err = %v", err)

				dialog.ShowInformation("Meeting Room Reservation failed", "Please select differnt slot", w)
			}
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

	var buildingSelect *widget.Select
	var floorSelect *widget.Select
	//	var options []string
	buildingOptions := []string{"1", "2", "3"}
	buildingSelect = widget.NewSelect(buildingOptions, func(s string) {
		fmt.Println("Building selected:", s)
	})

	floorOptions := []string{}
	for i := 1; i <= 3; i++ {
		floorOptions = append(floorOptions, fmt.Sprintf("%d", i))
	}
	floorSelect = widget.NewSelect(floorOptions, func(s string) {
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
		startTimeOptions = append(startTimeOptions, fmt.Sprintf("%02d:%02d", i, 0))
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

func createUserCtx(username string, password string) context.Context {
	//Get username and password API and set it
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	md := metadata.Pairs("username", username, "password", password)
	newctx := metadata.NewOutgoingContext(ctx, md)
	return newctx
}

func main() {
	flag.Parse()
	InitConnection(*addr)
	defer ConnectionClose()
	Init()
}
