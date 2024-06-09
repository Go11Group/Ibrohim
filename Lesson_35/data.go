package main

import (
	"encoding/json"
	"fmt"
	"gorilla_pg/model"
	"gorilla_pg/storage/postgres"
	"math/rand"
	"os"
	"time"
)

func PopulateUsers(uRepo *postgres.UserRepo) {
	f, err := os.Open("users.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	users := []model.User{}
	err = json.NewDecoder(f).Decode(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	for _, u := range users {
		err = uRepo.CreateUser(model.User{Username: u.Username, Email: u.Email, Password: u.Password})
		if err != nil {
			panic(err)
		}
	}
}

func PopulateProblems(pRepo *postgres.ProblemRepo) {
	problems := [10]model.Problem{
        {Title: "Two Sum", Description: "Given an array of integers, return indices of the two numbers such that they add up to a specific target.", Difficulty: "Easy", Acceptance: 0.45},
        {Title: "Add Two Numbers", Description: "Given two non-empty linked lists representing non-negative integers, add them and return the sum as a linked list.", Difficulty: "Medium", Acceptance: 0.35},
        {Title: "Longest Substring Without Repeating Characters", Description: "Given a string s, find the length of the longest substring without repeating characters.", Difficulty: "Medium", Acceptance: 0.30},
        {Title: "Median of Two Sorted Arrays", Description: "Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.", Difficulty: "Hard", Acceptance: 0.25},
        {Title: "ZigZag Conversion", Description: "Given a string, convert it to a zigzag pattern.", Difficulty: "Medium", Acceptance: 0.40},
        {Title: "Reverse Integer", Description: "Given a 32-bit signed integer, reverse digits of an integer.", Difficulty: "Easy", Acceptance: 0.50},
        {Title: "Palindrome Number", Description: "Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.", Difficulty: "Easy", Acceptance: 0.45},
        {Title: "Regular Expression Matching", Description: "Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.", Difficulty: "Hard", Acceptance: 0.20},
        {Title: "Container With Most Water", Description: "Given n non-negative integers representing points on a 2D plane, find the container with the most water.", Difficulty: "Medium", Acceptance: 0.35},
        {Title: "Longest Palindromic Substring", Description: "Given a string s, return the longest palindromic substring in s.", Difficulty: "Medium", Acceptance: 0.30},
    }
	for _, p := range problems {
		err := pRepo.CreateProblem(p) 
		if err != nil {
			panic(err)
		}
	}
}

func PopulateUserProblems(uRepo *postgres.UserRepo, pRepo *postgres.ProblemRepo, upRepo *postgres.UserProblemRepo) {
	users, err := uRepo.GetUser(model.User{})
	if err != nil {
		panic(err)
	}
	problems, err := pRepo.GetProblem(model.Problem{})
	if err != nil {
		panic(err)
	}
	for _,u := range users {
		numProblems := rand.Intn(len(problems)) + 1
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})

		for i := 0; i < numProblems; i++ {
			randProblem := problems[rand.Intn(len(problems))]
			err := upRepo.AddProblemToUser(u.ID, randProblem.ID, time.Now())
			if err != nil {
				fmt.Println("Error adding problem to user:", err)
				continue
			}
		}
	}
}