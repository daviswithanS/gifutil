# gifutil
gifutil makes use of the gif.GIF struct found in the "image/gif" library. This struct, although useful, is relatively manual, requiring one to edit the fields and set up the options themselves. As such, I have created this library as a supplement to "image/gif", containing useful functions to help create and manage these data structures. 

NOTE: This library offers basic creation and output of gifs. If one requires more advanced features such as Disposal Methods, they will have to edit the GIF struct themselves.

# Populate()
While most of the functions in the library are rather self-explanatory, there is one that perhaps requires a bit more explanation. The Populate() function offers a way to automatically create the frames of a gif according to a custom function defined by the user. This function is passed in the parameters of the Populate() function. 

The function should take as input an integer, representing the current frame number, and return as output a pointer to an image.Image, representing what the frame of the gif should look like. This is useful for a number of reasons, ranging from loading a group of files with consistent filenames (e.g. image0.png, image1.png, image2.png...) to creating the images using a drawing library and returning the output at each step. 

If one wishes to bind variables to their custom function, such as an increasing variable or a drawing context, they can use closure, which allows a function to keep persistent variables through its calls. For more information on closures, see https://tour.golang.org/moretypes/25 and the examples provided with this library.
