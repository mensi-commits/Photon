package core

type ServiceSpec struct {
    ID          string
    Image       string
    MinReplicas int
    MaxReplicas int
    TargetCPU   float64
}

type Pod struct {
    ID        string
    ServiceID string
    NodeID    string
    Status    string
}

type Node struct {
    ID       string
    CPUUsage float64
    MemUsage float64
    Alive    bool
}