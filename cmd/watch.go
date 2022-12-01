package cmd

import (
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/urfave/cli/v2"
)

var Watch = &cli.Command{
    Name: "watch",
    Usage: "Monitor clerk.yml and run the build command each time a change is made.",
    Flags: []cli.Flag{},
    Action: func(c *cli.Context) error {

        log.Println("Start monitoring clerk.yml...")

        filename := "./clerk.yml"
        watcher, err := fsnotify.NewWatcher()
        if err != nil {
            return err
        }
        defer watcher.Close()

        err = watcher.Add(filename)
        if err != nil {
            return err
        }

        renameCh := make(chan bool)
        removeCh := make(chan bool)
        errCh := make(chan error)

        go func() { // (3)
            for {
                select {
                case event:= <-watcher.Events:
                    switch {
                    case event.Op&fsnotify.Write == fsnotify.Write:
                        log.Println("update clerk.yml ...")
                        Build.Action(c)
                    case event.Op&fsnotify.Create == fsnotify.Create:
                        log.Println("Created file: ", event.Name)
                    case event.Op&fsnotify.Remove == fsnotify.Remove:
                        log.Println("Removed file: ", event.Name)
                        removeCh <- true
                    case event.Op&fsnotify.Rename == fsnotify.Rename:
                        log.Println("Renamed file: ", event.Name)
                        renameCh <- true
                    case event.Op&fsnotify.Chmod == fsnotify.Chmod:
                        log.Println("update clerk.yml ...")
                        Build.Action(c)
                    }
                case err:= <-watcher.Errors:
                    errCh <-err
                }
            }
        }()

        go func() {
            for {
                select {
                case <-renameCh:
                    err = waitUntilFind(filename)
                    if err != nil {
                        log.Fatalln(err)
                    }
                    err = watcher.Add(filename)
                    if err != nil {
                        log.Fatalln(err)
                    }
                case <-removeCh:
                    err = waitUntilFind(filename)
                    if err != nil {
                        log.Fatalln(err)
                    }
                    err = watcher.Add(filename)
                    if err != nil {
                        log.Fatalln(err)
                    }
                }
            }
        }()


        return <-errCh
    },
}

func waitUntilFind(filename string) error {
	for {
		time.Sleep(1 * time.Second)
		_, err := os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return err
			}
		}
		break
	}
	return nil
}
