package 青蛙跳台阶问题

func numWays(n int) int {
	k := 1000000007
	dp := make([]int, n+1)
	if n < 1 {
		return 1
	}
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % k
	}
	return dp[n]
}
