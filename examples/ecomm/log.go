package ecomm

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex

type EventLog struct {
	EventID       string
	Event         string
	StartTime     time.Time
	EndTime       time.Time
	KafkaReceived time.Time
	Cost          uint64
	Note          string
	TimeElapsed   time.Duration
	KafkaTime     time.Duration
}

func (e EventLog) toSlice() []string {
	return []string{
		e.EventID,
		e.Event,
		e.StartTime.String(),
		e.EndTime.String(),
		e.KafkaReceived.String(),
		strconv.FormatUint(e.Cost, 10),
		e.Note,
		e.TimeElapsed.String(),
		e.KafkaTime.String(),
	}
}

func UpdateLog(filePath, event, eventID string, record_time time.Time, note string, cost uint64) (*EventLog, error) {
	mu.Lock()
	defer mu.Unlock()

	log.Print("Update Log.csv for event:", event, " with id:", eventID, " at time:", record_time)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	// Check if file is empty or newly created, if so, add the headers
	if len(records) == 0 {
		headers := []string{
			"EventID", "Event", "StartTime", "EndTime", "KafkaReceived", "Cost", "Note", "TimeElapsed", "KafkaTime",
		}
		records = append(records, headers)
	}

	var event_log EventLog
	// currentTime := time.Now()

	// Check if the event with the given eventID exists
	var existingIndex = -1
	for i, record := range records {
		if record[0] == eventID && record[1] == event {
			fmt.Println("Find record: ", record)
			existingIndex = i
			break
		}
	}

	if existingIndex == -1 {
		event_log.EventID = eventID
		event_log.Event = event
		event_log.StartTime = record_time
		event_log.Cost = cost
		event_log.Note = note

		log.Println("Save time:", event_log.StartTime)

		records = append(records, event_log.toSlice())
	} else {
		event_log = EventLog{
			EventID:       records[existingIndex][0],
			Event:         event,
			StartTime:     parseTime(records[existingIndex][2]),
			EndTime:       parseTime(records[existingIndex][3]),
			KafkaReceived: parseTime(records[existingIndex][4]),
			Cost:          parseCost(records[existingIndex][5]),
			Note:          records[existingIndex][6],
		}

		log.Println("StartTime: ", event_log.StartTime, "from: ", records[existingIndex][2])

		if event_log.EndTime.IsZero() {
			event_log.EndTime = record_time
		} else if event_log.KafkaReceived.IsZero() {
			event_log.KafkaReceived = record_time
		}

		if cost > 0 {
			event_log.Cost = cost
		}
		if note != "" {
			event_log.Note = note
		}
		if !event_log.EndTime.IsZero() && !event_log.StartTime.IsZero() {
			event_log.TimeElapsed = event_log.EndTime.Sub(event_log.StartTime)
		}
		if !event_log.EndTime.IsZero() && !event_log.KafkaReceived.IsZero() {
			event_log.KafkaTime = event_log.KafkaReceived.Sub(event_log.EndTime)
		}

		records[existingIndex] = event_log.toSlice()

	}

	file.Truncate(0)
	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(records); err != nil {
		return nil, err
	}

	return &event_log, nil
}

// func parseTime(s string) time.Time {
// 	const layout = "2006-01-02 15:04:05.99999 -0700 -07"
// 	t, _ := time.Parse(layout, s)
// 	return t
// }

func parseTime(s string) time.Time {
	const layout = "2006-01-02 15:04:05.999999 -0700 MST"

	// Remove the monotonic clock reading
	if idx := strings.Index(s, " m="); idx != -1 {
		s = s[:idx]
	}

	t, err := time.Parse(layout, s)
	if err != nil {
		fmt.Printf("Failed to parse time: %v\n", err)
		return time.Time{}
	}
	return t
}

func parseCost(s string) uint64 {
	cost, _ := strconv.ParseUint(s, 10, 64)
	return cost
}