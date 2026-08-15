class Solution {
public:
	vector<int> subset;
	vector<vector<int>> result;
	vector<int> nums;
    vector<vector<int>> subsetsWithDup(vector<int>& nums) {
        this->nums = nums;
		sort(this->nums.begin(), this->nums.end());
		result.push_back(subset);
		backtrack(0);
		return result;
    }

	void backtrack(int idx) {
		if (idx == nums.size()) {
			return;
		}

		int prev = INT_MIN;
		for (int i = idx; i < nums.size(); i++) {
			int num = nums[i];
			if (prev == num) {
				continue;
			}
			// cout << "i: " << i << " idx: " << idx << " num: " << num << endl;
			subset.push_back(num);
			result.push_back(subset);
			backtrack(i + 1);
			subset.pop_back();
            prev = num;
		}
	}
};
