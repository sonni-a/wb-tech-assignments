package storage

import (
	"testing"
	"time"
)

func TestCreateAndGet(t *testing.T) {
	cal := NewCalendar()

	date, _ := time.Parse("2006-01-02", "2025-01-01")

	e, _ := cal.Create(1, date, "test")

	if e.ID == 0 {
		t.Fatal("expected id")
	}

	events := cal.EventsForPeriod(1, date, date)

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}
}

func TestUpdate(t *testing.T) {
	cal := NewCalendar()

	date, _ := time.Parse("2006-01-02", "2025-01-01")
	e, _ := cal.Create(1, date, "old")

	if err := cal.Update(e.ID, "new"); err != nil {
		t.Fatal(err)
	}

	events := cal.EventsForPeriod(1, date, date)
	if len(events) != 1 || events[0].Text != "new" {
		t.Fatalf("expected updated event, got %+v", events)
	}

	if err := cal.Update(999, "x"); err == nil {
		t.Fatal("expected error for missing event")
	}
}

func TestDelete(t *testing.T) {
	cal := NewCalendar()

	date := time.Now()
	e, _ := cal.Create(1, date, "test")

	err := cal.Delete(e.ID)
	if err != nil {
		t.Fatal(err)
	}

	err = cal.Delete(e.ID)
	if err == nil {
		t.Fatal("expected error")
	}
}
