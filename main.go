package main

import (
	"embed"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:embed lib/caj2pdf.exe
var qpdfData embed.FS

func main() {
	// os.Setenv("TMPDIR", "/home/temp")

	convert := flag.String("convert", "convert.docx", "convert mode")
	flag.Parse()
	// get all file name
	file := filepath.Base(*convert)
	fmt.Printf("Input file: %s\n", file)
	ext := filepath.Ext(file)
	ext = strings.TrimLeft(ext, ".")
	newpath, err := getPath(*convert, ext)
	if err != nil {
		panic(err)
	}
	fmt.Println(newpath)
	if _, err := os.Stat(newpath); os.IsNotExist(err) {
		// mkdir
		if err := os.MkdirAll(newpath, 0755); err != nil {
			fmt.Println(err)
		}
	}
	nameWithoutExt := strings.TrimSuffix(file, ext)
	fmt.Println(nameWithoutExt) // 输出：file
	new_file_name := newpath + "/" + nameWithoutExt + "pdf"
	fmt.Printf("newpath: %s\n", new_file_name)
	cmd := exec.Command("soffice", "--headless", "--convert-to", "pdf", *convert, "--outdir", newpath)
	// check if caj
	if ext == "caj" {
		// // start caj2pdf
		exeData, err := qpdfData.ReadFile("lib/caj2pdf.exe")
		if err != nil {
			panic(err)
		}

		exePath := filepath.Join(os.TempDir(), "caj2pdf.exe")
		err = ioutil.WriteFile(exePath, exeData, 0644)
		if err != nil {
			panic(err)
		}
		// caj2pdf convert test.caj -o output.pdf
		// ./caj2pdf.exe convert E:\常用工具/go/convert/1.caj -o E:\常用工具/go/convert/1.pdf
		cmd = exec.Command(exePath, "convert", *convert, "-o", new_file_name)
	} else {
		// libreoffice7.4 --headless --convert-to pdf /home/test.docx --outdir /home
		cmd = exec.Command("soffice", "--headless", "--convert-to", "pdf", *convert, "--outdir", newpath)
	}

	fmt.Println(cmd)
	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("done")
	}
	// out, err := cmd.CombinedOutput()

	// if err != nil {
	// 	fmt.Println("Error executing command:", err)
	// 	fmt.Println("Command output:", string(out))
	// } else {
	// 	fmt.Println("Command output:", string(out))
	// }

	// // Clean up temporary files
	// os.Remove(filepath.Join(os.TempDir(), "qpdf29.dll"))
	os.Remove(filepath.Join(os.TempDir(), "caj2pdf.exe"))
}
