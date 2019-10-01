package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

const (
	multiHashes = 6
)

func ExecutePipeline(jobs ...job) {
	current := make(chan interface{})

	for _, j := range jobs {
		out := make(chan interface{})

		go func(j job, in, out chan interface{}) {
			j(in, out)
			close(out)
		}(j, current, out)
		current = out
	}
	<-current
}

var mu sync.Mutex

func calcMd5(data string) string {
	mu.Lock()
	defer mu.Unlock()
	return DataSignerMd5(data)
}

func SingleHash(in, out chan interface{}) {
	var wg = &sync.WaitGroup{}

	for data := range in {
		s := strconv.Itoa(data.(int))
		wg.Add(1)

		go func() {
			defer wg.Done()
			crc32d := make(chan string)
			crc32Md5 := make(chan string)

			go func() {
				crc32d <- DataSignerCrc32(s)
			}()

			go func() {
				crc32Md5 <- DataSignerCrc32(calcMd5(s))
			}()
			out <- <-crc32d + "~" + <-crc32Md5
		}()
	}
	wg.Wait()
}

func MultiHash(in, out chan interface{}) {
	var wg = &sync.WaitGroup{}

	for chanData := range in {
		data := chanData.(string)
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			var waitGroup = &sync.WaitGroup{}
			hashes := make([]string, multiHashes)

			for i := 0; i < multiHashes; i++ {
				waitGroup.Add(1)

				go func(crc int, wg *sync.WaitGroup) {
					hashSrc := strconv.Itoa(crc) + data
					hashes[crc] = DataSignerCrc32(hashSrc)
					wg.Done()
				}(i, waitGroup)
			}
			waitGroup.Wait()
			out <- strings.Join(hashes, "")
			wg.Done()
		}(wg)
	}

	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	results := make([]string, 0)

	for data := range in {
		results = append(results, data.(string))
	}

	sort.Strings(results)
	result := strings.Join(results, "_")
	out <- result
}
