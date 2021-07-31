# Go (Golang) Channels

|                                 | Send:<br>`ch <- v`<br>([send]) | Receive:<br>`<-ch`<br>([recv]) | Close:<br>`close(ch)`<br>([close]) |
| ------------------------------- | ------------------------------ | ------------------------------ | ---------------------------------- |
| **Bidirectional: <br> `chan`**  | \[Not Closed\]<br>Sends normally<br><br>\[Closed\]<br>Run-time panic | \[Not Closed\]<br>Receives normally<br><br>\[Closed\]<br>Returns zero value and `false` | \[Not Closed\]<br>Closes the channel<br><br>\[Closed\]<br>Run-time panic |
| **Send-Only:<br>`chan<-`**      | \[Not Closed\]<br>Sends normally<br><br>\[Closed\]<br>Run-time panic | Compilation error | \[Not Closed\]<br>Closes the channel<br><br>\[Closed\]<br>Run-time panic |
| **Receive-Only:<br>`<-chan`**   | Compilation error | \[Not Closed\]<br>Receives normally<br><br>\[Closed\]<br>Returns zero value and `false` | Compilation error |
| **Uninitialized:<br>`nil`**     | Blocks forever | Blocks forever | Run-time panic |

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
