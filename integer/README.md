
## 四则运算（arithmetic）

### 描述
expression仅支持加减乘除以及小括号
    
    expression := "2*(3+4)+5*6"

### 解决方案
从左向右依次遍历整个表达式，使用两个栈来分别存储操作数和操作符。

ds存储操作数，ss存储操作符。

读取的字符分为四类：
1. 左括弧
2. 右括弧
3. 操作符
4. 操作数

#### 左括弧
碰到左括弧，直接将左括弧存入ss

#### 右括弧
碰到右括弧，要开始向前计算，直到把小括号里的值算出来。

#### 操作符
乘除法优先级高于加减法。
1. 如果ss栈顶的操作符是'('，则将该操作符入栈即可。
2. 如果当前待处理的操作符op的优先级高于ss栈顶的操作符，则将该操作符op入栈即可。
3. 否则循环从ds和ss中取出两个操作数和一个操作符计算，并将计算结果压回ds。直到ss栈顶的操作符优先级不高于待处理的操作符，此时将待处理的操作符op压入ss

#### 操作数
1. 当读取到一个字符介于'0'和'9'之间时，说明碰到了操作数，此时使用`fmt.Sscanf`将该操作数完整的读出来。并压入ds.
2. 通过该操作数计算该操作数的位数，以使表达式expression的游标向前正确移动。

#### 收尾
当遍历完整个表达式expression之后，会有两种情况：
1. ds和ss中还有一个操作符和两个操作数。
2. ds中有三个操作数，ss中有一个加减号和乘除号两个操作符。例如：['+','*']
最终要将这ds和ss中的数据计算完毕。

### Tips
`单元过程`：每个括号中的内容是一个单元过程。如果没有括号，整个表达式是一个单元过程。
1. 整个算法过程中，每个`单元过程`，ss中最多存储一个加减号和乘除号