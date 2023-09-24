package main

import (
	// "fmt"
	// "math/rand"

	// "strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "golang.org/x/mobile/app"
)

type Todo struct {
	TodoItem string
	isTodoComplete bool
}


func main() {
	// Setup a list of Todo items
	var todos []Todo

	// Create some todo items
	q1 := Todo{
		TodoItem: "Testing",
		isTodoComplete: true,
	}
	q2 := Todo {
		TodoItem: "testing 1",
		isTodoComplete: true,
	}
	q3 := Todo {
		TodoItem: "testing 2",
		isTodoComplete: true,
	}
	q4 := Todo {
		TodoItem: "testing 3",
		isTodoComplete: true,
	}

	// Append todo items to the todos slice/list/array/whatever
	todos = append(todos, q1, q2, q3, q4)

	// Create the app
	a := app.New()

	// Create a new window
	w := a.NewWindow("Todo List")
	w.Resize(fyne.NewSize(480, 640)) // Resize window to 640x480

	// Create a label widget
	lbl_ := widget.NewLabel("Some Text")

	// Create a list widget
	todo_list := widget.NewList(
		func() int {
			// Return the length of the array
			return len(todos) 
		},
		func() fyne.CanvasObject {
			// Return a template of the label used inside the list widget
			return widget.NewLabel("")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			// Draw the list item from the array to the label in the list. 
			if todos[lii].isTodoComplete == false {
				co.(*widget.Label).SetText(todos[lii].TodoItem)
			} else {
				co.(*widget.Label).SetText(todos[lii].TodoItem + " - Complete")
			}
		})
	
	// Create an entry widget to add new items to the todo slice.
	add_todo := widget.NewEntry()
	add_todo.SetPlaceHolder("Add todo item here")
	add_todo_btn := widget.NewButton("Add", func() {
		todo := Todo{
			TodoItem: add_todo.Text,
			isTodoComplete: false,
		}
		todos = append(todos, todo)
		todo_list.Refresh()
	})

	// Set the content into a container.
	w.SetContent(container.NewGridWithRows(
		3,
		container.NewVBox(
			lbl_,
		),
		container.NewVBox(

			container.NewAdaptiveGrid(
				2,
				container.NewVBox(
					add_todo,
				),
				container.NewVBox(
					add_todo_btn,
				),
			),
		),
		todo_list,

	))

	// What it says.
	w.ShowAndRun()
}