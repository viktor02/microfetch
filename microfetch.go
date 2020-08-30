package microfetch

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"os"
	"time"
)

func main() {
	v, _ := mem.VirtualMemory()
	h, _ := host.Info()
	p, _ := cpu.Info()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"MicroFetch", time.Now().Format(time.RFC822)})
	t.AppendRows([]table.Row{
		{"Host", h.Hostname},
		{"Platform", h.Platform},
		{"Version", h.PlatformVersion},
		{"KernelArch", h.KernelArch},
		{"Uptime", h.Uptime},
	})

	t.AppendRows([]table.Row{
		{""},
		{"CPU"},
		{"ModelName:", p[0].ModelName},
		{"Cores:", p[0].Cores},
		{"Mhz:", p[0].Mhz},
	})

	t.AppendRows([]table.Row{
		{""},
		{"MEMORY"},
		{"Total:", v.Total},
		{"Free:", v.Free},
		{"Used:", fmt.Sprintf("%.0f%%", v.UsedPercent)},
	})
	t.Render()

}
