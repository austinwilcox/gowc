# GO Word Count
This is a simple implementation of the gnu utility ```wc```.

The cli takes the following inputs:
```
gowc test.txt
```
will produce output like:
```
7145  58164 342190 test.txt
```

You can supply the following flags to gather misc. details
```-c: number of bytes in a file```
```-l: number of lines in a file```
```-w: number of words in a file```
```-m: number of characters in utf8 in a file```

The program will also work by piping input it like so:
```cat test.txt | gowc```
This will produce:
```
7145  58164 342190
```

and just like calling ```gowc test.txt``` you can supply the command line arguments to filter out what you would like.
