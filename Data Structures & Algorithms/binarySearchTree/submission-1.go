type TreeMap struct {
    root *TreeNode
}

type TreeNode struct {
    key int
    val int
    left *TreeNode
    right *TreeNode
}

func NewTreeMap() *TreeMap {
    return &TreeMap{nil}
}

func (tm *TreeMap) Insert(key, val int) {
    tm.root = insertNode(tm.root, key, val)    
}

func insertNode(root *TreeNode, key, val int) *TreeNode {
    if root == nil || root.key == key {
        return &TreeNode{key, val, nil, nil}
    }

    if key < root.key {
        root.left = insertNode(root.left, key, val)
    } else {
        root.right = insertNode(root.right, key, val)
    }

    return root
}

func (tm *TreeMap) Get(key int) int {
    found := getNode(tm.root, key)
    if found == nil {
        return -1
    }
    return found.val
}

func getNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    if root.key == key {
        return root
    } else if key < root.key {
        return getNode(root.left, key)
    } else {
        return getNode(root.right, key)
    }
}

func (tm *TreeMap) GetMin() int {
    min := getMinNode(tm.root)
    if min == nil {
        return -1
    }
    return min.val
}

func getMinNode(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if root.left == nil {
        return root
    }

    return getMinNode(root.left)
}

func (tm *TreeMap) GetMax() int {
    return getMax(tm.root)
}

func getMax(root *TreeNode) int {
    if root == nil {
        return -1
    }
    if root.right == nil {
        return root.val
    }

    return getMax(root.right)
}

func (tm *TreeMap) Remove(key int) {
    tm.root = removeNode(tm.root, key)
}

func removeNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    if root.key == key {
        // if 0..1 child
        if root.left == nil {
            return root.right
        } else if root.right == nil {
            return root.left
        } else {
            // 2 children
            minNode := getMinNode(root.right)
            root.key = minNode.key
            root.val = minNode.val
            root.right = removeNode(root.right, minNode.key)
            return root
        }
    
    } else if key < root.key {
        root.left = removeNode(root.left, key)
    } else {
        root.right = removeNode(root.right, key)
    }
    return root
}

func (tm *TreeMap) GetInorderKeys() []int {
    res := []int{}
    inorder(tm.root, &res)
    return res
}

func inorder(root *TreeNode, res *[]int) {
    if root == nil {
        return
    }

    inorder(root.left, res)
    *res = append(*res, root.key)
    inorder(root.right, res)
}
























