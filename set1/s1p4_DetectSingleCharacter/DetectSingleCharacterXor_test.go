package detectsinglecharacter

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/phuongdoan13/Cryptopals/set1/s1p3_SinglebyteXorCipher"
)

func readFile() (<-chan string, error){
	fptr := flag.String("fpath", "puzzle.txt", "file path to read from")
  flag.Parse()

  f, err := os.Open(*fptr)
  if err != nil {
		return nil, err
  }

	lines := make(chan string)

	go func() {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines <- scanner.Text();
		}

		close(lines)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return lines, nil
}

func TestReadFile(t *testing.T) {
	expected := "Now that the party is jumping\n"
	lines, err := readFile()
	if err != nil {
		t.Fatal(err)
	}

	for line := range lines {
		result, err := singlebyteXorCipher.SinglebyteXorCipher(line) 
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(result)
		if result == expected {
			return
		}
	}
	
	t.Fatalf("No valid result. The expected plaintext is '%s'", expected)
}