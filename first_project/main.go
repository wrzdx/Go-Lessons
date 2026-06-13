package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todos/event"
	"todos/task"
)

func printCommands() {
	fmt.Println("- help — эта команда позволяет узнать доступные команды и их формат")
	fmt.Println("- add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("- list — эта команда позволяет получить полный список всех задач")
	fmt.Println("- del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("- done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную")
	fmt.Println("- events — эта команда позволяет получить список всех событий")
	fmt.Println("- exit — эта команда позволяет завершить выполнение программы")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := []task.Task{}
	events := []event.Event{}

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		fields := strings.Fields(input)
		if len(fields) == 0 {
			continue
		}
		fmt.Println()
		eventDesc := ""

		switch fields[0] {
		case "exit":
			if len(fields) != 1 {
				eventDesc = "Invalid number of parameters, enter `help` command to check list of available commands"
				fmt.Println(eventDesc)
				break
			}
			fmt.Println("\rBye-bye...")
			return
		case "help":
			if len(fields) != 1 {
				eventDesc = "Invalid number of parameters, enter `help` command to check list of available commands"
				fmt.Println(eventDesc)
				break
			}
			printCommands()
		case "add":
			if len(fields) != 3 {
				eventDesc = "Invalid number of parameters, enter `help` command to check list of available commands"
				fmt.Println(eventDesc)
				break
			}
			title, description := fields[1], strings.Join(fields[2:], "")
			task, err := task.CreateTask(title, description)
			if err != nil {
				eventDesc = err.Error()
				fmt.Println("Failure: ", err)
			} else {
				tasks = append(tasks, task)
				fmt.Print("Created new task: ")
				fmt.Println(task)
			}
		case "list":
			if len(fields) != 1 {
				eventDesc = "Invalid number of parameters, enter `help` command to check list of available commands"
				fmt.Println(eventDesc)
				break
			}
			fmt.Println(tasks)
		case "events":
			if len(fields) != 1 {
				eventDesc = "Invalid number of parameters, enter `help` command to check list of available commands"
				fmt.Println(eventDesc)
				break
			}
			fmt.Println(events)
		case "del":
			if len(fields) != 2 {
				eventDesc = "Invalid number of parameters, enter `help` command to check list of available commands"
				fmt.Println(eventDesc)
				break
			}
			title := fields[1]
			var deleted *task.Task
			tasks, deleted = DeleteTaskByTitle(title, tasks)
			if deleted != nil {
				fmt.Println("Deleted:", deleted)
			} else {
				fmt.Println("Not found")
			}
		case "done":
			if len(fields) != 2 {
				eventDesc = "Invalid number of parameters, enter `help` command to check list of available commands"
				fmt.Println(eventDesc)
				break
			}
			title := fields[1]
			task, ok := DoneTaskByTitle(title, tasks)
			if ok {
				fmt.Println("Task:", task)
			} else {
				fmt.Println("Not found")
			}
		default:
			eventDesc = "Unknown command, enter `help` " +
				"to see the list of available commands"
			fmt.Println(eventDesc)

		}

		fmt.Println()
		events = append(events, event.CreateEvent(input, eventDesc))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n\nBye-bye...")

}

func DeleteTaskByTitle(title string, tasks []task.Task) ([]task.Task, *task.Task) {
	for i := range tasks {
		if tasks[i].Title() == title {
			deleted := &tasks[i]
			tasks = append(tasks[:i], tasks[i+1:]...)
			return tasks, deleted
		}
	}
	return tasks, nil
}

func DoneTaskByTitle(title string, tasks []task.Task) (*task.Task, bool) {
	for i := range tasks {
		if tasks[i].Title() == title {
			tasks[i].CompleteTask()
			return &tasks[i], true
		}
	}
	return nil, false
}
