# diamondSquare-go
Go language implementation of diamond square algorithm for heightmap generation

The implementations follows the procedure described in [the Wikipedia article for the algorithm](https://en.wikipedia.org/wiki/Diamond-square_algorithm) 

usage: 

```
$diamondSquare-go <size> <h> <output filename>
```
Paremeter _\<h\>_ is used to define the factor by which the scale of random part of each node is mutiplied each iteration, as described in the Wikipedia article. 
The value must be between 0 and 1 represented in _float_. 
Values closer to 0 generate rougher heightmaps while  values closer to 1 generate smoother heightmaps.

Parameter _\<output filename\>_ must end in either ".png", ".jpg", ".jpeg" or ".gif"
