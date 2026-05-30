class Solution {
public:
    int carFleet(int target, vector<int>& position, vector<int>& speed) {
        vector<std::pair<int, int>> info(position.size());
        for (int i = 0; i < position.size(); i++) {
            info[i] = make_pair(position[i], speed[i]);
        }
        sort(info.begin(), info.end(), [](pair<int, int> &a, pair<int, int> &b){
            return a.first > b.first;
        });

        vector<float> times;
        for (const auto &i : info) {
            cout << "position: " << i.first << endl;
            cout << "speed: " << i.second << endl;
            float t = float(target - i.first) / i.second;
            cout << "t: " << t << endl;
            times.push_back(t);
            if (times.size() >= 2 && times[times.size() - 2] >= t) {
                cout << "last item: " << times[times.size() - 2] << endl;
                times.pop_back();
            }
        }

        return times.size();
    }
};
