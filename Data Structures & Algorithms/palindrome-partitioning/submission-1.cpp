class Solution {
public:
	string s;
	vector<string> curr;
	vector<vector<string>> result;
    vector<vector<string>> partition(string s) {
        this->s = s;
		backtrack(0);
		return this->result;
    }

	void backtrack(int i) {
		if (i == s.size()) {
			result.push_back(curr);
			return;
		}

		for (int j = i; j < s.size(); j++) {
			if (!isPalindrome(s, i, j)) {
				continue;
			}

			curr.push_back(s.substr(i, j - i + 1));
			// std::cout << "c: ";
			// for (auto &c : curr) {
			//	std::cout << c << " ";
			// }
			std::cout << std::endl;
			backtrack(j + 1);
			curr.pop_back();
		}
	}

	bool isPalindrome(string s, int left, int right) {
		while (left < right) {
			if (s[left] != s[right]) {
				return false;
			}
			left++;
			right--;
		}
		return true;
	}
};
