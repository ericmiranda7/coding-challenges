class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        s, e = 0, len(matrix) - 1

        while s <= e:
            mid = (s+e) // 2

            if target < matrix[mid][0]:
                e = mid - 1
            elif target > matrix[mid][-1]:
                s = mid + 1
            else:
                # target row, binsearch now
                target_arr = matrix[mid]
                l, r = 0, len(target_arr) - 1
                
                while l <= r:
                    mid = (l+r) // 2

                    if target < target_arr[mid]:
                        r = mid - 1
                    elif target > target_arr[mid]:
                        l = mid + 1
                    else:
                        return True

                return False

        return False