# Learn GO!

Complete list the thing that I learn about go programming. 

# Article (Learning Source)

- Understanding how the go routine works to avoid leak
 ..* https://medium.com/golangspec/goroutine-leak-400063aef468 
 ..* https://blog.detectify.com/2019/09/05/how-we-tracked-down-a-memory-leak-in-one-of-our-go-microservices/
 ..* https://hackernoon.com/avoiding-memory-leak-in-golang-api-1843ef45fca8
 
### Key Points
 
 - Go is very attached to memory that it allocates.
 - Careful with Go garbage collector or its runtime that maybe hogging the memory.
 - Always set the timeout if you use big concurrent call for your apps.
 - Be careful with channel! 
