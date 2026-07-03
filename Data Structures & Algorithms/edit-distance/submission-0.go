func minDistance(word1 string, word2 string) int {
    dp := make([][]int, len(word1)+1)
    for i := range dp {
        dp[i] = make([]int, len(word2)+1)
    }

    for i := 0; i <= len(word1); i++ {
        for j := 0; j <= len(word2); j++ {
            if i == 0 {
                dp[0][j] = j
                continue
            }
            if j == 0 {
                dp[i][0] = i
                continue
            }

            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
                continue
            }

            insertOp := dp[i][j-1]
            replaceOp := dp[i-1][j-1]
            deleteOp := dp[i-1][j]
            dp[i][j] = 1 + min(insertOp, replaceOp, deleteOp)
        }
    }

    return dp[len(word1)][len(word2)]
}