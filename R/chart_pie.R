x<-c(21,62,10,53)
labels <- c("London","New York","Singapore","Mumbai")

png(file="city.jpg")

pie(x,labels)

dev.off()
