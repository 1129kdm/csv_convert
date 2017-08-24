package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {

	log.Printf("start program")
	var wg sync.WaitGroup

	for i := 1; i < len(os.Args); i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()

			// read
			fp1, err := os.Open(os.Args[num])
			failOnError(err)
			reader := csv.NewReader(fp1)
			reader.Comma = ','
			reader.LazyQuotes = true
			defer fp1.Close()

			// write
			fp2, err := os.OpenFile("converted_"+os.Args[num], os.O_WRONLY|os.O_CREATE, 0755)
			failOnError(err)
			writer := csv.NewWriter(fp2)
			defer fp2.Close()

			for {
				record, err := reader.Read()
				if err == io.EOF {
					break
				} else {
					failOnError(err)
				}

        // convert
        // left word  : before
        // right word : after
				for k := 0; k < len(record); k++ {
					r := strings.NewReplacer(
						"A", "*****",
						"B", "*****",
						"C", "*****",
					)
					resStr := r.Replace(record[k])
					record[k] = resStr
				}
				writer.Write(record)
			}
			writer.Flush()
			log.Printf("Finish convert: " + os.Args[num])
		}(i)
	}
	wg.Wait()
	log.Printf("end program")
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}
