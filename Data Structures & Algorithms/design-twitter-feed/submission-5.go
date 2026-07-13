type PriorityQueue []*Tweet

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}
func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].createdAt > pq[j].createdAt
}
func (pq *PriorityQueue) Push(x any) {
    (*pq) = append(*pq, x.(*Tweet))
}
func (pq *PriorityQueue) Pop() any {
    tmp := (*pq)
    n := len(tmp)
    x := tmp[n-1]
    (*pq) = tmp[:n-1]
    return x
}

type Tweet struct {
    id int
    createdAt int
}

type User struct {
    id int
    tweets []*Tweet
    followees map[int]*User
}

type Twitter struct {
    users map[int]*User
    currentTime int
}

func Constructor() Twitter {
    return Twitter{
        users: map[int]*User{},
    }
}

func (this *Twitter) createUser(userId int) *User {
    user := &User{
        id: userId,
        tweets: []*Tweet{},
        followees: map[int]*User{},
    }
    this.users[userId] = user
    return user
}

func (this *Twitter) PostTweet(userId int, tweetId int)  {
    user, exists := this.users[userId]
    if !exists {
        user = this.createUser(userId)
    }
    this.currentTime++
    user.tweets = append(user.tweets, &Tweet{
        id: tweetId,
        createdAt: this.currentTime,
    })
}

func (this *Twitter) GetNewsFeed(userId int) []int {
    user, exists := this.users[userId]
    if !exists {
        return []int{}
    }

    pq := &PriorityQueue{}
    heap.Init(pq)
    for _, tweet := range user.tweets {
        heap.Push(pq, tweet)
    }
    for _, followee := range user.followees {
        for _, tweet := range followee.tweets {
            heap.Push(pq, tweet)
        }
    }
    res := []int{}
    for pq.Len() > 0 && len(res) < 10 {
        x := heap.Pop(pq)
        tweet := x.(*Tweet)
        res = append(res, tweet.id)
    }
    return res
}

func (this *Twitter) Follow(followerId int, followeeId int)  {
    if followerId == followeeId {
        return
    }

    user, exist := this.users[followerId]
    if !exist {
        user = this.createUser(followerId)
    }
    followee, exist := this.users[followeeId]
    if !exist {
        followee = this.createUser(followeeId)
    }
    user.followees[followeeId] = followee
}

func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    user, exists := this.users[followerId]
    if !exists {
        return
    }
    delete(user.followees, followeeId)
}
