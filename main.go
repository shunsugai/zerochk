package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "zerochk"
	app.Usage = "Checks given binary file is 0-filled or not"
	app.Version = "0.0.1"
	app.Author = "Shun Sugai"
	app.Email = "sugaishun@gmail.com"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) < 1 {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
		for _, path := range []string(c.Args()) {
			if err := checkZero(path); err != nil {
				fmt.Println(err)
			}
		}
		os.Exit(0)
	}
	app.Run(os.Args)
}

func checkZero(path string) (err error) {
	fmt.Println(path)
	f, err := os.Open(path)
	if err != nil {
		return
	}

	var count uint64

	r := bufio.NewReader(f)
	for {
		c, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		if int(c) != 0x00 {
			fmt.Printf("NOT 0-filled: found 0x%X at 0x%08X\n", c, count)
			return nil
		}
		count++
	}
	fmt.Printf("0-filled: read %d bytes", count)
	return nil
}
