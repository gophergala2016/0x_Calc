# 0x_Calc : Hexadecimal Calculator with GO

Helllo! You may noticed the famous prefix for hexadecimal literals(0x).
What we made is simple calculator with golang. 

We studied golang since **10 days** ago. Our focus was learning several syntax features of golang, which means the quality of program is **TERRIBLE**.   
But we don't care. We are satisfied with this short codes.

Have a nice day!

## Dependencies
To implement the GUI, we used `go-gtk`. Thanks for their endeavors.  
link : https://mattn.github.io/go-gtk/

## Screenshot
1. Full figure
```
[full]
```
As you can see in this figure, there is nothing complex. 

1. Top half
```
[part]
```
The buttons at top-left corner are for radix. When you click one of them, Labels at the right of them will react to your selection.  
But, **only** button click. We wanted to use glib for keyboard input, but we couldn't make it.  

1. Bottom half
```
[part]
```
To talk turkey, We are also first at GTK. We didn't have enough knowledge for it. We used table and jagged slice for 2 panels.  
Left frame, is for number buttons. The other frame is for operators. That's it!

## How to use
Well, I don't recommend not to use them. 

## License
We DON'T care. It's totally out of our hand. 
