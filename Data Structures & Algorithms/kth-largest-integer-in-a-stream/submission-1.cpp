class KthLargest {
public:
	priority_queue<int, vector<int>, greater<int>> pq;
	int k;
    KthLargest(int k, vector<int>& nums) {
        for (const auto& n : nums) {
			pq.push(n);
		}
		this->k = k;
		while (pq.size() > k) {
			pq.pop();
		}
    }
    
    int add(int val) {
        pq.push(val);
		if (pq.size() > k) {
			pq.pop();
		}
		int res = pq.top();
		return res;
    }
};
