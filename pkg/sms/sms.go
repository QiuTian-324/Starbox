package sms

//type SmsOperation interface {
//	SendVerificationCode(phoneNumber string) error
//	CheckVerificationCode(phoneNumber, verificationCode string) error
//}
//
//func NewSms() SmsOperation {
//	return getAliyunEntity()
//}
//
//// CreateRandCode 创建6位随机数
//func CreateRandCode() string {
//	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
//}
