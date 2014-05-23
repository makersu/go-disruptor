package main

import "github.com/smartystreets/go-disruptor"

func publish(writer *disruptor.SharedWriter) {

	for {
		// TODO: return lower, upper instead? or some kind of struct "Reservation"
		// upon which commit can be invoked?
		if sequence := writer.Reserve(ItemsToPublish); sequence != disruptor.Gating {
			ringBuffer[sequence&RingMask] = sequence
			writer.Commit(sequence, sequence+ItemsToPublish-1)
		}
	}
}
