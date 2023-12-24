# diamondSquare-go
Go language implementation of diamond square algorithm for heightmap generation.

The implementations follows the procedure described in [the Wikipedia article for the algorithm](https://en.wikipedia.org/wiki/Diamond-square_algorithm).

# diamondSquare-cli
diamondSquare-cli located in /cmd/diamondSquare-cli.go is an application to generate heightmaps from CLI interface.

usage: 

```
$ go run cmd/diamondSquare-cli/main.go <size> <h> <output filename>
```
Parameter _\<size\>_ must must be positive integer (2^n)+1 where n natural. Example: 3, 9, 17, 33...

Paremeter _\<h\>_ is used to define the factor by which the scale of random part of each node is mutiplied each iteration, as described in the Wikipedia article. 
The value must be between 0 and 1 represented in _float_. 
Values closer to 0 generate rougher heightmaps while  values closer to 1 generate smoother heightmaps.

Parameter _\<output filename\>_ must end in either ".png", ".jpg", ".jpeg" or ".gif".

# diamondSquare-gui
diamondSquare-gui located in /cmd/diamondSquare-gui.go is an application to run the implementation of the diamondSquare algorithm with configurable parameters and output to a [raylib](https://www.raylib.com/) GUI window.

# pkg/heightmap
The package heightmap can be used to generate new and interact with heightmaps.
