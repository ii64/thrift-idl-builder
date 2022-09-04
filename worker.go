// SPDX-License-Identifier: MIT
package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

type metaCU struct {
	Err  error
	Path string
	Done bool
}
type mon struct {
	max   int
	finsh int
	m     []*metaCU
	mu    *sync.Mutex
}

func (m *mon) show(wg *sync.WaitGroup) {
	defer wg.Done()

	buf := bytes.NewBuffer([]byte{})
	i := 0
	max := 0
	pb := [...]rune{'|', '/', '-', '\\'}
	_ = pb
	var rEdge bool
	for {
		m.mu.Lock()

		buf.Reset()
		buf.WriteRune('\r')

		fmt.Fprintf(buf, "[%d/%d] ", m.finsh, m.max)

		if rEdge {
			buf.WriteString(fmt.Sprintf("All processed."))
		} else if len(m.m) > 0 {
		sw_start:
			l := len(m.m)
			if l <= 0 {
				goto sw_out
			}
			idx := (i / (2000 / l)) % l
			meta := m.m[idx]
			if !meta.Done {
				fmt.Fprintf(buf, "working: %s ...", meta.Path)
			} else {
				m.m = append(m.m[:idx], m.m[idx+1:]...)
				goto sw_start
			}
		} else {
			buf.WriteString("waiting...")
		}
	sw_out:
		m.mu.Unlock()
		if !rEdge { // active indicator
			buf.WriteRune(' ')
			buf.WriteRune(pb[(i/100)%len(pb)])
		}
		// pad to clear previous written strings
		if l := buf.Len(); l > max {
			max = l
		} else {
			rem := max - l
			sp := make([]byte, rem)
			for i := 0; i < rem; i++ {
				sp[i] = ' '
			}
			buf.Write(sp)
		}
		// buf.WriteRune('\n')
		buf.WriteTo(os.Stdout)
		i++
		time.Sleep(time.Millisecond)
		// edge case.
		if m.finsh >= m.max {
			if !rEdge {
				rEdge = true
			} else {
				fmt.Println()
				break
			}
		}
	}
}

func worker(id int, qpath <-chan string, m *mon, start, done func(meta *metaCU)) {
	var err error
	for {
		var meta metaCU
		meta.Path = <-qpath
		start(&meta)

		cmd := buildThriftCmd(meta.Path, thriftAddtionalFlags.value)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err = cmd.Run(); err != nil {
			meta.Err = err
			done(&meta)
			continue
		}
		done(&meta)
	}
}
