func carFleet(target int, position []int, speed []int) int {
    infos := make([][2]int, len(position))
    for i := 0; i < len(position); i++ {
        infos[i][0] = position[i]
        infos[i][1] = speed[i]
    }

    sort.Slice(infos, func(i int, j int) bool {
        return infos[i][0] > infos[j][0]
    })

    // fmt.Println(infos)
    fleet := []float64{}
    for _, info := range infos {
        t := float64(target - info[0]) / float64(info[1])
        fleet = append(fleet, t)
        if len(fleet) >= 2 && fleet[len(fleet) - 2] >= t {
            fleet = fleet[:len(fleet)-1]
        }
    }
    return len(fleet)
}