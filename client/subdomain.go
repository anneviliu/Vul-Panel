package main

import (
	"fmt"
	"log"
	"os/exec"
)

func (s *Service) runSubdomain(urlList string) {
	// 读取文件中的txt
	cmdStr := fmt.Sprintf("python3 ./plugin/OneForAll/oneforall/oneforall.py --valid True --target %s run", urlList)
	//execShell("osascript -e 'tell application \"Terminal\" to do script \"echo hello\"'")
	err := execShell(cmdStr)
	if err != nil {
		log.Println("subdomain 脚本执行失败", err)
	}
}

func execShell(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}
	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func (s *Service) getUrls() {

}
