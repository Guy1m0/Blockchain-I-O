package ecomm

import (
	"encoding/csv"
	"os"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex

type EventLog struct {
	Event         string
	EventID       int
	StartTime     time.Time
	EndTime       time.Time
	KafkaReceived time.Time
	Cost          float64
	Note          string
	TimeElapsed   time.Duration
	KafkaTime     time.Duration
}

func (e EventLog) toSlice() []string {
	return []string{
		e.Event,
		strconv.Itoa(e.EventID),
		e.StartTime.String(),
		e.EndTime.String(),
		e.KafkaReceived.String(),
		strconv.FormatFloat(e.Cost, 'f', 2, 64),
		e.Note,
		e.TimeElapsed.String(),
		e.KafkaTime.String(),
	}
}

func UpdateLog(filePath, event string, eventID int, note string, cost float64) error {
	mu.Lock()
	defer mu.Unlock()

	// Open the CSV file
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Prepare the log entry and update/append as required
	var log EventLog
	currentTime := time.Now()
	if eventID == 0 { // New Event
		log.Event = event
		log.StartTime = currentTime
		log.Cost = cost
		log.Note = note

		// Find the maximum existing event ID
		maxID := 0
		for _, row := range records {
			id, err := strconv.Atoi(row[1])
			if err == nil && id > maxID {
				maxID = id
			}
		}
		log.EventID = maxID + 1
		records = append(records, log.toSlice())
	} else { // Updating an existing event
		for index, row := range records {
			if row[1] == strconv.Itoa(eventID) {
				log = EventLog{
					Event:         row[0],
					EventID:       eventID,
					StartTime:     parseTime(row[2]),
					EndTime:       parseTime(row[3]),
					KafkaReceived: parseTime(row[4]),
					Cost:          parseCost(row[5]),
					Note:          row[6],
				}

				if log.EndTime.IsZero() {
					log.EndTime = currentTime
				} else if log.KafkaReceived.IsZero() {
					log.KafkaReceived = currentTime
				}

				if note != "" {
					log.Note = note
				}
				if cost != 0 {
					log.Cost = cost
				}
				if !log.StartTime.IsZero() && !log.EndTime.IsZero() {
					log.TimeElapsed = log.EndTime.Sub(log.StartTime)
				}
				if !log.EndTime.IsZero() && !log.KafkaReceived.IsZero() {
					log.KafkaTime = log.KafkaReceived.Sub(log.EndTime)
				}
				records[index] = log.toSlice()
				break
			}
		}
	}

	// Clear the file before writing
	file.Truncate(0)
	file.Seek(0, 0)

	// Write back the updated CSV
	writer := csv.NewWriter(file)
	return writer.WriteAll(records)
}

func parseTime(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

func parseCost(s string) float64 {
	cost, _ := strconv.ParseFloat(s, 64)
	return cost
}
