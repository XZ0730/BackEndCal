namespace go calculate


struct BaseResponse{
    1: required i64 code,
    2: required string message,
}

struct CalculateRequest {
    1: required string expression,
}
struct CalculateResponse {
    1: required BaseResponse base
    2: required string result,
}

struct RateRequest {
    1: required string money,
    2: required i64 type, // 0 存款 1 贷款
    3: required i64 duration, 
    // 1三个月 2半年 3一年 4二年 5三年 6五年
    // 7 一至三年 8三至五年 9五年
}

struct RateResponse {
    1: required BaseResponse base
    2: required string interest,
}

struct GetRateRequest{
}
struct GetRateResponse{
    1: BaseResponse base
    2: list<string> rate_list
}

struct SetRateRequest{
    1: required i64 the_type
    2: required string data
}

struct HistoryRequest{
}

struct HistoryResponse{
    1: required BaseResponse base
    2: required list<string> history
}

service CalculateService {
    CalculateResponse Calculate(1: CalculateRequest req) (api.post="/cs/calculate")
    RateResponse RateCall(1: RateRequest req)(api.post="/cs/rate")
    BaseResponse SetRate(1:SetRateRequest req)(api.post="/cs/rate/set")
    GetRateResponse GetRate(1:GetRateRequest req)(api.get="/cs/rate/get")

    HistoryResponse GetHistory(1:HistoryRequest req)(api.get="/cs/calculate/history")
}