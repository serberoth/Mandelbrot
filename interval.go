package main

// Interval is an interval range type with start, end, and step fields.
// It offers the range method which returns a channel so that it can
// be used in conjunction within a for ... range expression.
type Interval struct {
	Start, End, Step float64
}

// Make an Interval instance with the provided start and end taking
// the specified number of steps (this will calculate the value for
// step)
func LinearSpacing(start, end float64, numSteps int) Interval {
	step := (end - start) / float64(numSteps)
	return Interval{start, end, step}
}

// Range is a method on Interval that returns a channel that is
// populated by the values within the interval.
func (i Interval) Range() <-chan float64 {
	c := make(chan float64)
	go func() {
		for t := i.Start; t < i.End; t += i.Step {
			c <- t
		}
		close(c)
	}()
	return c
}
