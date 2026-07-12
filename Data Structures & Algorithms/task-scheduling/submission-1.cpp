class Solution {
public:
    int leastInterval(vector<char>& tasks, int n) {
        priority_queue<pair<int, int>> pq;
        vector<int> count(26, 0);
        for (const auto& t : tasks) {
            count[t - 'A']++;
        }

        for (const auto& c : count) {
            if (!c) {
                continue;
            }
            pq.push(pair(c, 0));
        }

        queue<pair<int, int>> queue_;
        int cyclesTime = 0;
        while (pq.size() > 0 || queue_.size()) {
            cyclesTime++;
            if (pq.size() > 0) {
                auto t = pq.top();
                pq.pop();
                t.first--;
                if (t.first > 0) {
                    t.second = cyclesTime + n;
                    queue_.push(t);
                }
            }
            if (queue_.size() > 0) {
                auto t = queue_.front();
                if (t.second <= cyclesTime) {
                    queue_.pop();
                    pq.push(t);
                }
            }
        }
        return cyclesTime;
    }
};