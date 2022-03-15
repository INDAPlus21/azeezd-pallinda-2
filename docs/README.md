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

# Answers to "Task 2 - Many Senders; Many Receivers"

### What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the main function?
The program crashes because the channel is closed in the main goroutine before any of the producer goroutines are finished with using it. The `wgp.Wait()` assured that the main goroutine waits before closing, changing their order removes the assurance.

### What happens if you move the `close(ch)` from the main function and instead close the channel in the end of the function Produce?
It would work fine at first but then crashes at the end when the first goroutines hits the `close(ch)` ruining all other goroutines' channel communication.


### What happens if you remove the statement close(ch) completely?
The program would still run (although not correctly) because no other goroutine is stuck in receiving. The consumption goroutines use the `for range` channel receiving which avoids deadlocks.

### What happens if you increase the number of consumers from 2 to 4?
The program should finish faster.

### Can you be sure that all strings are printed before the program stops?
No, this is not assured because the main goroutine might exit early before all consumers had the chance to print due to the lack of any lock or wait group.