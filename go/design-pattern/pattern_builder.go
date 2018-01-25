package designPattern

// Builder 建造者
type Builder struct {
	senderList []Sender
}

// ProduceMailSender 建造方法
func (b *Builder) ProduceMailSender(count int) {
	for i := 0; i < count; i++ {
		b.senderList = append(b.senderList, MailSender{})
	}
}

// ProduceSmsSender 建造方法
func (b *Builder) ProduceSmsSender(count int) {
	for i := 0; i < count; i++ {
		b.senderList = append(b.senderList, SmsSender{})
	}
}
