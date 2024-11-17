package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	CmdFlags := NewCmdFlags()
	CmdFlags.Execute(&todos)
	todos.print()
	storage.Save(todos)
}
