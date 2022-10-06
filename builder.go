// SPDX-License-Identifier: MIT
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var (
	appname = "thrift-idl-builder"
	cmd     = flag.NewFlagSet(appname, flag.ExitOnError)

	wg        sync.WaitGroup
	verbosity = 0

	recurseFailFirst     = false
	thriftAddtionalFlags = stringSlice{value: []string{}}
	thriftBinaryLocation = "thrift-0.13.0"
	thriftgoPlugins      = stringLiteralSlice{}
	outGenDir            = "./dist"
	dirPath              = "./idl/thrift"

	genLang = "go"

	defaultLookup     = []string{"service.thrift", "svc.thrift"}
	lookupForFilename = stringSlice{value: defaultLookup}

	showErrors  bool
	workerCount = 1

	maxParalelTaskChanBufSz = 512
)

type stringLiteralSlice []string

func (s *stringLiteralSlice) Set(p string) error {
	*s = append(*s, p)
	return nil
}

func (s *stringLiteralSlice) String() string {
	return ""
}

type stringSlice struct {
	value        []string
	hasCalledSet bool
}

func (s *stringSlice) Set(p string) error {
	// persist default value until `flag`.Parse call `Set`.
	if !s.hasCalledSet {
		s.value = make([]string, 0, 8)
		s.hasCalledSet = true
	}
	parts := SplitAtCommas(p)
	for _, part := range parts {
		part = strings.Trim(part, " ")
		str, err := strconv.Unquote(part)
		if err == nil {
			part = str
		}
		s.value = append(s.value, part)
	}
	return nil
}

func (s *stringSlice) String() string {
	return strings.Join(s.value, ",")
}

type MatcherFunc func(filename string) bool

func init() {
	setIfOSEnv(&thriftBinaryLocation, "THRIFT_PATH")
	setIfOSEnv(&genLang, "GEN_LANG")

	cmd.StringVar(&thriftBinaryLocation, "bin", thriftBinaryLocation, "thrift binary location")
	cmd.StringVar(&genLang, "gen", genLang, "thrft generate language")
	cmd.StringVar(&outGenDir, "o", outGenDir, "output location for gen-* directory")
	cmd.StringVar(&dirPath, "source-dir", dirPath, "thrift file source directory")
	cmd.BoolVar(&showErrors, "errors", false, "show errors at end of process")
	cmd.IntVar(&workerCount, "wrk", workerCount, "worker count")
	cmd.IntVar(&verbosity, "v", verbosity, "verbosity level")

	cmd.Var(&lookupForFilename, "lookup", "lookup files")
	cmd.BoolVar(&recurseFailFirst, "recfail", recurseFailFirst, "recursive fail first")

	cmd.Var(&thriftAddtionalFlags, "t", "thrift addtitional flags")
	cmd.Var(&thriftgoPlugins, "tp", "thrift kitex plugin flags")
}

func matcher(filename string) bool {
	//return g.Match(filename)
	filename = strings.ToLower(filename)
	for _, v := range lookupForFilename.value {
		if strings.HasSuffix(filename, v) {
			return true
		}
	}
	return false
}

// findFilesRecurse walk folder
func findFilesRecurse(dir string, matcher MatcherFunc, ignoreError bool) (targets []string, err error) {
	d, err := os.Open(dir)
	if err != nil {
		fmt.Printf("err: failed read: %s\n", err)
		return
	}
	defer d.Close()
	var names []fs.FileInfo
	names, err = d.Readdir(-1)
	if err != nil {
		fmt.Printf("err: failed read dir: %s\n", err)
		return
	}
	for _, item := range names {
		switch {
		case item.IsDir():
			var targetsOnSub []string
			targetsOnSub, err = findFilesRecurse(filepath.Join(dir, item.Name()), matcher, ignoreError)
			if err != nil {
				if ignoreError { // continue..
					goto Cont
				}
				return
			}
			targets = append(targets,
				targetsOnSub...)
		Cont:
			continue
		case matcher(item.Name()):
			path := filepath.Join(dir, item.Name())
			if verbosity > 0 {
				fmt.Printf("found: %s\n", path)
			}
			targets = append(targets, path)
		}
	}
	return
}

