# Inspiration

Here are some experiments playing with buffers and channels in Go.

## Below is text from chatgpt:

In Go, channels are used for communication between concurrent goroutines. Buffers can be added to channels to allow them to store a limited number of values. This can be useful in situations where the sending or receiving of values on a channel may be slower than the other end.

To use a buffer in a channel, you need to create a buffered channel by specifying the buffer size when you make the channel:

```go
ch := make(chan int, 10)
```

In this example, ch is a buffered channel with a capacity of 10. You can now send and receive values on this channel just like any other channel. The difference is that if the channel is full, sending a value on the channel will block until there is space in the buffer. Similarly, if the channel is empty, receiving a value from the channel will block until a value is available.

Here's an example program that demonstrates how to use a buffered channel:

```go
package main

import "fmt"

func main() {
    // Create a buffered channel with a capacity of 3
    ch := make(chan int, 3)

    // Send three values into the channel
    ch <- 1
    ch <- 2
    ch <- 3

    // This fourth send will block until a value is received from the channel
    // to make space in the buffer
    ch <- 4

    // Receive values from the channel
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch) // This will block until a value is sent on the channel
}
```

In this program, we create a buffered channel with a capacity of 3. We send three values into the channel and then try to send a fourth value. Since the channel is full, the fourth send operation will block until one of the values is received from the channel.

Finally, we receive values from the channel using the <- operator. If the channel is empty, receiving a value will block until a value is available in the channel.

Note that you should be careful when using buffered channels, as they can lead to deadlocks if not used properly. Always make sure that the sender and receiver are in sync and that you don't accidentally fill up the buffer without clearing it.
