package archiver

import (
	"archive/zip"
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

/**
TODO:
status() - A string representing the status of the download, either Waiting, Running or Complete
progress() - A number between 0 and 1, indicating how much progress the archive job has made
run() - Starts a new archive job (if the current status is Waiting)
reset() - Cancels the current archive job, if any, and resets to the “Waiting” state
archive_file() - The path to the archive file that has been created on the server, so we can send it to the client
get() - A class method that lets us get the Archiver for the current user
**/

type Status byte

const (
	Waiting Status = iota
	Running
	Complete
)

// don't change the order
func (s Status) String() string {
	return []string{"Waiting", "Running", "Complete"}[s]
}

type Archiver struct {
	Path     string
	Status   Status
	Progress int
	Logs     []string
}

func NewArchiver() *Archiver {
	return &Archiver{
		Status: Waiting,
	}
}

func (a *Archiver) Run(ctx context.Context, rd io.Reader) {
	fmt.Println("running")
	a.Status = Running
	a.Logs = append(a.Logs, "Running archiver")

	fileName := a.RandomName(5)

	p, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	path := filepath.Join(p, "assets", "contacts", fileName)
	a.Path = path

	fz, err := os.Create(fmt.Sprintf("%s.zip", path))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fz.Close()

	w := zip.NewWriter(fz)

	f, _ := w.Create("contact.txt")

	b := bufio.NewReader(rd)
	b.WriteTo(f)

	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second)

	a.Status = Complete
	a.Logs = append(a.Logs, "Complete")
	fmt.Println("complete")

}

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func (a Archiver) RandomName(length int) string {

	b := make([]byte, length)

	for i := range b {
		b[i] = Charset[seededRand.Intn(len(Charset))]
	}

	return string(b)
}

type ID int

type Archivers struct {
	Data map[ID]Archiver
}