func buildThriftCmd(path string, addtFlags []string) *exec.Cmd {
	// thrift flag to specify direct directory output path without `gen-go` folder.
	outFlag := "-out"
	// thriftgo flag compatibility
	if strings.HasSuffix(thriftBinaryLocation, "thriftgo") {
		outFlag = "-o"
	}
	var args = []string{
		outFlag, outGenDir, // output
		"-r",             // recurse
		"--gen", genLang, // generate params
	}
	// append thriftgo plugin flags
	if len(thriftgoPlugins) > 0 {
		for _, val := range thriftgoPlugins {
			args = append(args, "-p", val)
		}
	}
	// append additional compile flags
	args = append(args, addtFlags...)
	// append the source path
	args = append(args, path)
	if verbosity > 1 {
		fmt.Printf("> %s\n", append([]string{thriftBinaryLocation}, args...))
	}
	return exec.Command(thriftBinaryLocation, args...)
}

func setIfOSEnv(dst *string, envKey string) {
	if dst == nil {
		return
	}
	envValue := os.Getenv(envKey)
	if envValue != "" {
		*dst = envValue
	}
}

func postArgParse() error {
	if len(lookupForFilename.value) <= 0 {
		lookupForFilename.value = defaultLookup
	}
	if len(dirPath) == 0 {
		return fmt.Errorf("dir path source cannot be empty")
	}
	fmt.Println("used plugins:", thriftgoPlugins)
	return nil
}

func Main(args []string) (err error) {
	if err = cmd.Parse(args); err != nil {
		return
	}
	if err = postArgParse(); err != nil {
		return
	}

	//	g = glob.MustCompile("*service.thrift")
	fmt.Printf("Output language: %q\n", genLang)
	fmt.Printf("Dir source path: %q\n", dirPath)
	fmt.Printf("Dir output path: %q\n", outGenDir)

	os.MkdirAll(outGenDir, 0644)
	_, err = os.Stat(outGenDir)
	if os.IsNotExist(err) {
		fmt.Printf("err: dir out is not present. create it first.")
		return
	} else if err != nil {
		fmt.Printf("err: dir out error: %s\n", err)
		return
	}

	fmt.Println("Looking for [*service.thrift]...")
	var candidates []string
	candidates, err = findFilesRecurse(dirPath, matcher, recurseFailFirst)
	if err != nil {
		return
	}

	fmt.Printf("total candidate: %d\n", len(candidates))
	if len(candidates) == 0 {
		fmt.Println("nothing to do.")
		return
	}
	fmt.Println("running build....")

	res := []*metaCU{}
	if workerCount > 1 {
		var mu sync.Mutex
		var wg sync.WaitGroup
		m := &mon{
			max: len(candidates),
			mu:  &mu,
		}
		start := func(meta *metaCU) {
			mu.Lock()
			defer mu.Unlock()
			m.m = append(m.m, meta)
		}
		done := func(meta *metaCU) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			m.finsh++
			meta.Done = true
			if meta != nil && meta.Err != nil {
				res = append(res, meta)
			}
		}
		bufSz := (((len(candidates) - 1) >> 4) + 1) << 4 // aligned to 16
		if bufSz > maxParalelTaskChanBufSz {
			bufSz = maxParalelTaskChanBufSz
		}
		qpath := make(chan string, bufSz)
		for i := 0; i < workerCount; i++ {
			go worker(i, qpath, m, start, done)
		}
		wg.Add(len(candidates))
		wg.Add(1) // "status" G
		go m.show(&wg)
		for _, path := range candidates {
			qpath <- path
		}
		// wait task done.
		wg.Wait()
	} else {
		for _, path := range candidates {
			cmd := buildThriftCmd(path, thriftAddtionalFlags.value)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Printf("working: %s .... failed :(\n", path)
				res = append(res, &metaCU{err, path, true})
				continue
			}
			fmt.Printf("working: %s .... success\n", path)
		}
	}

	var addt string
	addt = fmt.Sprintf(" %d total", len(candidates))
	if len(res) > 0 {
		if addt != "" {
			addt = "," + addt
		}
		addt = fmt.Sprintf(" %d error(s)%s", len(res), addt)
	}
	fmt.Printf("task done.%s\n", addt)
	if showErrors && len(res) > 0 {
		fmt.Println("errors:")
		for i, err := range res {
			fmt.Printf("[%d] file: %s\n", i, err.Path)
			fmt.Printf("[%d] err:  %s\n", i, err.Err)
		}
	}

	return
}
