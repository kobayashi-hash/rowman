package main
import (
	flag "github.com/spf13/pflag"
	"fmt"
	"log"
	"os"
	"strings"
	"encoding/csv"
	"io"
)

type options struct{
	rowNum int
	colNum int
	filterText string
	help bool
	completion bool
}

func helpMessage(command string) string{
	return fmt.Sprintf(`%s [OPTIONS] [csv_file]
    -r, --rows <N>           output the first N rows
    -c, --cols <N>           output the first N columns
    -f, --filter <TEXT>      output only rows that contain the specified TEXT

    -h, --help               Print this message`, command)
}


func buildFlagSet() (*flag.FlagSet, *options){
	opts := &options{rowNum: 0, colNum:0, filterText:"", help:false}
	flags := flag.NewFlagSet("rowman",flag.ContinueOnError)
	flags.Usage = func(){fmt.Println(helpMessage("rowman"))}
	flags.IntVarP(&opts.rowNum,        "rows",   "r", 0,      "output the first N rows.")
	flags.IntVarP(&opts.colNum,        "cols",   "c", 0,      "output the first N columns.")
	flags.StringVarP(&opts.filterText, "filter", "f", "", "output only rows that contain the specified TEXT.")
	flags.BoolVarP(&opts.help,         "help",   "h", false,  "Print this message")
	flags.BoolVarP(&opts.completion, "generate-completions", "", false, "Generate shell completions")

	return flags, opts
}

// 他的オプション(countでオプションが指定されてる数を計算)
func validateOpts(opts *options) error{
    count := 0
    if opts.rowNum > 0 { count++ }
    if opts.colNum > 0 { count++ }
    if opts.filterText != ""{ count++ }
	if opts.help == true { count++ }
	if opts.completion == true { count++}
    if count != 1 {
		return fmt.Errorf("Please specify exactly one of --rows, --cols, or --filter.")
    }
	return nil
}

// 標準入力からcsvを読み込み
func readCSVFromStdin(r io.Reader) ([][]string, error) {
    reader := csv.NewReader(r)
    return reader.ReadAll()
}

// csvを標準出力
func writeCSVToStdout(records [][]string) error {
    writer := csv.NewWriter(os.Stdout)
    defer writer.Flush()
    return writer.WriteAll(records)
}

// 先頭からrowNum行抽出
func headRows(records [][]string, rowNum int) [][]string{
	if(rowNum > len(records)){
		return records
	}
	return records[:rowNum]
}

// 先頭からcolNum列抽出
func headCols(records [][]string, colNum int) [][]string{
	result := make([][]string, len(records))
	for i,record := range records{
		if colNum > len(record){
			result[i]=record
		}else{
			result[i]=record[:colNum]
		}
	}
	return result
}


// keywordが含まれる行をフィルタ
func filter(records [][]string, keyword string) [][]string{
	var result [][]string
	for _,record := range records{
		for _,cell := range record{
			if strings.Contains(cell, keyword){
				result = append(result, record)
				break
			}
		}
	}
	return result
}
// ファイル名の数をチェックする
func validateArgs(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("only one CSV file can be specified")
	}
	return nil
}


func main(){
	flags, opts := buildFlagSet();

	if 	err := flags.Parse(os.Args[1:]); err != nil {
        log.Fatalf("Error parsing flags: %v", err)
	}

	// 排他的オプション
	if err := validateOpts(opts); err != nil {
        log.Fatalf("Error: %v", err)
	}

	// helpの表示
	if opts.help {
		flags.Usage()
		return
	}

	// 補完作成
	if opts.completion {
		if err := generateCompletions(flags, "rowman"); err != nil {
			log.Fatalf("Failed to generate completions: %v", err)
		}
		fmt.Println("Completions generated successfully.")
		return
	}


	// 入力ファイルの数をチェック（1つまで）
	args := flags.Args()
	if err := validateArgs(args); err != nil {
        log.Fatalf("Error: %v", err)
	}

	// csv読み込み
	var input *os.File
	if len(args) == 1 { // ファイル名が指定された時はファイル読み込みの結果を渡す.
		var err error
		input, err = os.Open(flags.Args()[0])
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer input.Close()
	} else {		    // csv形式のテキストの場合はそのまま渡す.
		input = os.Stdin
	}

	records, err := readCSVFromStdin(input)
    if err != nil {
        log.Fatalf("Failed to read input: %v", err)
    }

    // 排他条件なので1つだけ判定
    if opts.filterText != ""{
        records = filter(records, opts.filterText)
    } else if opts.rowNum > 0 {
        records = headRows(records, opts.rowNum)
    } else if opts.colNum > 0 {
        records = headCols(records, opts.colNum)
    }

    if err := writeCSVToStdout(records); err != nil {
        log.Fatalf("Failed to write output: %v", err)
    }
}