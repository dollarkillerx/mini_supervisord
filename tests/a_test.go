package tests

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	s := "./aaaaaaa 0.0.0.0:8083"
	split := strings.Split(s, " ")
a:
	cmd := exec.Command(split[0])
	cmd.Args = split

	errPipe, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalln(err)
	}

	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}


	outReader := bufio.NewReader(outPipe)
	errReader := bufio.NewReader(errPipe)

	err = cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	go func() {
		go func() {
			for {
				fmt.Println("in")
				buf := make([]byte,1024)
				idx, err := outReader.Read(buf)
				if err != nil {
					log.Println(err)
					break
				}
				fmt.Printf("%s", buf[:idx])
			}
		}()

		go func() {
			for {
				readString, err := errReader.ReadString('\n')
				if err != nil {
					log.Println(err)
					break
				}
				fmt.Printf("%s", readString)
			}
		}()


	}()

	if err := cmd.Wait(); err != nil {
		log.Println(err)
		if !cmd.ProcessState.Success() {

			goto a
		}
	}

	fmt.Println(cmd.Args)
	time.Sleep(time.Second * 10)

	fmt.Println("in")

	for {
		select {}
	}
}

//func main() {
//	cmd := exec.Command("./aaaaaaa")
//	var stdout, stderr bytes.Buffer
//	cmd.Stdout = &stdout
//	cmd.Stderr = &stderr
//	err := cmd.Start()
//	if err != nil {
//		log.Fatalf("cmd.Run() failed with %s\n", err)
//	}
//	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
//	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
//	fmt.Println("pid: ", cmd.Process.Pid)
//
//	if err := cmd.Wait(); err != nil {
//		log.Println(err)
//	}
//
//
//	fmt.Println(cmd.Args)
//	time.Sleep(time.Second * 10)
//	outStr, errStr = string(stdout.Bytes()), string(stderr.Bytes())
//	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
//	fmt.Println(cmd.ProcessState.Success())
//	//fmt.Println("kill")
//
//	//if err := syscall.Kill(cmd.Process.Pid, 9); err != nil {
//	//	fmt.Println(err)
//	//}
//
//	fmt.Println("in")
//
//	for {
//		select {}
//		//readString, err := stdout.ReadString('\n')
//		//if err != nil {
//		//	fmt.Println(readString)
//		//}else {
//		//	fmt.Println(err)
//		//}
//	}
//}

//func main() {
//	go func() {
//		command := exec.Command("echo", "helloworld")
//
//		if err := command.Start(); err != nil {
//			log.Fatalln(err)
//		}
//	}()
//
//
//	for {
//		select {
//
//		}
//	}
//}

