var.1 = c(0,1,2,3)
var.2 <- c("learn","R")
c(TRUE,1) -> var.3
print(var.1)
cat ("var.1 is ", var.1 ,"\n")
cat ("var.2 is ", var.2 ,"\n")
cat ("var.3 is ", var.3 ,"\n")

ls()
rm(var.3)


# 
v <- c( 2,5.5,6)
t <- c(8, 3, 4)
#
print(v+t)
print(v-t)
print(v*t)
print(v/t)
print(v%%t)
print(v%/%t)
print(v^t)
#
print(v>t)
print(v < t)
print(v==t)
print(v<=t)
print(v>=t)
print(v!=t)
#
v1 <- c(3,1,TRUE,2+3i)
v2 <<- c(3,1,TRUE,2+3i)
v3 = c(3,1,TRUE,2+3i)
print(v1)
print(v2)
print(v3)
c(3,1,TRUE,2+3i) -> v1
c(3,1,TRUE,2+3i) ->> v2 
print(v1)
print(v2)

v <- 2:8
print(v) 

v1 <- 8
v2 <- 12
t <- 1:10
print(v1 %in% t) 
print(v2 %in% t) 

M = matrix( c(2,6,5,1,10,4), nrow=2,ncol=3,byrow = TRUE)
t = M %*% t(M)
print(t) 


