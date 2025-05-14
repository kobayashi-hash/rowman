import flag "github.com/spf13/pflag"
import ("fmt")

type options struct{
	rowNum int
	colNum int
	filterText string
	help bool
}

func buildFlagSet() (*flag.FlagSet, *options){
	opts := &options{targets: &OutputTargets{}, printer: &PrintOptions{}, rowNum: 0, colNum:0}
	flags := flag.NewFlagSet("rowman",flag.ContinueOnError)
	flag.Usage = func(){fmt.Println(helpMessage())}
	flags.IntVarP(&opts.rowNum,        "rows",   "r", 0,      "output the first N rows.")
	flags.IntVarP(&opts.colNum,        "cols",   "c", 0,      "output the first N columns.")
	flags.StringVarP(&opts.filterText, "filter", "f", "TEXT", "output only rows that contain the specified TEXT.")
	flags.BoolVarP(&opts.help,         "help",   "h", false,  "Print this message")
	return flags, opts
}