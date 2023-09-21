package backend

import (
	"encoding/json"
	"fmt"
	"os"
	"bufio"
	"io"
)

func ProcessBodyJson(body io.ReadCloser) (map[string]interface{}, error){
	decoder := json.NewDecoder(body)
	var j interface{}
	err := decoder.Decode(&j)
	mapped_j := j.(map[string]interface{})
	return mapped_j, err
}

func JsonToCss(j map[string]interface{}) error{
	fmt.Printf("PRINT STRUCT:\n %v\n", j)
	f, err := os.Create("./web/public/json.css")
    if err != nil{
		return err
	}
    defer f.Close()
	w := bufio.NewWriter(f) 

	for atom, vals := range j{
		fmt.Println(atom)
		values := vals.(map[string]interface{})
		write_err := printVarCss(w,atom, values)
		if write_err != nil{
			return write_err
		}
	}
	return nil
}

func printVarCss(w *bufio.Writer, val string, values map[string]interface{}) error{
	_, err := fmt.Fprintf(w, "\n%s {\n", val)
	if err != nil{
		return err
	}
	for property, value := range values{
		_, err_wr := fmt.Fprintf(w, "\t%s: %s;\n", property, value)
		if err_wr != nil{
			return err_wr
		}
	}
	_, err = fmt.Fprintf(w, "}\n")
	if err != nil{
		return err
	}
	w.Flush()
	return nil
}
