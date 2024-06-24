package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"eko", "kurniawan", "khanedy"})
	_ = writer.Write([]string{"eko", "kurniawan", "khanedy"})
	_ = writer.Write([]string{"eko", "kurniawan", "khanedy"})

	writer.Flush()
}
