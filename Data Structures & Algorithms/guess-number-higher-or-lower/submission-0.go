/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    start := 1
    end := n

    for start <= end {
        m := ((end - start) / 2) + start

        if guess(m) == -1 {
            end = m - 1
        } else if guess(m) == 1 {
            start = m + 1
        } else {
            return m
        }
    }
    return -1
}
