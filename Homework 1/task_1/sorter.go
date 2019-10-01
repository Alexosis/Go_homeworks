package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

const errMsg = "Something went wrong..."

var flagF = flag.Bool("f", false, "игнорировать регистр букв")
var flagU = flag.Bool("u", false, "выводить только первое среди нескольких равных")
var flagR = flag.Bool("r", false, "сортировка по убыванию")
var flagO = flag.String("o", "", "выводить в файл, без этой опции выводить в stdout")
var flagN = flag.Bool("n", false, "сортировка чисел")
var flagK = flag.Int("k", 0, "сортировать по столбцу")

type flags struct {
	flagF bool
	flagU bool
	flagR bool
	flagO string
	flagN bool
	flagK int
}

type withoutRegister []string

func (a withoutRegister) Len() int           { return len(a) }
func (a withoutRegister) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a withoutRegister) Less(i, j int) bool { return strings.ToLower(a[i]) < strings.ToLower(a[j]) }

func main() {
	flag.Parse()
	flags := flags{
		flagF: *flagF,
		flagU: *flagU,
		flagR: *flagR,
		flagO: *flagO,
		flagN: *flagN,
		flagK: *flagK,
	}

	if arr, err := Sorter("data.txt" /*flag.Arg(0)*/, flags); err != nil || len(arr) == 0 {
		fmt.Print(err)
	} else {
		fmt.Print(arr)
	}
}

// чтение из файла
func readFromFile(pathFile string) ([]string, error) {
	data := make([]string, 0)
	lines, err := ioutil.ReadFile(pathFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data = strings.Fields(string(lines))
	return data, nil
}

//запись в файл
func writeInFile(output []string, outputFile string) {
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	file.WriteString(strings.Join(output, "\n"))
}

func Sorter(fileName string, flags flags) ([]string, error) {

	fileData, err := readFromFile(fileName)

	if err != nil {
		return []string{}, err
	}

	if flags.flagU {
		m := make(map[string]bool)

		for i := len(fileData) - 1; i >= 0; i-- {
			if m[fileData[i]] {
				fileData = append(fileData[:i], fileData[i+1:]...)
			}
			m[fileData[i]] = true
		}
		//выводить только первое среди нескольких равных
	}

	if flags.flagK > 0 {
		fileData = byRow(fileData, flags.flagK) //сортировать по столбцу
	} else if flags.flagK < 0 {
		return []string{}, errors.New(errMsg)
	}

	if flags.flagF {
		sort.Sort(withoutRegister(fileData))
	} else if flags.flagN {
		newData := make([]int, len(fileData))
		for i := range fileData {
			newData[i], _ = strconv.Atoi(fileData[i])
		}
		sort.Ints(newData) //сортировать числа
		for i := range newData {
			fileData[i] = strconv.Itoa(newData[i])
		}
	} else {
		sort.Strings(fileData)
	}

	if flags.flagR {
		for i := 0; i < len(fileData)/2; i++ {
			fileData[i], fileData[len(fileData)-i-1] = fileData[len(fileData)-i-1], fileData[i]
		} //сортировка по убыванию
	}

	if flags.flagO != "" {
		writeInFile(fileData, flags.flagO) //вывод в output.txt
	}

	return fileData, nil
}

func byRow(strings []string, num int) []string {
	newStrings := make([]string, 0)
	for key, value := range strings {
		if (key+1)%num == 0 {
			newStrings = append(newStrings, value)
		}
	}
	return newStrings
}
