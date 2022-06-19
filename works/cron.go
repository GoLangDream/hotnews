package works

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/robfig/cron/v3"
	"os"
)

var cronTask *cron.Cron = cron.New()

func Start() {
	cronTask.Start()
}

func printCronTask(name string, taskID cron.EntryID) {
	task := cronTask.Entry(taskID)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.Style().Format.Header = text.FormatTitle

	t.AppendHeader(table.Row{"#", "Task Name", "下次执行时间"})

	t.AppendRow([]interface{}{
		1,
		name,
		task.Next,
	})
	t.Render()
}
