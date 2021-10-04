package path

/*
// Package path implements utility routines for manipulating slash-separated
// paths.
//
// The path package should only be used for paths separated by forward
// slashes, such as the paths in URLs. This package does not deal with
// Windows paths with drive letters or backslashes; to manipulate
// operating system paths, use the path/filepath package.
//path包实现了用于操作斜杠分隔(slash-separated)的实用程序例程
// path包只能用于前向路径斜杠，例如url中的路径。
path包不涉及带有驱动器号或反斜杠的Windows路径;
用于操作操作系统路径，使用path/filepath包。
*/

import (
	goPath "path"
)

// Clean returns the shortest path name equivalent to path
// by purely lexical processing. It applies the following rules
// iteratively until no further processing can be done:
//
//	1. Replace multiple slashes with a single slash.
//	2. Eliminate each . path name element (the current directory).
//	3. Eliminate each inner .. path name element (the parent directory)
//	   along with the non-.. element that precedes it.
//	4. Eliminate .. elements that begin a rooted path:
//	   that is, replace "/.." by "/" at the beginning of a path.
//
// The returned path ends in a slash only if it is the root "/".
//当参数只有"/"时，就返回斜杠"/"
// If the result of this process is an empty string, Clean
// returns the string ".".
//
// See also Rob Pike, ``Lexical File Names in Plan 9 or
// Getting Dot-Dot Right,''
// https://9p.io/sys/doc/lexnames.html
//Clean函数通过单纯的词法操作返回和path代表同一地址的最短路径
//它会不断的依次应用如下的规则，直到不能再进行任何处理
/*
	1. 将连续的多个斜杠替换为单个斜杠
	2. 剔除每一个.路径名元素（代表当前目录）
	3. 剔除每一个路径内的..路径名元素（代表父目录）和它前面的非..路径名元素(父级目录)
	4. 剔除开始一个根路径的..路径名元素，即将路径开始处的"/.."替换为"/"
	详情查看example
*/
func Clean(path string) string {
	return goPath.Clean(path)
}

// Split 根据文件的路径分割成路径和文件
func Split(path string) (dir, file string) {
	return goPath.Split(path)
}

// Join 合并子路径，并Clean返回路径
func Join(element ...string) string {
	return goPath.Join(element...)
}

// Ext 返回路径的文件的扩展名
func Ext(path string) string {
	return goPath.Ext(path)
}

// Base 返回path的最后一个元素
func Base(path string) string {
	return goPath.Base(path)
}

// Dir 返回path的目录，并且经过Clean
func Dir(path string) string {
	return goPath.Dir(path)
}

// IsAbs 判断path是否为绝对路径
func IsAbs(path string) bool {
	return goPath.IsAbs(path)
}

// Match 匹配方法，如果name匹配shell文件名模式匹配字符串，Match函数返回真
//（注意：Match要求匹配整个name字符串，不是它的一部分。如果pattern语法错误时，会返回"syntax error in pattern"的错误）
/*
	'*'                                  匹配0或多个非/的字符
	'?'                                  匹配1个非/的字符
	'[' [ '^' ] { character-range } ']'  字符组（必须非空）(支持三种格式[abc],[^abc],[a-c])
	c                                    匹配字符c（c != '*', '?', '\\', '['）
	'\\' c                               匹配字符c(可上面c的区别是 可以支持字符 * ? \\ [的匹配)

*/
func Match(pattern, name string) (bool, error) {
	return goPath.Match(pattern, name)
}
