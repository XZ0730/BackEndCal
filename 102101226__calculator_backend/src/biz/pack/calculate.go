package pack

import (
	"github.com/XZ0730/tireCV/biz/model/calculate"
)

func SetCalculateResponse(code int64, message string, result string, resp *calculate.CalculateResponse) {

	resp.Base = &calculate.BaseResponse{
		Code:    code,
		Message: message,
	}
	resp.Result = result
}

func SetRateResponse(code int64, message string, result string, resp *calculate.RateResponse) {
	resp.Base = &calculate.BaseResponse{
		Code:    code,
		Message: message,
	}
	resp.Interest = result
}

func SetGetRateResponse(code int64, message string, result []string, resp *calculate.GetRateResponse) {
	resp.Base = &calculate.BaseResponse{
		Code:    code,
		Message: message,
	}
	resp.RateList = append(resp.RateList, result...)
}

func SetBaseResponse(code int64, message string, resp *calculate.BaseResponse) {
	resp.Code = code
	resp.Message = message
}

func SetHistoryResponse(code int64, message string, history []string, resp *calculate.HistoryResponse) {
	resp.Base = &calculate.BaseResponse{
		Code:    code,
		Message: message,
	}
	resp.History = history
}
