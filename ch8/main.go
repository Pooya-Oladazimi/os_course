package main

import "fmt"

func main() {
	jobs := []Job{{id: 1, length: 400.0}, {id: 2, length: 200.0}, {id: 3, length: 1000.0}}
	tat, rsp := mlfq(jobs)
	fmt.Println()
	fmt.Println("Average Turn around time:", tat)
	fmt.Println("Average Response time:", rsp)
}

type Job struct {
	id     int
	length float64
}

func mlfq(jobs []Job) (tat float64, rsp float64) {
	tat = 0.0
	rsp = 0.0
	curt := 0
	q1 := make([]Job, 0, 3)
	q2 := make([]Job, 0, 3)
	history := make(map[int]float64)
	ts1 := 5.0
	ts2 := 10.0
	resetS := 100
	allotT := 30.0
	q1 = append(q1, jobs...)
	for _, j := range jobs {
		history[j.id] = 0.0
	}

	for {
		if len(q1) == 0 && len(q2) == 0 {
			break
		}
		i := 0
		fmt.Println("q1:", q1)
		for len(q1) != 0 {
			if q1[i].length <= 0.0 {
				q1 = append(q1[:i], q1[i+1:]...)
				tat += float64(curt)
				fmt.Println("q1:", q1)
				continue
			}
			q1[i].length -= ts1
			curt += int(ts1)
			history[q1[i].id] += ts1
			if history[q1[i].id] >= allotT {
				q2 = append(q2, q1[i])
				q1 = append(q1[:i], q1[i+1:]...)
				fmt.Println("q1:", q1)
				continue
			}
			i += 1
			if i >= len(q1) {
				i = 0
			}
		}
		i = 0
		fmt.Println("q2:", q2)
		for len(q2) != 0 {
			if q2[i].length <= 0.0 {
				q2 = append(q2[:i], q2[i+1:]...)
				tat += float64(curt)
				fmt.Println("q2:", q2)
				continue
			}
			q2[i].length -= ts2
			curt += int(ts2)
			if curt%resetS == 0 {
				q1 = append(q1, q2...)
				q2 = q2[:0]
				for k, _ := range history {
					history[k] = 0.0
				}
				break
			}
			i += 1
			if i >= len(q2) {
				i = 0
			}
		}
	}
	return tat / float64(len(jobs)), (float64(len(jobs)-1) * ts1) / float64(len(jobs))
}
