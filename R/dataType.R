
# 逻辑 logical
v <- TRUE ;print(class(v))
# 数字 numeric
v <- 12 ;print(class(v))
# 整数 integer
v <- 12L ;print(class(v))
# 复数 complex
v <- 2+3i ;print(class(v))
# 字符 character
v <- "TRUE" ;print(class(v))
# 原始 character
v <- charToRaw("Hello");print(class(v))


# 向量 vector
color <- c(1,'a',3)
print(color);print(class(color))

# 列表 list
list1 <- list(c(2,5,3),21.3,sin)
print(list1);print(class(list1))

# 矩阵
M = matrix( c('a','a','b','c','b','a'), nrow=2,ncol=3,byrow = TRUE)
print(M);print(class(M))

# 数组
a <- array(c('green','yellow'),dim=c(3,3,2))
print(a);print(class(a))

# 因子
apple_colors <- c('green','green','yellow','red','red','red','green')
factor_apple <- factor(apple_colors)
print(factor_apple);print(nlevels(factor_apple))

# 数据帧
BMI <- 	data.frame(
			gender = c("Male", "Male","Female"), 
			height = c(152, 171.5, 165), 
			weight = c(81,93, 78),
			Age =c(42,38,26)
			)
print(BMI)
