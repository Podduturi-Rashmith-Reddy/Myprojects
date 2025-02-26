Interface:
An interface in Go is a type that defines a set of method signatures but does not implement them.
It allows different types to implement these methods in their own way, enabling polymorphism and 
decoupling of code.

Goroutines and Channels in Go:
Goroutines are lightweight concurrent threads in Go that help run multiple tasks concurrently, making programs more efficient. Channels are the primary communication between these goroutines, allowing them to send and receive data safely.

Buffered vs. Unbuffered Channels:
Buffered Channels:
A buffered channel has a fixed capacity and can hold multiple values before requiring a receiver.
The sender can send data without blocking if there is space in the buffer. If the buffer is complete, the sender blocks until space is available. The receiver blocks only when the channel is empty. Suitable for asynchronous communication, where the sender and receiver do not need to operate simultaneously.

Deadlock Example: If the sender tries to send data to an entire buffered channel and no goroutine is ready to receive, the sender will block, causing a deadlock.

Unbuffered Channels:
Unbuffered channels have no internal storage. Data is passed directly between the sender and receiver. Both the sender and receiver need to be ready at the same time. If either is not prepared, both goroutines will block. It is useful when synchronization between the sender and receiver is necessary.  Unbuffered channels can sometimes perform better since they ensure synchronization between the sending and receiving goroutines.

Choosing Between Buffered and Unbuffered Channels:
Use buffered channels when you know the number of goroutines in advance and want to limit the work or capacity.
Use unbuffered channels when exchanging goroutines, which must happen simultaneously, ensuring synchronization.


