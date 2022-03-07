# otp_service
otp service written in go lang 


Mutation 
 *sendOTP(input: sendOTP!): Status!
 *verifyOTP(input: verifyOTP!): Status!
 *addOtpType(input: addOtpType!): Status!
 *removeOtpType(input: removeOtpType!): Status!
 
Query
  *getOtpType: [OtpType]!

Thing we can do 
* Add otp type 
* remove otp type
* get all otp type 
* send otp
* verify otp



ToDO 
Adding email and sms sending layer

Boot 
go run server.go
open http://localhost:3001/ 

note: make sure redis is running in your local machine 


examples :

mutation{
  addOtpType(input:{otpLength:6,keyPrefix:"pop-otp",otp_validity_minutes:10,retry_limit:3, resent_limit:5 ,mobile_message: "Your OTP to complete your mobile verification on CaratLane is {{otp}}. It will be valid for the next {{validity}} mins. Please do not share your OTP with anyone.",retry_lock_minutes:30,resend_lock_minutes:5,otpType:"pop"}){
    status
    message
    other_message
  }
}

mutation{
  sendOTP(input:{channel:"8123456789",otp_type:"pop"}){
    status
    message
  }
}	

mutation {
  verifyOTP(input:{otp:"434497",otp_type:"pop",channel:"8123456789"})
{
  status
  message
}
}



{
  getOtpType{
    otpType
    keyPrefix
    retry_limit
    resent_limit
    mobile_message
    retry_lock_minutes
    resend_lock_minutes
    
  }
}

