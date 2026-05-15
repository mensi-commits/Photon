package core

import "sync"

type State struct {
    Services map[string]ServiceSpec
    Pods     map[string]Pod
    Nodes    map[string]Node
    mu       sync.RWMutex
}

func NewState() *State {
    return &State{
        Services: map[string]ServiceSpec{},
        Pods:     map[string]Pod{},
        Nodes:    map[string]Node{},
    }
}