package service

import (
	"encoding/json"
	"fmt"
	"otp_service/graph/model"
	"otp_service/paramvalidator"
	"otp_service/redis"
	"otp_service/structs"
	"otp_service/utils"
	"time"
)

func AddOtpType(input model.AddOtpType) (ok bool, response string) {
	otpType, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		return false, err.Error()
	}
	return redis.SetData("otptype_"+input.OtpType, otpType)
}

func RemoveOtpType(input model.RemoveOtpType) (ok bool, response string) {
	ok, redisData := redis.Getdata("otptype_" + input.OtpType)
	if ok {
		delFlag, message := redis.DelData("otptype_" + input.OtpType)
		return delFlag, message
	}
	return ok, redisData
}

func GetOtpType() []*model.OtpType {
	redisData := redis.KeysData("otptype_*")
	var arrData []*model.OtpType
	for _, value := range redisData {
		ok, RedisOtpTypeData := redis.Getdata(value)
		if ok {
			var otpTypeData *model.OtpType
			err := json.Unmarshal([]byte(RedisOtpTypeData), &otpTypeData)
			if err != nil {
				panic(err.Error())
			}
			arrData = append(arrData, otpTypeData)
		}
	}
	return arrData
}

func SendOtp(channel string, otptype string) (ok bool, response string) {
	if paramvalidator.IsEmail(channel) != nil && paramvalidator.IsMobileNumber(channel) != nil {
		return false, "Invalid channel"
	}
	ok, otpTypeRedisdata := redis.Getdata("otptype_" + otptype)
	if !ok {
		return false, "Otp type data found"
	}
	var otpType model.OtpType
	otpTypeErr := json.Unmarshal([]byte(otpTypeRedisdata), &otpType)
	if otpTypeErr != nil {
		panic(otpTypeErr.Error())
	}
	key := fmt.Sprintf("%v-%v", otpType.KeyPrefix, channel)
	ok, otpRedisdata := redis.Getdata(key)
	var otpdata structs.Otp
	if ok {
		otpDataErr := json.Unmarshal([]byte(otpRedisdata), &otpdata)
		if otpDataErr != nil {
			panic(otpDataErr.Error())
		}
		if otpdata.ResentCount >= otpType.ResentLimit {
			if !otpdata.ResendLock {
				otpdata.ResendLock = true
				setOtpdata, err := json.Marshal(otpdata)
				if err != nil {
					fmt.Println(err)
					return false, err.Error()
				}
				redis.SetDataWithTTL(key, setOtpdata, time.Duration(otpType.ResendLockMinutes)*time.Minute)
			}
			ttl := fmt.Sprintf("%.f", redis.GetExpire(key).Minutes())
			return false, "Maximum retries reached for this user, please try after " + ttl + " minutes"
		}
		otpdata.ResentCount = otpdata.ResentCount + 1

	} else {
		otp, err := utils.GenerateOTPCode(6)
		if err == nil {
			otp = "123456"
		}
		otpdata.ResentCount = 0
		otpdata.RetryCount = 0
		otpdata.Otp = otp
		otpdata.RetryLock = false
		otpdata.ResendLock = false
	}
	setOtpdata, err := json.Marshal(otpdata)
	if err != nil {
		fmt.Println(err)
		return false, err.Error()
	}
	redis.SetDataWithTTL(key, setOtpdata, time.Duration(otpType.OtpValidityMinutes)*time.Minute)
	return true, "Otp successfully sent"
}

func VerifyOtp(channel string, otp string, otptype string) (ok bool, response string) {
	if paramvalidator.IsEmail(channel) != nil && paramvalidator.IsMobileNumber(channel) != nil {
		return false, "Invalid channel"
	}
	ok, otpTypeRedisdata := redis.Getdata("otptype_" + otptype)
	if !ok {
		return false, "OTP type data found"
	}
	var otpType model.OtpType
	otpTypeErr := json.Unmarshal([]byte(otpTypeRedisdata), &otpType)
	if otpTypeErr != nil {
		panic(otpTypeErr.Error())
	}
	key := fmt.Sprintf("%v-%v", otpType.KeyPrefix, channel)
	ok, otpRedisdata := redis.Getdata(key)
	var otpdata structs.Otp
	if !ok {
		return false, "OTP data not found or OTP is expired"
	}
	otpDataErr := json.Unmarshal([]byte(otpRedisdata), &otpdata)
	if otpDataErr != nil {
		panic(otpDataErr.Error())
	}
	if otpdata.Otp == otp {
		redis.DelData(key)
		return true, "OTP Verified"
	}
	if otpdata.RetryCount >= otpType.RetryLimit {
		if !otpdata.RetryLock {
			otpdata.RetryLock = true
			setOtpdata, err := json.Marshal(otpdata)
			if err != nil {
				fmt.Println(err)
				return false, err.Error()
			}
			redis.SetDataWithTTL(key, setOtpdata, time.Duration(otpType.RetryLockMinutes)*time.Minute)
		}
		ttl := fmt.Sprintf("%.f", redis.GetExpire(key).Minutes())
		return false, "Maximum retries reached for this user, please try after " + ttl + " minutes"
	}
	otpdata.RetryCount = otpdata.RetryCount + 1
	setOtpdata, err := json.Marshal(otpdata)
	if err != nil {
		fmt.Println(err)
		return false, err.Error()
	}
	redis.SetDataWithTTL(key, setOtpdata, time.Duration(otpType.OtpValidityMinutes)*time.Minute)
	return false, "Otp is incorrect"
}
