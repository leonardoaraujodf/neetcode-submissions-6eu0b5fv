/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })

    rooms := make([][]Interval, 0)
    for _, interval := range intervals {
        if len(rooms) == 0 {
            rooms = append(rooms, []Interval{interval})
            continue
        } else {
            inserted := false
            for idx, room := range rooms {
                // fmt.Println("room:", room)
                // fmt.Println("room[len(room)-1]: ", room[len(room)-1])
                // fmt.Println("interval: ", interval)
                if interval.start >= room[len(room)-1].end {
                    // fmt.Println("Adding ", interval, "to ", room)
                    rooms[idx] = append(rooms[idx], interval)
                    inserted = true
                    // fmt.Println("After insertion: ", room)
                    break
                }
            }
            if !inserted {
                rooms = append(rooms, []Interval{interval})
                // fmt.Println("Creating new room: ", rooms[len(rooms)-1])
            }
        }
    }

    return len(rooms)
}
