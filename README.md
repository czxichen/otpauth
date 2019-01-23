# 谷歌otpauth验证器

## Install

    go get -v github.com/czxichen/otpauth

## Example

    package main

    import (
        "fmt"

        "github.com/czxichen/otpauth"
    )

    func main() {
        str := GenerateOTP("Di", "dijielin@qq.com")

        fmt.Printf("OTPAUTH字符串: %s\n",str)

        if CompareCode(3, 865946, "MNTIZ73RIWUUO2PJ") {
            fmt.Println("Verification Pass!")
        } else {
            fmt.Println("Verification Faild!")
        }
    }