package storage

import (
	"errors"
	"sync"
	"time"
)

type Event struct {
	ID     int
	UserID int
	Date   time.Time
	Text   string
}

type Calendar struct {
	mu     sync.RWMutex
	events map[int]Event
	nextID int
}

func NewCalendar() *Calendar {
	return &Calendar{
		events: make(map[int]Event),
	}
}

func (c *Calendar) Create(userID int, date time.Time, text string) (Event, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.nextID++
	e := Event{
		ID:     c.nextID,
		UserID: userID,
		Date:   date,
		Text:   text,
	}

	c.events[e.ID] = e
	return e, nil
}

func (c *Calendar) Update(id int, text string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	e, ok := c.events[id]
	if !ok {
		return errors.New("event not found")
	}

	e.Text = text
	c.events[id] = e
	return nil
}

func (c *Calendar) Delete(id int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.events[id]; !ok {
		return errors.New("event not found")
	}

	delete(c.events, id)
	return nil
}

func (c *Calendar) EventsForPeriod(userID int, from, to time.Time) []Event {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var result []Event
	for _, e := range c.events {
		if e.UserID == userID &&
			!e.Date.Before(from) &&
			!e.Date.After(to) {
			result = append(result, e)
		}
	}
	return result
}
