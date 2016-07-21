# Go (Golang) Channels

Type of `ch`            | Status        | Operation   | Result                         | References
----------------------- | ------------- | ----------- | ------------------------------ | ----------
Bidirectional (`chan`)  | not closed    | `ch <- v`   | Sends normally                 |
Bidirectional (`chan`)  | not closed    | `<-ch`      | Receives normally              |
Bidirectional (`chan`)  | not closed    | `close(ch)` | Closes the channel             |
Bidirectional (`chan`)  | closed        | `ch <- v`   | Causes a Run-Time Panic        | [close]
Bidirectional (`chan`)  | closed        | `<-ch`      | Returns zero value and `false` | [close], [recv]
Bidirectional (`chan`)  | closed        | `close(ch)` | Causes a Run-Time Panic        | [close]
Send Only (`chan<-`)    | any           | `<-ch`      | Compilation Error              |
Send Only (`chan<-`)    | not closed    | `ch <- v`   | Sends normally                 |
Send Only (`chan<-`)    | not closed    | `close(ch)` | Closes the channel             |
Send Only (`chan<-`)    | closed        | `ch <- v`   | Causes a Run-Time Panic        | [close]
Send Only (`chan<-`)    | closed        | `close(ch)` | Causes a Run-Time Panic        | [close], [send]
Receive Only (`<-chan`) | any           | `ch <- v`   | Compilation Error              |
Receive Only (`<-chan`) | any           | `close(ch)` | Compilation Error              | [close]
Receive Only (`<-chan`) | not closed    | `<-ch`      | Receives normally              |
Receive Only (`<-chan`) | closed        | `<-ch`      | Returns zero value and `false` | [close], [recv]
Uninitialized (`nil`)   | uninitialized | `ch <- v`   | Blocks forever                 | [send]
Uninitialized (`nil`)   | uninitialized | `<-ch`      | Blocks forever                 | [recv]
Uninitialized (`nil`)   | uninitialized | `close(ch)` | Causes a Run-Time Panic        | [close]

## References:

- [chan]
- [make]
- [close]
- [send]
- [recv]
- [for]
- [select]
- [len_cap]


[chan]: https://golang.org/ref/spec#Channel_types

[make]: https://golang.org/ref/spec#Making_slices_maps_and_channels

[close]: https://golang.org/ref/spec#Close

[send]: https://golang.org/ref/spec#Send_statements

[recv]: https://golang.org/ref/spec#Receive_operator

[for]: https://golang.org/ref/spec#For_statements

[select]: https://golang.org/ref/spec#Select_statements

[len_cap]: https://golang.org/ref/spec#Length_and_capacity
