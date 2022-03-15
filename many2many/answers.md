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