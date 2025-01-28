package core

import (
	"encoding/json"
	"log"
	"time"
)

type Number interface{}
type Id interface{}
type ConvertibleTimestamp interface{}
type Duration interface{}
type Data map[string]interface{}

func parseTimestamp(tsIn ConvertibleTimestamp) time.Time {
	var ts time.Time
	switch v := tsIn.(type) {
	case string:
		parsedTime, err := time.Parse(time.RFC3339, v)
		if err != nil {
			log.Fatalf("Error parsing time: %v", err)
		}
		ts = parsedTime
	case time.Time:
		ts = v
	}

	// Set resolution to milliseconds instead of microseconds
	ts = ts.Truncate(time.Millisecond)

	// Add timezone if not set
	if ts.Location() == time.UTC {
		log.Printf("timestamp without timezone found, using UTC: %v", ts)
		ts = ts.UTC()
	}
	return ts
}

type Event struct {
	Id        Id        `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Duration  Duration  `json:"duration"`
	Data      Data      `json:"data"`
}

func NewEvent(id Id, timestamp ConvertibleTimestamp, duration Duration, data Data) *Event {
	if timestamp == nil {
		log.Println("Event initializer did not receive a timestamp argument, using now as timestamp")
		timestamp = time.Now().UTC()
	}
	return &Event{
		Id:        id,
		Timestamp: parseTimestamp(timestamp),
		Duration:  duration,
		Data:      data,
	}
}

func (e *Event) ToJSONDict() map[string]interface{} {
	jsonData := map[string]interface{}{
		"id":        e.Id,
		"timestamp": e.Timestamp.Format(time.RFC3339),
		"duration":  e.Duration.(time.Duration).Seconds(),
		"data":      e.Data,
	}
	return jsonData
}

func (e *Event) ToJSONString() string {
	data := e.ToJSONDict()
	jsonStr, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error converting to JSON string: %v", err)
	}
	return string(jsonStr)
}

func (e *Event) HasProp(propName string) bool {
	switch propName {
	case "id":
		return e.Id != nil
	case "data":
		return e.Data != nil
	case "timestamp":
		return !e.Timestamp.IsZero()
	case "duration":
		return e.Duration != nil
	default:
		return false
	}
}

func (e *Event) Equal(other *Event) bool {
	return e.Timestamp.Equal(other.Timestamp) && e.Duration == other.Duration && e.DataEqual(other.Data)
}

func (e *Event) DataEqual(otherData Data) bool {
	for k, v := range e.Data {
		if otherData[k] != v {
			return false
		}
	}
	return true
}

func (e *Event) LessThan(other *Event) bool {
	return e.Timestamp.Before(other.Timestamp)
}

func (e *Event) SetID(id Id) {
	e.Id = id
}

func (e *Event) SetData(data Data) {
	e.Data = data
}

func (e *Event) SetTimestamp(timestamp ConvertibleTimestamp) {
	e.Timestamp = parseTimestamp(timestamp)
}

func (e *Event) SetDuration(duration Duration) {
	switch v := duration.(type) {
	case time.Duration:
		e.Duration = v
	case float64:
		e.Duration = time.Duration(v) * time.Second
	case int:
		e.Duration = time.Duration(v) * time.Second
	default:
		log.Fatalf("Couldn't parse duration of invalid type %T", v)
	}
}
