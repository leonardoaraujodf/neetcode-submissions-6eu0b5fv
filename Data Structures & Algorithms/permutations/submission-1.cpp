class Solution {
public:
    vector<vector<int>> permute(vector<int>& nums) {
        vector<vector<int>> result;
		function<void(vector<int>)> backtrack;
		backtrack = [&](vector<int> curr) {
			if (curr.size() == nums.size()) {
				result.push_back(curr);
			}

			for (const auto& num : nums) {
				if (find(curr.begin(), curr.end(), num) == curr.end()) {
					curr.push_back(num);
					backtrack(curr);
					curr.pop_back();
				}
			}
		};
		backtrack(vector<int>{});
		return result;
    }
};
