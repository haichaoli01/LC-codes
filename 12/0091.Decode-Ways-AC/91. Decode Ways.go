package ltcode

// 如果连续的两位数符合条件，就相当于一个上楼梯的题目，可以有两种选法：
// 1.一位数决定一个字母
// 2.两位数决定一个字母
// 就相当于dp(i) = dp[i-1] + dp[i-2];
// 如果不符合条件，又有两种情况
// 1.当前数字是0：
// 	不好意思，这阶楼梯不能单独走，
// 	dp[i] = dp[i-2]
// 2.当前数字不是0
// 	不好意思，这阶楼梯太宽，走两步容易扯着步子，只能一个一个走
// 	dp[i] = dp[i-1];
func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)

	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' { //lhcerr: s[i-1] != 0
			dp[i] = dp[i-1]
		}
		if i > 2 {
			if t := (s[i-2]-'0')*10 + (s[i-1] - '0'); t < 27 && t > 9 {
				dp[i] += dp[i-2]
			}
		}

	}
	return dp[n]
}
