# gifutil
gifutil makes use of the gif.GIF struct found in the "image/gif" library. This struct, although useful, is relatively manual, requiring one to edit the fields and set up the options themselves. As such, I have created this library as a supplement to "image/gif", containing useful functions to help create and manage these data structures. 

NOTE: This library offers basic creation and output of gifs. If one requires more advanced features such as Disposal Methods, they will have to edit the GIF struct themselves.

Documentation can be found at https://godoc.org/github.com/daviswithanS/gifutil.
