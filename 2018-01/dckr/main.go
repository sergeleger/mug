package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// docker run <container> cmd args
// dckr run cmd args
func main() {
	switch os.Args[1] {
	case "run":
		run(os.Args[2:])
	case "client":
		client(os.Args[2:])
	default:
		panic("What?!")
	}
}

func run(args []string) {
	cmd := exec.Command("/proc/self/exe")
	cmd.Args = append(cmd.Args, "client")
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	os.Mkdir("/sys/fs/cgroup/pids/dckrRules", 0700)
	ioutil.WriteFile(filepath.Join("/sys/fs/cgroup/pids/dckrRules", "pids.max"), []byte("10"), 0600)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWNET | syscall.CLONE_NEWPID,
	}

	must(cmd.Run())

	os.Remove("/sys/fs/cgroup/pids/dckrRules")
}

func client(args []string) {
	ioutil.WriteFile(filepath.Join("/sys/fs/cgroup/pids/dckrRules", "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0600)

	fmt.Printf("Running %v PID: %d\n", args, os.Getpid())

	must(syscall.Chroot("/images"))
	must(syscall.Chdir("/"))

	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(args[0])
	cmd.Args = append(cmd.Args, args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
