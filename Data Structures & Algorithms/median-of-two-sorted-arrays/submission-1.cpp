class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
		if (nums1.size() > nums2.size()) {
			return findMedianSortedArrays(nums2, nums1);
		}
        vector<int>& A = nums1, B = nums2;
		int m = A.size();
		int n = B.size();
		cout << "m, n = " << m << ", " << n << endl;
		int left = 0, right = m;
		while (left <= right) {
			int i = left + (right - left)/2;
			int j = (m + n + 1)/2 - i;
			int A_left = INT_MIN;
			int A_right = INT_MAX;
			int B_left = INT_MIN;
			int B_right = INT_MAX;
			if (i > 0) {
				A_left = A[i-1];
			}
			if (i < m) {
				A_right = A[i];
			}
			if (j > 0) {
				B_left = B[j-1];
			}
			if (j < n) {
				B_right = B[j];
			}
			if (A_left <= B_right && B_left <= A_right) {
				if ((m+n) % 2 == 0) {
					return double(max(A_left, B_left) + min(A_right, B_right)) / 2;
				} else {
					return double(max(A_left, B_left));
				}
			} else if (A_left > B_right) {
				right = i - 1;
			} else {
				left = i + 1;
			}
		}

		return -1;
    }
};
