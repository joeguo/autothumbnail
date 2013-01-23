autothumbnail
=============

Autothumbnail is a simple golang application, which can generated image thumbnail based on file system notification(fsnotify) and imagemagic easily.

Usage
=============
1. Install [fsnotify](https://github.com/howeyc/fsnotify) and [imagemagic](http://www.imagemagick.org/) first.
2. Clone the source and build:

git clone git://github.com/joeguo/autothumbnail.git

go build -o autothumb

3. Use the autothumb to monitor a folder:

autothumb [options]

The options are:
    folder      the folder to monitor, default folder is /var/www/images
    target      target thumbs folder, default target is /var/www/images/thumbnail
    size        thumb size, default size is 341x267
    wait        Time(seconds) to wait the image written done before generating thumb image, default 10 seconds

for example:

autothumb -folder=/var/www/images -target /var/www/thumbnail -size=400x300 -wait=5






