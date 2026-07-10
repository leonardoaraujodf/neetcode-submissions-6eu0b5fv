class Solution {
public:
    struct Point {
        double distance;
        int idx;
        Point(double d, int i) : distance(d), idx(i) {}
        bool operator<(const Point& other) const {
            return this->distance < other.distance;
        }
    };
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        priority_queue<Point> pq;
        for (int i = 0; i < points.size(); i++) {
            double x = points[i][0];
            double y = points[i][1];
            double distance = sqrt(pow(x, 2) + pow(y, 2));
            auto p = Point(distance, i);
            pq.push(p);
        }
        while (pq.size() > k) {
            pq.pop();
        }

        vector<vector<int>> ans;
        while (pq.size() > 0) {
            Point p = pq.top();
            pq.pop();
            ans.push_back(points[p.idx]);
        }
        return ans;
    }
};