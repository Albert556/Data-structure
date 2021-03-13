<!--
 * @Author: Albert
 * @Date: 2021-01-17 00:53:40
 * @LastEditors: Albert
 * @LastEditTime: 2021-01-17 09:44:28
 * @Desctiption:
-->
# leetcode

## LinkList
## 解题技巧
1. 反转链表，用三指针
    ```python
    while n:
        c.next = n.next
        n.next = p.next
        p.next = n
        n = c.next
    ```
2. 链表排序<br>
   归并排序O(nlogn),O(1)，用迭代的方法，直接归并，省略了分治的过程

3. 环形链表
   * 链表是否有环用快慢指针，相遇有环
   * 二次相遇的位置是环的入口
## Tree
### 遍历
* 递归中序遍历
    ```python
    def dfs_mid(root):
        if not (root): return

        dfs_mid(root.left)
        # 逻辑代码
        path.append(root.val)
        dfs_mid(root.right)
    ```


### 计算器
双栈做法可以通用
```go
    func calculate(s string) int {
        // 栈储存数字
        nums := []int{}
        // 栈储存操作数
        ops := []byte{}
        // map储存操作符优先级
        priority := map[byte]int{
            '-': 1,
            '+': 1,
            '*': 2,
            '/': 2,
            '%': 2,
            '^': 3,
        }
        // 删除所有空格
        s = strings.Replace(s, " ", "", -1)
        // 防止括号后第一个数字是负数，如果负号看作一个二元操作符，会导致少一个数，所以在负号前补0也可以起到统一操作
        s = strings.Replace(s, "(-", "(0-", -1)
        // 定义计算函数，用匿名函数可以使用闭包，解决切片传值的问题
        var calc func()
        calc = func() {
            // 传入的式子是空的
            if len(nums) < 2 {
                return
            }
            // 计算完成，包括0
            if (len(ops)) == 0 {
                return
            }
            // 注意两个操作数的顺序，先出栈的在操作符的右边，后出栈的在操作符的左边
            b := nums[len(nums)-1]
            nums = nums[:len(nums)-1]
            a := nums[len(nums)-1]
            nums = nums[:len(nums)-1]
            op := ops[len(ops)-1]
            ops = ops[:len(ops)-1]
            // 记录结果
            // 结果入栈
            switch op {
            case '+':
                nums = append(nums, a+b)
            case '-':
                nums = append(nums, a-b)
            case '*':
                nums = append(nums, a*b)
            case '/':
                nums = append(nums, a/b)
            case '^':
                nums = append(nums, int(math.Pow(float64(a), float64(b))))
            case '%':
                nums = append(nums, a%b)
            }
        }
        // 开始解析算式
        // 首先入栈0，防止第一个数是复数
        nums = append(nums, 0)
        // 按byte循环式子，32是分开3和2
        // 式子中没有中文所以用byte索引不影响结果
        // 中文是rune类型，用range索引
        for i := 0; i < len(s); i++ {
            // 先解决括号的问题
            // ( 直接入栈
            // ) 相当于优先级提前，要先计算括号中的内容
            switch s[i] {
            case '(':
                ops = append(ops, s[i])
            case ')':
                for len(ops) > 0 {
                    // 循环计算直至遇到(
                    if ops[len(ops)-1] != '(' {
                        calc()
                    } else {
                        // 遇到 ( 出栈并结束计算的循环，继续遍历算式
                        ops = ops[:len(ops)-1]
                        break
                    }
                }
            default:
                // 括号解决完就是正常的数字和加减乘除
                // 先处理数字
                if '0' <= s[i] && s[i] <= '9' {
                    // 保存遍历到的数字，先不入栈，防止数字不止一位
                    num := int(s[i] - '0')
                    // 判断下一位是否是数字
                    j := i + 1
                    for ; j < len(s) && '0' <= s[j] && s[j] <= '9'; j++ {
                        // 是数字就和上一个组合
                        num = num*10 + int(s[j]-'0')
                    }
                    // 入栈，并修正i的值，j是判断的下一位，当j不是数字出循环时，j-1才是数的最后一位
                    // 不能把j直接赋值给i因为最外层的大循环会自动+1
                    nums = append(nums, num)
                    i = j - 1
                } else {
                    // 操作符
                    // 判断将要入栈的操作符和最后一个已经入栈的操作符的优先级
                    // +号入栈前，已入栈并在）上的操作符中不能有*号
                    // 3*2+3，+入栈时要把3*2的结果计算出来，直接入栈就会把式子变为3*(2+3)
                    for len(ops) > 0 && ops[len(ops)-1] != '(' {
                        preOps := ops[len(ops)-1]
                        // 有一个新操作要入栈时，先把栈内可以算的都算了
                        // 只有满足「栈内运算符」比「当前运算符」优先级高/同等，才进行运算
                        // 防止出现2-30-56变为2-(30-56)
                        if priority[preOps] >= priority[s[i]] {
                            calc()
                        } else {
                            break
                        }
                    }
                    // 无论是否需要计算，最后都要把操作符入栈
                    ops = append(ops, s[i])
                }
            }
        }
        // 算式遍历完成，所有元素入栈，计算
        for len(ops) > 0 {
            calc()
        }
        return nums[len(nums)-1]
    }
```

## 疑问
1. go中```'1'```和```byte(1)```的区别(leetcode使用```'1'```通过（1.15）,本地```byte(1)```通过（1.16）)
<details>
  <summary>答案</summary>
  （占位）
</details>

