class TimeMap {
public:
	// map that pairs a key (string) to an array of (timestamp (int), value (string))
	unordered_map<string, vector<pair<int, string>>> timeMap;
    TimeMap() {
		
    }
    
    void set(string key, string value, int timestamp) {
        timeMap[key].push_back(pair(timestamp, value));
    }
    
    string get(string key, int timestamp) {
        if (timeMap.count(key) == 0) {
			return "";
		}

		const auto& arr = timeMap[key];
		int left = 0, right = arr.size();
		// cout << "here" << endl;
		while (left < right) {
			// cout << "left: " << left << endl;
			// cout << "right: " << right << endl;
			int mid = left + (right - left)/2;
			// cout << "mid: " << mid << endl;
			const auto& entry = arr[mid];
			int t = entry.first;
			string value = entry.second;

			if (t > timestamp) {
				right = mid;
			} else {
				left = mid + 1;
			}
		}
		left -= 1;
		if (left < 0) {
			return "";
		}
		// std::cout << left - 1 << std::endl;
		string res = arr[left].second;
		return res;
    }
};
