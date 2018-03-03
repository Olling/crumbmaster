package main


import (

	"github.com/olling/slog"
)


func GetNext () {
	slog.PrintTrace("Func called: GetNext")
	queue := GetCurrentQueue()
	queue.MoveFirstToBack()
	queue.Write()
	//SendReminderBread(queue.GetResponsible())
}

func Remind () {
	slog.PrintTrace("Func called: GetReminder")
	//Send something about 
}

func BreadDate () {
	slog.PrintTrace("Func called: GetBreadDate")
	slog.PrintWarning("Remember to eat your bread!")
}

//sweaters := Inventory{"wool", 17}
//tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
//if err != nil { panic(err) }
//err = tmpl.Execute(os.Stdout, sweaters)
//if err != nil { panic(err) }


func main() {
	InitializeConfiguration()
	CurrentConfiguration.CronNext = "@every 100s"
	InitializeCron()
	NextDate()
	InitializeWebinterface()


//	slog.PrintInfo(CronNext.Entries()[0].Schedule)

}

func test () {
	//slog.PrintTrace("Hello")

}
