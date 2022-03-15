# Answers to "Task 1 - Debugging Concurrent Programs"

These answers the current questions
1. Explain what is wrong with the code.
2. Fix the bug.
3. Explain why your solution fixes the bug.

## Bug 1
1. The code creates a deadlock when sending the message through the channel using `ch <- "..."` but it waits for someone to receive it which is on the next line and thus it blocks for ever.
2. See [./bug01.go](./bug01.go)
3. By making the send in a separate goroutine the main routine now is able to receive from the channel is now being sent to from the newly created goroutine.

## Bug 2
1. The code closes the channel before it printing the messages. That is, the delay before printing out creates a gap for the `close(ch)` to run on the last iteration before any the `Println` having the chance to run and thus the channel closes before the goroutine prints.
2. See See [./bug02.go](./bug02.go)
3. By adding a `WaitGroup` the program now increments the `WaitGroup` counter before the delay then `Done`s it after printing it. The `WaitGroup` waits before closing the channel. This ensures that the closing is done after all 11 `Println`s and delays have been done before closing the channel.