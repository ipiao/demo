list_data <- list("Red", "Green", c(21,32,11), TRUE, 51.23, 119.1)
print(list_data[3])

list_data <- list(c("Jan","Feb","Mar"), matrix(c(3,9,5,1,-2,8), nrow=2), list("green",12.3))
names(list_data) <- c("1st Quarter", "A_Matrix", "A Inner list")
print(list_data)
print(list_data$A_Matrix)

list_data[4] <- "New element"
print(list_data[4])
list_data[4] <- NULL
print(list_data[4])
list_data[3] <- "updated element"
print(list_data[3])


list1 <- list(1,2,3)
list2 <- list("Sun","Mon","Tue")
merged.list <- c(list1,list2)
print(merged.list)

list1 <- list(1:5)
print(list1)
list2 <-list(10:14)
print(list2)
v1 <- unlist(list1)
v2 <- unlist(list2)
print(v1)
print(v2)
result <- v1+v2
print(result)

M <- matrix(c(3:14), nrow=4, byrow=TRUE)
print(M)

rownames = c("row1", "row2", "row3", "row4")
colnames = c("col1", "col2", "col3")
print(list(rownames, colnames))
matrix(c(3:14), nrow=4, byrow=TRUE, dimnames=list(rownames, colnames))

vector1 <- c(5,9,3)
vector2 <- c(10,11,12,13,14,15)
new.array <- array(c(vector1,vector2),dim=c(3,3,2))
print(new.array)
result <- apply(new.array, c(1), sum)
print(result)

