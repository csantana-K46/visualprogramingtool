package scripts

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const SCRIPTROOT = "./scripts/script.py"

func ExeC() string {
	var result string
	var out bytes.Buffer

	cmd := exec.Command(SCRIPTROOT)
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
	result = out.String()
	return result
}

func ExecuteCode(code string) string {
	Write(code)
	return ExeC()
}

func EvalCode(code string) string {
	return ExecuteCode(code)
}

func AstScriptEvaluation() string {
	var code string

	return code
}

func ClearScript() {
	removeLines(SCRIPTROOT, 8, 100)
}
func Write(data string) {
	ClearScript()
	file, err := os.OpenFile(SCRIPTROOT, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString(data); err != nil {
		log.Fatal(err)
	}
}

func Delete() {
	fpath := SCRIPTROOT

	f, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var bs []byte
	buf := bytes.NewBuffer(bs)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ss := scanner.Text()
		print(ss)
		if scanner.Text() != "[manifest_version]" {
			_, err := buf.Write(scanner.Bytes())
			if err != nil {
				log.Fatal(err)
			}
			_, err = buf.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(fpath, buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func removeLines(fn string, start, n int) (err error) {
	var f *os.File
	if f, err = os.OpenFile(fn, os.O_RDWR, 0); err != nil {
		return
	}
	defer func() {
		if cErr := f.Close(); err == nil {
			err = cErr
		}
	}()
	var b []byte
	if b, err = ioutil.ReadAll(f); err != nil {
		return
	}
	cut, ok := skip(b, start-1)
	/*if !ok {
		return fmt.Errorf("less than %d lines", start)
	}*/
	if n == 0 {
		return nil
	}
	tail, ok := skip(cut, n)
	if !ok {
		//return fmt.Errorf("less than %d lines after line %d", n, start)
	}
	t := int64(len(b) - len(cut))
	if err = f.Truncate(t); err != nil {
		return
	}
	if len(tail) > 0 {
		_, err = f.WriteAt(tail, t)
	}
	return
}

func skip(b []byte, n int) ([]byte, bool) {
	for ; n > 0; n-- {
		if len(b) == 0 {
			return nil, false
		}
		x := bytes.IndexByte(b, '\n')
		if x < 0 {
			x = len(b)
		} else {
			x++
		}
		b = b[x:]
	}
	return b, true
}
