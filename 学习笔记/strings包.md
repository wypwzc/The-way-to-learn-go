Go 语言标准库中的 **`strings`** 包专门用于提供对字符串（String）的各种常用高效操作。由于 Go 语言中的字符串是不可变类型（Immutable），`strings` 包中的函数在处理字符串时，都不会直接修改原字符串，而是会**返回一个新的字符串或处理结果**。

`strings` 包的核心作用和常用功能主要可以分为以下几大类：

### 一、 查找与包含判断（Searching & Checking）

用来检测字符串中是否存在某个子串，或者确定子串的位置。

- **`strings.Contains(s, substr)`**：检查字符串 `s` 中是否包含子串 `substr`，返回布尔值。   
- **`strings.HasPrefix(s, prefix)`**：检查字符串 `s` 是否以 `prefix` 开头。
- **`strings.HasSuffix(s, suffix)`**：检查字符串 `s` 是否以 `suffix` 结尾。
- **`strings.Index(s, substr)`**：返回子串 `substr` 在 `s` 中第一次出现的位置索引（字节级），如果不存在返回 `-1`。
- **`strings.Count(s, substr)`**：统计子串 `substr` 在 `s` 中出现的次数。

### 二、 空白与字符修剪（Trimming）

用于清理字符串两端不需要的字符（如空格、特定标点等）。

- **`strings.TrimSpace(s)`**：剔除字符串开头和结尾的所有空白字符（包括空格、制表符 `\t`、换行符 `\n` 等）。   
- **`strings.Trim(s, cutset)`**：剔除字符串开头和结尾包含在 `cutset` 中的所有指定字符。
- **`strings.TrimLeft` / `strings.TrimRight`**：分别只修剪左边或右边的指定字符。

### 三、 切割与拼接（Splitting & Joining）

用于把长字符串拆分成多个片段，或者把多个片段组合成一个长字符串。

- **`strings.Split(s, sep)`**：以指定的分隔符 `sep` 将字符串切开，返回一个字符串切片（`[]string`）。   
- **`strings.Join(elems, sep)`**：将一个字符串切片中的所有元素用指定的分隔符 `sep` 拼接成一个完整的字符串。
- **`strings.Fields(s)`**：按照空白字符（空格、换行等）将字符串分割成多个单词切片。

### 四、 替换（Replacement）

用于修改或替换字符串中的指定内容。

- **`strings.Replace(s, old, new, n)`**：将字符串 `s` 中的前 `n` 个 `old` 替换为 `new`（如果 `n < 0`，则表示全部替换）。
- **`strings.ReplaceAll(s, old, new)`**：将字符串 `s` 中所有的 `old` 替换为 `new`（相当于 `n = -1` 的快捷方式）。

### 五、 大小写转换（Case Conversion）

- **`strings.ToUpper(s)`**：将字符串中的所有英文字母转换为大写。
- **`strings.ToLower(s)`**：将字符串中的所有英文字母转换为小写。

###  总结

在日常开发中，处理用户输入、解析文本配置、格式化输出时，`strings` 包是使用频率极高的基础工具包。使用时只需在文件头部导入：

```Go
import "strings"
```