package scanner

import (
	"fmt"
	"todolist/todo"

	"github.com/k0kubun/pp"
)

func printPrompt() {
	fmt.Print("Введите команду: ")
}

func printExit() {
	fmt.Println("Завершение работы... До скорого!")
}

func printAdd(title string) {
	fmt.Println("Задача '" + title + "' успешно добавлена")
	fmt.Println()
}

func printTasks(tasks map[string]todo.Task) {
	pp.Println("Список дел:", tasks)
	fmt.Println()
}

func printDone(title string) {
	fmt.Println("Задача '" + title + "' помечена как выполненная")
	fmt.Println()
}

func printDel(title string) {
	fmt.Println("Задача '" + title + "' успешно удалена")
	fmt.Println()
}

func printHelp() {
	fmt.Println("Список команд:")
	fmt.Println("- help — эта команда позволяет узнать доступные команды и их формат")
	fmt.Println("- add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("- list — эта команда позволяет получить полный список всех задач")
	fmt.Println("- del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("- done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную")
	fmt.Println("- events — эта команда позволяет получить список всех событий")
	fmt.Println("- exit — эта команда позволяет завершить выполнение программы")
	fmt.Println()
}

func printEvents(events []Event) {
	pp.Println("События:", events)
	fmt.Println()
}

func printResult(result string) {
	fmt.Println("Результат выполнения команды:", result)
	fmt.Println()
}
