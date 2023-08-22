package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kafka_events/pkg/event"
	"kafka_events/pkg/utils"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
)

var eventTypes = []string{"view", "view", "view", "save", "save", "sample", "message"}
var trackableType = []string{"Product", "Brand"}

func getRandomAttribute(list []string) string {
	rand.Seed(time.Now().UnixNano())
	return list[rand.Intn(len(list))]
}

// getRandomId generates a random number between 1 and 100 for controlled chaos
func getRandomStringId() string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1

	return strconv.Itoa(randomNumber)
}

func generateRandomEvent() *event.Event {
	return &event.Event{
		ActorID:          uuid.New().String(),
		TrackableOwnerID: getRandomStringId(),
		EventType:        getRandomAttribute(eventTypes),
		TrackableType:    getRandomAttribute(trackableType),
		TrackableID:      getRandomStringId(),
		Metadata:         event.Metadata{"user-agent": randomdata.UserAgentString()},
	}
}

func sendRequest(event *event.Event, wg *sync.WaitGroup) {
	defer wg.Done()
	const url = "http://localhost:4000/event"
	payload, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}

	if resp.StatusCode != http.StatusCreated {
		panic(fmt.Sprintf("Request failed with status code: %d", resp.StatusCode))
	}

	fmt.Println("Request sent:", event.EventType)
}

func main() {
	const numRequests = 100000
	const maxConcurrentRequests = 1

	var wg sync.WaitGroup
	semaphore := utils.NewSemaphore(maxConcurrentRequests)

	for i := 0; i < numRequests; i++ {
		semaphore.Acquire()

		wg.Add(1)
		go func() {
			defer semaphore.Release()

			event := generateRandomEvent()
			sendRequest(event, &wg)
		}()

		wg.Wait()
	}

	fmt.Println("All requests sent.")
}
