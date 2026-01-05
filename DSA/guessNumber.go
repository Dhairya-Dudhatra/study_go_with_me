/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		mid := (l + r) / 2
		currentGuess := guess(mid)
		switch currentGuess {
		case 0:
			return mid
		case -1:
			r = mid - 1
		case 1:
			l = mid + 1
		}
	}
	return 0
}