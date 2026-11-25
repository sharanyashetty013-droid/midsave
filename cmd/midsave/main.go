package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	app := flag.String("app", "", "app to crash-test")
	check := flag.String("check", "", "program that checks the data")
	clean := flag.String("clean", "", "file to delete before each run")
	runs := flag.Int("runs", 10, "how many times to crash it")
	seed := flag.Int64("seed", 0, "random seed to replay a run (0 = pick a new one)")
	flag.Parse()
	if *app == "" || *check == "" {
		fmt.Println("usage: midsave --app <program> --check <checker> [--clean <file>] [--runs N] [--seed N]")
		os.Exit(2)
	}

	if *seed == 0 {
		*seed = time.Now().UnixNano()
	}
	rng := rand.New(rand.NewSource(*seed))
	fmt.Println("seed:", *seed)

	failed := 0
	for i := 1; i <= *runs; i++ {
		os.Remove(*clean) // fresh start each run

		cmd := exec.Command(*app)
		if err := cmd.Start(); err != nil { // 1. start the app
			fmt.Println("could not start app:", err)
			os.Exit(1)
		}

		wait := time.Duration(200+rng.Intn(2000)) * time.Millisecond
		time.Sleep(wait) // 2. wait a random time

		cmd.Process.Kill() // 3. kill it
		cmd.Wait()

		out, err := exec.Command(*check).Output() // 4. run the checker, keep what it prints
		msg := strings.TrimSpace(string(out))

		if err == nil {
			fmt.Printf("run %d: ok\n", i)
		} else if _, ok := err.(*exec.ExitError); ok {
			failed++
			fmt.Printf("run %d: BROKEN (killed after %v): %s\n", i, wait, msg)
		} else {
			fmt.Println("could not run checker:", err)
			os.Exit(1)
		}
	}
	fmt.Printf("\n%d of %d runs lost or broke data\n", failed, *runs)
	if failed > 0 {
		fmt.Printf("replay these exact crash times with: --seed %d\n", *seed)
	}
}
