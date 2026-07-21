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

	var rooms [][]Interval
	for _, interval := range intervals {
		// fmt.Println(interval)
		if len(rooms) == 0 {
			rooms = append(rooms, []Interval{})
			rooms[0] = append(rooms[0], interval)			
		} else {
			inserted := false
			for idx, room := range rooms {
				if room[len(room)-1].end <= interval.start {
					rooms[idx] = append(rooms[idx], interval)
					inserted = true
					break
				}
			}
			if !inserted {
				rooms = append(rooms, []Interval{})
				rooms[len(rooms)-1] = append(rooms[len(rooms)-1], interval)
			}
		}
	}

	// fmt.Println(rooms)
	return len(rooms)
}
