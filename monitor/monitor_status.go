package monitor

import (
	"os"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
)

var monitorStatusWg sync.RWMutex
var monitorStatus [][]string


func StartMonitorStatusDisplay(){
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Subject", "Class Number", "CRN", "Status"})
	
	for{
		table.ClearRows()
		monitorStatusWg.RLock()
		for _, v := range monitorStatus {
			table.Append(v)
		}
		monitorStatusWg.RUnlock()

		table.Render()
		time.Sleep(1 * time.Second)
	}
}

func (monitorTask *MonitorTask) updateStatus(msg string){
	monitorStatusWg.Lock()
	monitorStatus[monitorTask.MSIndex] = []string{monitorTask.Subject, monitorTask.ClassNumber, monitorTask.CRN, msg}
	monitorStatusWg.Unlock()
}