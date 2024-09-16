package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	length float64
	ticket int
}

func main() {
	j1 := &Job{length: 100.0, ticket: 100}
	j2 := &Job{length: 100.0, ticket: 100}
	j3 := &Job{length: 100.0, ticket: 100}
	jobs := map[int]*Job{1: j1, 2: j2, 3: j3}
	fairness, tat, rpt := lottory(jobs, 300)
	fmt.Println("Turn around time: ", tat)
	fmt.Println("Response time: ", rpt)
	fmt.Println("Fairness: ", fairness)
}

func lottory(jobs map[int]*Job, ttTickets int) (fairness float64, tat float64, rpt float64) {
	tat = 0.0
	rpt = 0.0
	fairness = 0.0
	ctime := 0.0
	qt := 2.0
	jobList := make([]int, 0, ttTickets)
	seenJobs := make(map[int]bool)
	for id, job := range jobs {
		tickets := job.ticket
		for tickets > 0 {
			jobList = append(jobList, id)
			tickets--
		}
	}
	doneJobsLen := 0
	fjTat := 0.0
	for len(jobs) > 0 {
		randomIndex := rand.Intn(ttTickets)
		jobId := jobList[randomIndex]
		jobs[jobId].length -= qt
		ctime += qt
		if jobs[jobId].length <= 0 {
			tat += ctime
			doneJobsLen += 1
			ttTickets -= jobs[jobId].ticket
			delete(jobs, jobId)
			jobList = removeJobFromJobsList(jobList, jobId)
			if doneJobsLen == 1 {
				fjTat = tat
			} else if doneJobsLen == 2 {
				fairness = fjTat / tat
			}
		}
		_, exist := seenJobs[jobId]
		if !exist {
			rpt += ctime
			seenJobs[jobId] = true
		}
	}
	tat /= float64(len(seenJobs))
	rpt /= float64(len(seenJobs))
	return fairness, tat, rpt
}

func removeJobFromJobsList(jobsList []int, target int) []int {
	start := -1
	end := -1
	for i, v := range jobsList {
		if start == -1 && v == target {
			start = i
		} else if start != -1 && end == -1 && v != target {
			end = i - 1
		}
	}
	if end == -1 {
		end = len(jobsList) - 1
	}
	return append(jobsList[:start], jobsList[end+1:]...)
}
