type entry struct {
	timestamp int
	value     string
}

type TimeMap struct {
	entries map[string][]entry
}

func Constructor() TimeMap {
	return TimeMap{
		entries: make(map[string][]entry),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.entries[key]; !ok {
		this.entries[key] = make([]entry, 0)
	}

	this.entries[key] = append(this.entries[key], entry{timestamp: timestamp, value: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	entries, ok := this.entries[key]
	if !ok {
		return ""
	}

	left := 0
	right := len(entries)
	for left < right {
		mid := left + (right-left)/2
		currTimestamp := entries[mid].timestamp
		if currTimestamp > timestamp {
			right = mid
		} else {
			left = mid + 1
		}
	}
	idx := left - 1 // rightmost timestamp <= query timestamp
	if idx < 0 {
		return ""
	}

	return entries[idx].value
}
