package main

import "fmt"

/*
	Asumption: All jobs arrive at the same time
*/

const Time_Slice = 20.37

func main() {
	jobs1 := []float64{200.0, 200.0, 200.0}
	jobs2 := []float64{300.0, 100.0, 200.0}
	fifoTatJobs1, fifoRptJobs1 := fifo(jobs1)
	sjfTatJobs1, sjfRptJobs1 := sjf(jobs1)
	rrTatJobs1, rrRptJobs1 := rr(jobs1, Time_Slice)
	fmt.Println("Jobs1: Avg Response Time ")
	fmt.Printf("Fifo: %v\n", fifoRptJobs1)
	fmt.Printf("SJF: %v\n", sjfRptJobs1)
	fmt.Printf("Round Robin: %v\n", rrRptJobs1)
	fmt.Println()
	fmt.Println("Jobs1: Avg Turn Around Time")
	fmt.Printf("Fifo: %v\n", fifoTatJobs1)
	fmt.Printf("SJF: %v\n", sjfTatJobs1)
	fmt.Printf("Round Robin: %v\n", rrTatJobs1)
	fmt.Println()
	fifoTatJobs2, fifoRptJobs2 := fifo(jobs2)
	sjfTatJobs2, sjfRptJobs2 := sjf(jobs2)
	rrTatJobs2, rrRptJobs2 := rr(jobs2, Time_Slice)
	fmt.Println("Jobs2: Avg Response Time")
	fmt.Printf("Fifo: %v\n", fifoRptJobs2)
	fmt.Printf("SJF: %v\n", sjfRptJobs2)
	fmt.Printf("Round Robin: %v\n", rrRptJobs2)
	fmt.Println()
	fmt.Println("Jobs2: Avg Turn Around Time")
	fmt.Printf("Fifo: %v\n", fifoTatJobs2)
	fmt.Printf("SJF: %v\n", sjfTatJobs2)
	fmt.Printf("Round Robin: %v\n", rrTatJobs2)
}

func fifo(jobs []float64) (tat float64, rpt float64) {
	tat = float64(0)
	rpt = float64(0)
	curt := float64(0)
	for _, jt := range jobs {
		tat += (tat + jt)
		rpt += curt
		curt += jt
	}
	tat /= float64(len(jobs))
	rpt /= float64(len(jobs))
	return tat, rpt
}

func sjf(jobs []float64) (tat float64, rpt float64) {
	tat = float64(0)
	rpt = float64(0)
	curt := float64(0)
	jobs = mergeSort(jobs)
	for _, jt := range jobs {
		tat += (tat + jt)
		rpt += curt
		curt += jt
	}
	tat /= float64(len(jobs))
	rpt /= float64(len(jobs))
	return tat, rpt
}

func rr(jobs []float64, ts float64) (tat float64, rpt float64) {
	jlist := make([]float64, len(jobs), len(jobs))
	copy(jlist, jobs)
	tat = float64(0)
	curt := float64(0)
	cOfFinishJobs := 0
	for cOfFinishJobs != len(jlist) {
		for i, _ := range jlist {
			if jlist[i] <= 0.0 {
				tat += curt
				cOfFinishJobs += 1
				continue
			}
			jlist[i] -= ts
			curt += ts
		}
	}
	rpt = (float64(len(jobs)-1) * ts) / float64(len(jobs))
	tat /= float64(len(jobs))
	return tat, rpt
}

func mergeSort(n []float64) []float64 {
	if len(n) < 2 {
		return n
	}
	lp := mergeSort(n[:len(n)/2])
	rp := mergeSort(n[len(n)/2:])
	res := make([]float64, 0, len(n))
	i := 0
	j := 0
	for {
		if i == len(lp) {
			res = append(res, rp[j:]...)
			break
		} else if j == len(rp) {
			res = append(res, lp[i:]...)
			break
		}
		if lp[i] <= rp[j] {
			res = append(res, lp[i])
			i++
		} else {
			res = append(res, rp[j])
			j++
		}
	}
	return res
}
