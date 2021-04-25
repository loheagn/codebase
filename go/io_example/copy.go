package io_example

import (
	"bytes"
	"io"
	"log"
)

func copyExample() {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	writer := new(bytes.Buffer)

	// io.Copy 完成了从 proverbs 读取数据并写入 file 的流程
	if _, err := io.Copy(writer, proverbs); err != nil {
		log.Fatal(err.Error())
	}

	// 这里将读到空
	bs, err := io.ReadAll(proverbs)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(string(bs))
}
