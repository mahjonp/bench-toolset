package cmd

import (
	"fmt"

	"github.com/5kbpers/bench-toolset/bench"
	"github.com/5kbpers/bench-toolset/workload"
	"github.com/spf13/cobra"
)

var (
	benchmark      string
	promethuesAddr string
)

func init() {
	analyzeCmd := NewAnalyzeCommand()

	analyzeCmd.PersistentFlags().StringVar(&logPath, "log", "", "log path of benchmark")
	analyzeCmd.PersistentFlags().StringVar(&benchmark, "benchmark", "", "benchmark name (tpcc, ycsb, sysbench)")
	analyzeCmd.PersistentFlags().IntVar(&intervalSecs, "interval", -1, "interval of metrics in seconds")

	rootCmd.AddCommand(analyzeCmd)
}

func NewAnalyzeCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "analyze",
		Short: "Analyze logs(support go-tpc, go-ycsb, sysbench) or metrics",
	}
	command.AddCommand(newLogCommand())
	return command
}

func newLogCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "log",
		Short: "Analyze benchmark(tpcc, ycsb, sysbench) logs",
		RunE: func(cmd *cobra.Command, args []string) error {
			switch benchmark {
			case "tpcc":
				records, err := workload.ParseTpccRecords(logPath)
				if err != nil {
					return err
				}
				results := bench.GetTpccJitter(records, intervalSecs)
				fmt.Println(results)
			default:
				panic("Unsupported benchmark name")
			}
			return nil
		},
	}
	return command
}
