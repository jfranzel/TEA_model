package main

import "fmt"

import (
    "fmt"
    "strconv"

    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    win := myApp.NewWindow("Name and Age Input")

    nameEntry := widget.NewEntry()
    nameEntry.SetPlaceHolder("Enter name")

    ageEntry := widget.NewEntry()
    ageEntry.SetPlaceHolder("Enter age")

    submitButton := widget.NewButton("Submit", func() {
        name := nameEntry.Text
        ageStr := ageEntry.Text
        age, err := strconv.Atoi(ageStr)
        if err != nil {
            // Handle error when age is not a valid integer
            fmt.Println("Invalid age:", err)
            return
        }

        // Do something with name and age
        fmt.Printf("Name: %s, Age: %d\n", name, age)

        // Clear entries after submission
        nameEntry.SetText("")
        ageEntry.SetText("")
    })

    inputForm := container.NewVBox(
        widget.NewLabel("Name:"),
        nameEntry,
        widget.NewLabel("Age:"),
        ageEntry,
        submitButton,
    )

    win.SetContent(inputForm)
    win.Resize(fyne.NewSize(300, 200))
    win.ShowAndRun()
}
