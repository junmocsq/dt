package concurrent

import (
	"os/exec"
	"fmt"
	"bufio"
	"io"
	"bytes"
	"sync"
	"os"
	"syscall"
	"os/signal"
	"time"
)

func F3_2() {
	cmd := exec.Command("ls", "-la", "/Users/junmo")

	stdout0, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Errorf("Error:this command No.0 can not be startup:%s\n", err)
		return
	}
	//var outputBuf bytes.Buffer
	//
	//for{
	//	output0 := make([]byte,30)
	//	n,err := stdout0.Read(output0)
	//	if err!=nil{
	//		if err==io.EOF{
	//			break
	//		}else {
	//			fmt.Printf("error:不能从管道读取数据。%s",err)
	//			return
	//		}
	//	}
	//	outputBuf.Write(output0[:n])
	//}
	//fmt.Printf("%s\n",outputBuf.String())

	outputBuf0 := bufio.NewReader(stdout0)
	for {
		output0, _, err := outputBuf0.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Printf("%s\n", output0)

	}
}

func F3_2_1() {
	cmd1 := exec.Command("ps", "-ef")
	cmd2 := exec.Command("grep", "php")

	var outputBuf1 bytes.Buffer
	cmd1.Stdout = &outputBuf1
	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error:cmd1 不能启动；%s\n", err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Printf("error:不能等待第一个命令；%d\n", err)
		return
	}

	cmd2.Stdin = &outputBuf1
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error:第二个命令不能启动；%s\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("Error:第二个命令不能等待；%s\n", err)
		return
	}
	fmt.Printf("%s\n", outputBuf2.Bytes())
}

// 命名管道
func F3_2_2() {
	input := []byte("junmocsq")
	reader, writer, _ := os.Pipe()

	var wg sync.WaitGroup
	output := make([]byte, 100)
	wg.Add(1)
	go func() {
		n, err := reader.Read(output)
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the named pipe:%s \n", err)
		}
		fmt.Printf("Read %d bytes. [file-based pipe]\n", n)
		fmt.Println(string(output[0:n]), len(output))
		wg.Done()
	}()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			n, err := writer.Write(input)
			if err != nil {
				fmt.Printf("Error: couldn't write data to the named pipe:%s\n", err)
			}
			fmt.Printf("Write %d bytes. [file-based pipe]\n", n)
			wg.Done()
		}()
	}

	fmt.Println(output, len(output))
	wg.Wait()
}

// 信号
func F3_2_3() {
	// kill -SIGINT pid
	var wg sync.WaitGroup
	wg.Add(3)
	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("Set notification for %s ...[sigRecv1]\n", sigs1)
	signal.Notify(sigRecv1, sigs1...)
	go func() {
		for sig := range sigRecv1 {
			fmt.Printf("Received a signal from sigRecv1: %s\n", sig)
		}
		fmt.Printf("end ,sigrecv1\n")
		wg.Done()
	}()
	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("Set notification for %s ...[sigRecv2]\n", sigs2)
	signal.Notify(sigRecv2, sigs2...)
	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Received a signal from sigRecv2: %s\n", sig)
		}
		fmt.Printf("end ,sigrecv2\n")
		wg.Done()
	}()

	signal.Notify(sigRecv2)
	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Received a signal from sigRecv3: %s\n", sig)
		}
		fmt.Printf("end ,sigrecv3\n")
		wg.Done()
	}()

	fmt.Println("wait for 20 seconds...")
	time.Sleep(20 * time.Second)
	fmt.Println("stop Notification ...")
	signal.Stop(sigRecv1)
	close(sigRecv1) // 结束for循环
	signal.Stop(sigRecv2)
	close(sigRecv2)
	wg.Wait()
}
