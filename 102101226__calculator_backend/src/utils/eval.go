package utils

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/cloudwego/kitex/pkg/klog"
)

var (
	ValidChar = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", "+", ".", "*", "/", "(", ")"}
	OtherChar = []string{"sin", "cos", "tan", "asin", "acos", "atan", "sqrt", "log", "ln"}
)

func Calculate(exp string) (string, error) {
	line := ""
	index := 0
	flag := 0
	flag0 := 0
	next_index := 0
	//right_cnt := 0
	left_cnt := 0
	exp = strings.ReplaceAll(exp, "^", "**")
	for {
		if !contain(string(exp[index])) {
			// TODO:科学计数法识别
			for i := index + 2; i < len(exp); i++ {
				if string(exp[i]) == "(" && flag == 0 {
					next_index = i
					flag = 1
					left_cnt++
				} else if string(exp[i]) == "(" {
					left_cnt++
				} else if string(exp[i]) == ")" {
					left_cnt--
					if left_cnt == 0 {
						flag = 0
						res := 0.0
						newline := exp[next_index+1 : i]
						klog.Info(newline)
						str, err := Calculate(newline)
						if err != nil {
							klog.Info(err.Error())
							return "", err
						}
						f, err := strconv.ParseFloat(str, 64)
						if err != nil {
							klog.Info(err.Error())
							return "", err
						}
						if string(exp[index]) != "a" {
							// TODO:整合到operateAT函数
							res = f
							klog.Info(exp[index : index+2])
							if string(exp[index:index+3]) == OtherChar[7] {
								klog.Info("log:", res)
								res = math.Log10(res)
								klog.Info("log:", res)
							} else if string(exp[index:index+2]) == OtherChar[8] {
								res = math.Log1p(res)
							} else if string(exp[index+1]) != "q" {
								res = f * math.Pi / 180
							}

							if exp[index:index+3] == OtherChar[0] {
								res = math.Sin(res)
							} else if exp[index:index+3] == OtherChar[1] {
								res = math.Cos(res)
							} else if exp[index:index+3] == OtherChar[2] {
								res = math.Tan(res)
							} else if exp[index:index+4] == OtherChar[6] {
								res = math.Sqrt(res)
							}
						} else if string(exp[index]) == "a" {
							ret, err := operateAT(exp[index:index+4], f)
							if err != nil {
								klog.Error(err.Error())
								return err.Error(), err
							}
							res = ret
						}
						str = fmt.Sprint(res)
						line += exp[flag0:index] + str
						index = i
						flag0 = i + 1
						break
					}
				}
			}
		}
		index++
		if index >= len(exp) {
			break
		}
	}
	line += exp[flag0:]
	fmt.Println(line)
	expr, err := govaluate.NewEvaluableExpression(line)
	if err != nil || expr == nil {
		klog.Error("[expression] error:", err.Error())
		return "", err
	}
	result, err := expr.Evaluate(nil)
	return fmt.Sprint(result), err
}

func contain(char string) bool {
	for _, v := range ValidChar {
		if char == v {
			return true
		}
	}
	return false
}

func operateAT(str string, target float64) (float64, error) {
	if str == OtherChar[3] {
		if target > 1 || target < (-1) {
			klog.Error("[asin] out of range")
			return 0, errors.New("[asin] out of range")
		}
		target = math.Asin(target)
	} else if str == OtherChar[4] {
		if target > 1 || target < (-1) {
			klog.Error("[acos] out of range")
			return 0, errors.New("[acos] out of range")
		}
		target = math.Acos(target)
	} else if str == OtherChar[5] {
		target = math.Atan(target)
	}
	return target, nil
}
