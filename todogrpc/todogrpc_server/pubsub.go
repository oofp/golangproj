package main

import (
	"log"
	pb "oofp/todogrpc/todoproto"
	"sync"
)

type PubSub struct {
	mu      sync.Mutex
	lastKey int64
	subs    map[int64](func(*pb.ToDoUpdate) bool)
}

func NewPubSub() *PubSub {
	ps := PubSub{
		subs: make(map[int64]func(*pb.ToDoUpdate) bool),
	}
	return &ps
}

func (ps *PubSub) AddSub(sub func(*pb.ToDoUpdate) bool) int64 {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.lastKey++

	ps.subs[ps.lastKey] = sub

	return ps.lastKey
}

func (ps *PubSub) RemoveSub(subId int64) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	delete(ps.subs, subId)
}

func (ps *PubSub) Publish(upd *pb.ToDoUpdate) {
	ns := ps.preparePublish(upd)

	for _, notif := range ns {
		notif()
	}
}

func (ps *PubSub) preparePublish(upd *pb.ToDoUpdate) []func() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	log.Printf("preparePublish subs:%d", len(ps.subs))

	ns := make([]func(), 0)
	for subId, sub := range ps.subs {
		curSubId := subId
		sub_func := sub
		notif := func() {
			log.Printf("going to notify subId:%d", curSubId)
			fl := sub_func(upd)
			if !fl {
				log.Printf("going to remove subId:%d", curSubId)
				ps.RemoveSub(subId)
			}
		}

		ns = append(ns, notif)
	}

	log.Printf("preparePublish subs to notify:%d", len(ns))
	return ns
}
