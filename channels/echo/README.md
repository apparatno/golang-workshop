# Tasks

Below are some tasks you can try out on your own.
They are ordered by complexity but you can solve
them in any order you like.
If you already know some Go you can probably skip
the first ones.

## Channels

Create a small program that prompts for input
and outputs what you wrote back to you.
Write the program so that it will show the prompt
again and let the user type a new message.

>
> $> ./my-program
>
> Enter text:
>

Simulate work by introducing a delay before printing
the message to the screen.
Use a channel and a goroutine to "work" on the message
and print it while accepting new messages from the
prompt at the same time.


> Tips:
>
> Don't be bothered by the program writing on top of
> your prompt.
>
> Have a look at the `bufio` package for reading
> keyboard input.
>
> Use `time.Sleep` to delay execution.

### Further work

What happens if you cancel the program before all messages
have been processed?
Can you do anything about it?
(Check out the `os/signal` package and `sync.Waitgroup`)
