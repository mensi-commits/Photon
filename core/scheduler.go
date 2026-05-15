package core

func PickNode(nodes map[string]Node) string {
    var best Node
    first := true

    for _, n := range nodes {
        if !n.Alive {
            continue
        }

        if first || n.CPUUsage < best.CPUUsage {
            best = n
            first = false
        }
    }

    return best.ID
}