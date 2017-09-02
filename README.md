# GoExpressionLanguage
This is an expression language written in Go. Please read the below description for more information about how to download and use the GoExpressionLanguage.

#Getting started
- Usage of the plugin - if you just want to use the plugin please refer to the Downloads section below. In it you will find download links for the three supported platforms: Linus, Windows, Darwin
- Contribution to the plugin - if you want to extend the plugin, you need to download the source code of the project. For this you will need Go installed.

# Supported operations
- Mathematical expressions - used for calculating a mathematical expressions ( + 2 3 ). The essential part here is that there MUST be opening and closing brackets as it is in predicate style. The drawback is that the brackets MUST be separated with space from the operators and the operands. Here is a list of the supported operators:
Operator | Supported symbols | Description
--- | --- | ---
`+` | + | Sums a set of numbers
`-` | - | Substracts a set of numbers
`*` | * | Multiplies a set of numbers
`÷` | ÷ | Divides a set of numbers
`^` | ^ | Powers two numbers a^b
`sqrt` | sqrt | Squares a give number
`sin` | sin | Calculates the sinus of a give number
`cos` | cos | Calculates the cosinus of a give number
`tan` | tan | Calculates the tangens of a give number
`&&` | && | Performs AND operation
`||` | or | Performs OR operation

- Number base systems - used for converting from a given number in a number base system to another base system. If one want to use this feature, the variables ibase and obase must be specified. The ibase - used for setting an input base number system, obase - used for setting the output number base system. 
Note: The supported base systems are from 2 - 36.
Example usage: 
    > ibase 2
    > obase 10
    > 101 => 5

# Downloads
Mac OS X 64 bit | Windows 64 bit | Linux 64 bit
--- | --- | ---
[go-exp-0.0.1-darwin.bin](http://google.com) | [go-exp-0.0.1-windows.exe](https://google.com) | [go-exp-0.0.1-linux.bin](http://google.com) |
