package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

func SingleHash(in, out chan interface{}) {
	wgGlobal := &sync.WaitGroup{}
	for i := range in {
		wgGlobal.Add(1)
		data := fmt.Sprintf("%v", i)
		md5 := DataSignerMd5(data)
		go func(data, md5 string) {
			wg := &sync.WaitGroup{}
			wg.Add(2)
			left, right := "", ""
			go func() {
				defer wg.Done()
				left = DataSignerCrc32(data)
			}()
			go func() {
				defer wg.Done()
				right = DataSignerCrc32(md5)
			}()
			wg.Wait()
			out <- left + "~" + right
			wgGlobal.Done()
		}(data, md5)
	}
	wgGlobal.Wait()
}

func MultiHash(in, out chan interface{}) {
	wgGlobal := &sync.WaitGroup{}
	for i := range in {
		wgGlobal.Add(1)
		go func(i interface{}) {
			hashTh := make([]string, 6)
			wg := &sync.WaitGroup{}
			wg.Add(6)
			for th := 0; th < 6; th++ {
				go func(it int) {
					hashTh[it] = DataSignerCrc32(fmt.Sprintf("%v%v", it, i))
					wg.Done()
				}(th)
			}
			wg.Wait()
			out <- strings.Join(hashTh, "")
			wgGlobal.Done()
		}(i)
	}
	wgGlobal.Wait()
}

func CombineResults(in, out chan interface{}) {
	var hash []string
	for i := range in {
		hash = append(hash, fmt.Sprint(i))
	}
	sort.Strings(hash)
	out <- strings.Join(hash, "_")

}

func ExecutePipeline(workers ...job) {
	wg := &sync.WaitGroup{}
	wg.Add(len(workers))
	in := make(chan interface{})
	defer close(in)
	for _, worker := range workers {
		out := make(chan interface{})
		go func(worker job, in, out chan interface{}) {
			defer close(out)
			defer wg.Done()
			worker(in, out)
		}(worker, in, out)
		in = out
	}
	wg.Wait()
}
